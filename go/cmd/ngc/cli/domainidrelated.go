package cli

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func domainidRelated() *cobra.Command {
	var resume *string
	var exact *bool
	cmd := &cobra.Command{
		Use:     "domainidrelated domain",
		Aliases: []string{"didrelated"},
		Short:   "Finds domains related through matching domain IDs",
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

			key := ""
			if len(args) == 2 {
				key = args[1]
			}

			r, err := c.DomainIDsForDomain(ctx, &dsetapi.DomainIDsForDomainRequest{
				Dataset: viper.GetString("dset"),
				Limit:   limit,
				Domain:  args[0],
				Key:     key,
				Resume:  *resume,
				Exact:   *exact,
			})
			if err != nil {
				return err
			}
			ids := []*dsetapi.DomainIDsForDomainResult{}
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
					Dataset: viper.GetString("dset"),
					Limit:   limit,
					Key:     id.Key,
					Value:   id.Value,
					Exact:   *exact,
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
	exact = cmd.Flags().BoolP("exact", "x", false, "Return exact domain results, instead of TLD+1")
	return cmd
}
