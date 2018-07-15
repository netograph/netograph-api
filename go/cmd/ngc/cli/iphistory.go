package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/netograph/netograph-api/go/proto/ngapi"
)

func ipHistoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "iphistory address",
		Aliases: []string{"iphist"},
		Short:   "Retrieve the capture history for an IP",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connect()
			if err != nil {
				return err
			}

			limit, err := cmd.Parent().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}

			r, err := c.IPHistory(ctx, &ngapi.IPHistoryRequest{
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
