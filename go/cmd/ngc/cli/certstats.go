package cli

import (
	"fmt"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func certStatsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "certstats query",
		Aliases: []string{"cstats"},
		Short:   "Retrieve certificate statistics for a domain query",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connectDset()
			if err != nil {
				return err
			}
			r, err := c.CertDomainStats(ctx, &dsetapi.CertDomainStatsRequest{
				Dataset: viper.GetString("dset"),
				Query:   args[0],
			})
			if err != nil {
				return err
			}
			if err := Output(cmd, r); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
