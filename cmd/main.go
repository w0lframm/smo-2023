package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/w0lframm/smo-2023/pkg"
)

func main() {
	fmt.Println("Hi there hello!")

	// check for unknown commandline arguments
	if len(flag.Args()) != 0 {
		pkg.CheckError(errors.New("error: invalid command line"))
	}
	pkg.StartSimulation()
}
