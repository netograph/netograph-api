package api

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func domainidDomainSearch() *cobra.Command {
	var resume *string
	var basedomain *bool
	cmd := &cobra.Command{
		Use:     "domiddomain domain [key] [value]",
		Aliases: []string{"diddom"},
		Short:   "DomainID entries for a domain, optionally limited by key and value",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 && len(args) != 2 && len(args) != 3 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectDset()
			if err != nil {
				return err
			}
			limit, err := cmd.Root().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}

			key := ""
			if len(args) == 2 {
				key = args[1]
			}

			value := ""
			if len(args) == 3 {
				value = args[2]
			}

			r, err := c.DomainIDDomainSearch(ctx, &dsetapi.DomainIDDomainSearchRequest{
				Dataset:    viper.GetString("dset"),
				Limit:      limit,
				Domain:     args[0],
				Key:        key,
				Value:      value,
				Resume:     *resume,
				Basedomain: *basedomain,
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
				} else if err := utils.Output(cmd, v); err != nil {
					return err
				}
			}
		},
	}
	resume = cmd.Flags().StringP("resume", "r", "", "Resume retrieval from a specified token")
	basedomain = cmd.Flags().BoolP(
		"basedomain", "b", false,
		"Return results for the base domain, rather than the exact domain",
	)
	return cmd
}
