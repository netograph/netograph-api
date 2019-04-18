package cli

import (
	"fmt"
	"io"
	"strconv"

	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func ipSearchCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:   "ipsearch address [mask]",
		Short: "IP addresses matching a netmask and address",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 || len(args) > 2 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			if len(args) == 2 {
				n, err := strconv.Atoi(args[1])
				if err != nil {
					return err
				}
				if n < 0 || n > 32 {
					return fmt.Errorf("Netmask must be 0-32")
				}
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connectDset()
			if err != nil {
				return err
			}
			n := 32
			if len(args) == 2 {
				var err error
				n, err = strconv.Atoi(args[1])
				if err != nil {
					return err
				}
			}

			limit, err := cmd.Parent().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}

			r, err := c.IPSearch(ctx, &dsetapi.IPSearchRequest{
				Dataset: viper.GetString("dset"),
				Ip:      args[0],
				Mask:    int32(n),
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
