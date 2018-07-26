package cli

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/proto/ngapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ipsForDomainCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:     "ipsfordomain address",
		Aliases: []string{"ipdom"},
		Short:   "Retrieve IP addresses associated with a domain query",
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

			r, err := c.IPsForDomain(ctx, &ngapi.IPsForDomainRequest{
				Dataset: viper.GetString("dset"),
				Query:   args[0],
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
				err = Output(cmd, v)
				if err != nil {
					return err
				}
			}
		},
	}
	resume = cmd.Flags().StringP("resume", "r", "", "Resume retrieval from a specified token")
	return cmd
}
