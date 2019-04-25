package cli

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func domainidTagSearch() *cobra.Command {
	var resume *string
	var exact *bool
	cmd := &cobra.Command{
		Use:     "domainidtagsearch key [value]",
		Aliases: []string{"didtag"},
		Short:   "DomainID entries for a tag key, optionally limited by value",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 && len(args) != 2 {
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

			value := ""
			if len(args) == 2 {
				value = args[1]
			}

			r, err := c.DomainIDTagSearch(ctx, &dsetapi.DomainIDTagSearchRequest{
				Dataset: viper.GetString("dset"),
				Limit:   limit,
				Key:     args[0],
				Value:   value,
				Resume:  *resume,
				Exact:   *exact,
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
	exact = cmd.Flags().BoolP("exact", "x", false, "Return exact domain results, instead of TLD+1")
	return cmd
}
