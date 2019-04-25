package cli

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func domainidsForDomain() *cobra.Command {
	var resume *string
	var exact *bool
	cmd := &cobra.Command{
		Use:     "domainidsfordomain domain [key]",
		Aliases: []string{"diddom"},
		Short:   "DomainID entries for a domain, optionally limited by key",
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
