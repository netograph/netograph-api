package api

import (
	"io"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ipHistoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "iphistory address",
		Aliases: []string{"iphist"},
		Short:   "Capture history for an IP address",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectDset()
			if err != nil {
				return err
			}

			limit, err := cmd.Root().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}

			r, err := c.IPHistory(ctx, &dsetapi.IPHistoryRequest{
				Dataset: viper.GetString("dset"),
				Ip:      args[0],
				Limit:   limit,
			})
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
