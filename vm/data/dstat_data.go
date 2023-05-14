package data

var DstatParam = []string{
	"-l",
	"-c",
	"--proc",
	// "--proc-count",
	"-y",
	"-i",
	"--ipc",
	"-m",
	"--vm",
	"-g",
	"-s",
	"-n",
	"--net-packets",
	"--socket",
	"--raw",
	"--tcp",
	"--udp",
	"--unix",
	"-d",
	"-r",
	"--aio",
	"--disk-tps",
	"--disk-util",
	"--fs",
	"--lock",
}

// Invalid metric dstat.proc-count.total (Derived metric definition failed).

type DstatData struct {
}
