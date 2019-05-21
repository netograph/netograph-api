package api

import (
	"io"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/userapi"
	"github.com/spf13/cobra"
)

func datasetsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "datasets",
		Aliases: []string{"dsets"},
		Short:   "List all available datasets",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectUser()
			if err != nil {
				return err
			}
			r, err := c.Datasets(ctx, &userapi.DatasetsRequest{})
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
				err = utils.Output(cmd, v)
				if err != nil {
					return err
				}
			}
		},
	}
	return cmd
}
