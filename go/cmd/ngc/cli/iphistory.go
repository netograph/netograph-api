package cli

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ipHistoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "iphistory address",
		Aliases: []string{"iphist"},
		Short:   "Capture history for an IP address",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connectDset()
			if err != nil {
				return err
			}

			limit, err := cmd.Parent().PersistentFlags().GetInt64("limit")
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
				err = Output(cmd, v)
				if err != nil {
					return err
				}
			}
		},
	}
	return cmd
}
