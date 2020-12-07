package api

import (
	"io"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func domainsForIPCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:     "domainsforip ip",
		Aliases: []string{"domip"},
		Short:   "Domains matching an IP query",
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

			r, err := c.DomainsForIP(ctx, &dsetapi.DomainsForIPRequest{
				Dataset: viper.GetString("dset"),
				Ip:      args[0],
				Limit:   limit,
				Resume:  *resume,
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
	resume = cmd.Flags().StringP("resume", "r", "", "Resume retrieval from a specified token")
	return cmd
}
