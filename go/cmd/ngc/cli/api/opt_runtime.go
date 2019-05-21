package api

import (
	"github.com/spf13/cobra"
)

// RunConfOpts collects runtime options
type RunConfOpts struct {
	Addr *string
}

func (o *RunConfOpts) Register(cmd *cobra.Command) {
	o.Addr = cmd.Flags().String("addr", "grpc.netograph.io:443", "API server address")
}
