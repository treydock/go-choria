package audit

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/choria-io/go-choria/choria"
	"github.com/choria-io/go-protocol/protocol/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFileContent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "McoRPC/Audit")
}

var _ = Describe("McoRPC/Audit", func() {
	It("Should correctly audit the request", func() {
		var cfg *choria.Config
		var err error

		if runtime.GOOS == "windows" {
			cfg, err = choria.NewConfig("testdata/audit_windows.cfg")
		} else {
			cfg, err = choria.NewConfig("testdata/audit.cfg")
		}

		os.Remove(cfg.Option("plugin.rpcaudit.logfile", "/tmp/rpc_audit.log"))

		Expect(err).ToNot(HaveOccurred())
		Expect(cfg.RPCAudit).To(BeTrue())
		Expect(cfg.Option("plugin.rpcaudit.logfile", "")).ToNot(BeAnExistingFile())

		req, err := v1.NewRequest("test_agent", "test.node", "choria=rip.mcollective", 120, "uniq_req_id", "mcollective")
		Expect(err).ToNot(HaveOccurred())

		ok := Request(req, "test_agent", "test_action", json.RawMessage(`{"hello":"world"}`), cfg)
		Expect(ok).To(BeTrue())
		Expect(cfg.Option("plugin.rpcaudit.logfile", "")).To(BeAnExistingFile())

		j, err := ioutil.ReadFile(cfg.Option("plugin.rpcaudit.logfile", ""))
		Expect(err).ToNot(HaveOccurred())

		am := Message{}
		err = json.Unmarshal(j, &am)
		Expect(err).ToNot(HaveOccurred())

		Expect(am.RequestID).To(Equal(req.RequestID()))
		Expect(am.RequestTime).To(Equal(req.Time().UTC().Unix()))
		Expect(am.CallerID).To(Equal("choria=rip.mcollective"))
		Expect(am.Sender).To(Equal("test.node"))
		Expect(am.Agent).To(Equal("test_agent"))
		Expect(am.Action).To(Equal("test_action"))
		Expect(am.Data).To(Equal(json.RawMessage(`{"hello":"world"}`)))
	})
})
