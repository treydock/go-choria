require "securerandom"

OWN_OS=`go env GOOS`.chomp
OWN_ARCH=`go env GOARCH`.chomp

task :default => [:test]

desc "Run just tests no measurements"
task :test do
  sh "ginkgo -r -skipMeasurements ."
end

desc "Run tests including measure tests"
task :test_and_measure do
  sh "ginkgo -r ."
end

desc "Builds a static binary"
task :build do
  ENV["GOOS"] ||= `go env GOOS`.chomp
  ENV["GOARCH"] ||= `go env GOARCH`.chomp
  version = ENV["VERSION"] || "development"
  sha = `git rev-parse --short HEAD`.chomp
  date = Time.now.strftime("%F %T %z")
  buildid = SecureRandom.hex

  flags = [
    "-X github.com/choria-io/go-choria/build.Version=%s" % version,
    "-X github.com/choria-io/go-choria/build.SHA=%s" % sha,
    "-X \"github.com/choria-io/go-choria/build.BuildDate=%s\"" % date,
    "-B 0x%s" % buildid
]

  if ENV["BUILD_XFLAGS"]
    ENV["BUILD_XFLAGS"].split("|").each do |flag|
      flags << "-X github.com/choria-io/go-choria/build.%s" % flag     
    end
  end

  args = [
    "-o %s" % output_name(version),
    "-ldflags '%s'" % flags.join(" ")
  ]

  args << "-race" if version == "development" && ENV["GOOS"] == OWN_OS

  cmd = "go build %s" % args.join(" ")

  sh cmd % [output_name(version), flags.join(" ")]
end

desc "Builds a Linux binary"
task :build_linux do
  ENV["GOOS"] = "linux"
  ENV["GOARCH"] = "amd64"

  Rake::Task["build"].execute
end

def output_name(version)
  return ENV["OUTPUT"] if ENV["OUTPUT"]

  if ENV["GOOS"] && ENV["GOARCH"]
    return "choria-%s-%s-%s" % [version, ENV["GOOS"].capitalize, ENV["GOARCH"]]
  else
    return "choria-%s" % [version]
  end
end
