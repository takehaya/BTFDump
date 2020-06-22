module github.com/takehaya/btfdump

go 1.14

require (
	github.com/cilium/ebpf v0.0.0-20200617135954-7acf5cc039f4
	github.com/pkg/errors v0.9.1
	github.com/urfave/cli v1.22.4
)

replace github.com/cilium/ebpf v0.0.0-20200617135954-7acf5cc039f4 => /home/vagrant/go/src/github.com/cilium/ebpf
