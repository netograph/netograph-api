package cli

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/netograph/netograph-api/go/proto/ngapi"
)

func satelliteLogSearchCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:     "satellitelogsearch query",
		Aliases: []string{"satlog"},
		Short:   "Search the log for captures where any satellite domain matches a query",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
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

			r, err := c.SatelliteLogSearch(ctx, &ngapi.SatelliteLogSearchRequest{
				Dataset: viper.GetString("dset"),
				Query:   args[0],
				Limit:   limit,
				Resume:  *resume,
			})
			if err != nil {
				return err
			}
			for {
				if v, err := r.Recv(); err != nil {
					if err == io.EOF {
						return nil
					}
					return err
				} else if err := Output(cmd, v); err != nil {
					return err
				}
			}
		},
	}
	resume = cmd.Flags().StringP("resume", "r", "", "Resume retrieval from a specified token")
	return cmd
}
