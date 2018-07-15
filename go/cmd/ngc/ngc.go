package main

import (
	"os"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
