package pkg

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func init() {
	var config Config
	f, err := os.ReadFile("utils/config.yaml")
	CheckError(err)
	err = yaml.Unmarshal(f, &config)
	CheckError(err)

	for i := 0; i < config.Sources; i++ {
		go createRequest(i)
	}
	in = inDIspatcher{}
	b = buffer{
		buf:         make([]*request, config.Bufsize),
		iterIn:      0,
		iterDecline: 0,
	}
	out = outDispatcher{
		devices: make([]device, config.Devices),
		iter:    0,
	}
	go out.run()
}

func CheckError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\nUsage of %v\n", err, os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func StartSimulation() {
	start = true
}
