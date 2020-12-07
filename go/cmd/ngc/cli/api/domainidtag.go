package api

import (
	"io"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func domainidTagSearch() *cobra.Command {
	var resume *string
	var basedomain *bool
	cmd := &cobra.Command{
		Use:     "domidtag key [value]",
		Aliases: []string{"didtag"},
		Short:   "DomainID entries for a tag key, optionally limited by value and domain",
		Args:    cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectDset()
			if err != nil {
				return err
			}
			limit, err := cmd.Root().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}

			value := ""
			if len(args) == 2 {
				value = args[1]
			}

			r, err := c.DomainIDTagSearch(ctx, &dsetapi.DomainIDTagSearchRequest{
				Dataset:    viper.GetString("dset"),
				Limit:      limit,
				Key:        args[0],
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
