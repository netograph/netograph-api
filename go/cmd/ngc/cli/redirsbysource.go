package cli

import (
	"fmt"
	"io"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func redirsBySourceCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:     "redirsbysource address",
		Aliases: []string{"rsrc"},
		Short:   "Redirections matching a domain query on source",
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

			r, err := c.RedirsBySource(ctx, &dsetapi.RedirsBySourceRequest{
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
