package tool

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func domainidRelated() *cobra.Command {
	var resume *string
	var basedomain *bool
	cmd := &cobra.Command{
		Use:     "domidrelated domain",
		Aliases: []string{"didrelated"},
		Short:   "Finds domains related through matching domain IDs",
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

			key := ""
			if len(args) == 2 {
				key = args[1]
			}

			r, err := c.DomainIDDomainSearch(ctx, &dsetapi.DomainIDDomainSearchRequest{
				Dataset:    viper.GetString("dset"),
				Limit:      limit,
				Domain:     args[0],
				Key:        key,
				Resume:     *resume,
				Basedomain: *basedomain,
			})
			if err != nil {
				return err
			}
			ids := []*dsetapi.DomainIDDomainSearchResult{}
			for {
				if v, err := r.Recv(); err != nil {
					if err == io.EOF {
						break
					}
					return err
				} else {
					ids = append(ids, v)
				}
			}
			for _, id := range ids {
				r, err := c.DomainIDTagSearch(ctx, &dsetapi.DomainIDTagSearchRequest{
					Dataset:    viper.GetString("dset"),
					Limit:      limit,
					Key:        id.Key,
					Value:      id.Value,
					Basedomain: *basedomain,
				})
				if err != nil {
					return err
				}
				header := false
				for {
					if v, err := r.Recv(); err != nil {
						if err == io.EOF {
							break
						}
						return err
					} else {
						if !header {
							fmt.Printf("tag %s = %s\n", id.Key, id.Value)
							header = true
						}
						fmt.Printf("\t%s\n", v.Domain)
					}
				}
			}
			return nil
		},
	}
	resume = cmd.Flags().StringP("resume", "r", "", "Resume retrieval from a specified token")
	basedomain = cmd.Flags().BoolP(
		"basedomain", "b", false,
		"Return results for the base domain, rather than the exact domain",
	)
	return cmd
}
