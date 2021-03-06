package natsstream

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/choria-io/go-choria/broker/adapter/stats"
	"github.com/choria-io/go-choria/choria"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/prometheus/client_golang/prometheus"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type stream struct {
	servers   func() ([]choria.Server, error)
	clusterID string
	clientID  string
	topic     string
	conn      stan.Conn
	log       *log.Entry
	name      string

	work chan adaptable
	quit chan bool
}

type msg struct {
	Protocol string    `json:"protocol"`
	Data     string    `json:"data"`
	Sender   string    `json:"sender"`
	Time     time.Time `json:"time"`
}

func newStream(name string, work chan adaptable, logger *log.Entry) ([]*stream, error) {
	prefix := fmt.Sprintf("plugin.choria.adapter.%s.stream.", name)

	instances, err := strconv.Atoi(config.Option(prefix+"workers", "10"))
	if err != nil {
		return nil, fmt.Errorf("%s should be a integer number", prefix+"workers")
	}

	servers := config.Option(prefix+"servers", "")
	if servers == "" {
		return nil, fmt.Errorf("No Stream servers configured, please set %s", prefix+"servers")
	}

	topic := config.Option(prefix+"topic", "")
	if topic == "" {
		topic = name
	}

	clusterID := config.Option(prefix+"clusterid", "")
	if clusterID == "" {
		return nil, fmt.Errorf("No ClusterID configured, please set %s", prefix+"clusterid'")
	}

	workers := []*stream{}

	for i := 0; i < instances; i++ {
		logger.Infof("Creating NATS Streaming Adapter %s NATS Streaming instance %d / %d publishing to %s on cluster %s", name, i, instances, topic, clusterID)

		iname := fmt.Sprintf("%s_%d-%s", name, i, strings.Replace(uuid.NewV4().String(), "-", "", -1))

		st := &stream{
			clusterID: clusterID,
			clientID:  iname,
			topic:     topic,
			name:      fmt.Sprintf("%s.%d", name, i),
			work:      work,
			log:       logger.WithFields(log.Fields{"side": "stream", "instance": i}),
		}
		st.servers = st.resolver(strings.Split(servers, ","))

		workers = append(workers, st)
	}

	return workers, nil
}

func (sc *stream) resolver(parts []string) func() ([]choria.Server, error) {
	servers, err := choria.StringHostsToServers(parts, "nats")
	return func() ([]choria.Server, error) {
		return servers, err
	}
}

func (sc *stream) connect(ctx context.Context, cm choria.ConnectionManager) error {
	if ctx.Err() != nil {
		return fmt.Errorf("Shutdown called")
	}

	nc, err := cm.NewConnector(ctx, sc.servers, sc.clientID, sc.log)
	if err != nil {
		return fmt.Errorf("Could not start NATS connection: %s", err)
	}

	sc.log.Infof("%s connecting to NATS Stream", sc.clientID)

	for {
		if ctx.Err() != nil {
			return fmt.Errorf("Shutdown called")
		}

		sc.conn, err = stan.Connect(sc.clusterID, sc.clientID, stan.NatsConn(nc.Nats()))
		if err != nil {
			sc.log.Errorf("Could not create initial STAN connection, retrying: %s", err)
			time.Sleep(time.Second)
			continue
		}

		break
	}

	sc.log.Infof("%s connected to NATS Stream", sc.clientID)

	return nil
}

func (sc *stream) disconnect() {
	if sc.conn != nil {
		sc.log.Info("Disconnecting from NATS Streaming")
		sc.conn.Close()
	}
}

func (sc *stream) publisher(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	bytes := stats.BytesCtr.WithLabelValues(sc.name, "output", config.Identity)
	ectr := stats.ErrorCtr.WithLabelValues(sc.name, "output", config.Identity)
	ctr := stats.ReceivedMsgsCtr.WithLabelValues(sc.name, "output", config.Identity)
	timer := stats.ProcessTime.WithLabelValues(sc.name, "output", config.Identity)

	transformerf := func(r adaptable) {
		obs := prometheus.NewTimer(timer)
		defer obs.ObserveDuration()

		m := msg{
			Protocol: "choria:adapters:natsstream:output:1",
			Data:     r.Message(),
			Sender:   r.SenderID(),
			Time:     r.Time().UTC(),
		}

		j, err := json.Marshal(m)
		if err != nil {
			sc.log.Warnf("Cannot JSON encode message for publishing to STAN, discarding: %s", err)
			ectr.Inc()
			return
		}

		sc.log.Debugf("Publishing registration data from %s to %s", m.Sender, sc.topic)

		bytes.Add(float64(len(j)))

		err = sc.conn.Publish(sc.topic, j)
		if err != nil {
			sc.log.Warnf("Could not publish message to STAN %s, discarding: %s", sc.topic, err)
			ectr.Inc()
			return
		}

		ctr.Inc()
	}

	for {
		select {
		case r := <-sc.work:
			transformerf(r)
		case <-ctx.Done():
			sc.disconnect()

			return
		}
	}
}
