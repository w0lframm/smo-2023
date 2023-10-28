package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/w0lframm/smo-2023/pkg"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Hi there hello!")

	// check for unknown commandline arguments
	if len(flag.Args()) != 0 {
		checkError(errors.New("error: invalid command line"))
	}

	var config pkg.Config
	f, err := os.ReadFile("utils/config.yaml")
	checkError(err)
	err = yaml.Unmarshal(f, &config)
	checkError(err)
	spew.Dump(config)
}

func checkError(err error) {
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\nUsage of %v\n", err, os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
}
