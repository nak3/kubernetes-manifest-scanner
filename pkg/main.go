package main

import (
	"os"
	"runtime"

	"github.com/nak3/kubernetes-manifest-scanner/pkg/cmd"
)

func main() {

	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	cmd := cmd.KmsNew(os.Stdin, os.Stdout, os.Stderr)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
