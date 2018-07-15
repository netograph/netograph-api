package cli

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/netograph/netograph-api/go/proto/ngapi"
)

func datasetsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "datasets",
		Aliases: []string{"dsets"},
		Short:   "List all available datasets",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connect()
			if err != nil {
				return err
			}
			r, err := c.Datasets(ctx, &ngapi.DatasetsRequest{})
			if err != nil {
				return err
			}
			for {
				v, err := r.Recv()
				if err != nil {
					if err == io.EOF {
						return nil
					}
					return err
				}
				err = Output(cmd, v)
				if err != nil {
					return err
				}
			}
		},
	}
	return cmd
}
