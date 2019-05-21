package api

import (
	"fmt"
	"io"
	"strconv"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func certIPCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:     "certip address [mask]",
		Aliases: []string{"cips"},
		Short:   "Certificates matching an IP address query",
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
			c, ctx, err := utils.ConnectDset()
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

			ip, err := resolveIP(args[0])
			if err != nil {
				return err
			}

			limit, err := cmd.Root().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}

			r, err := c.CertIPSearch(ctx, &dsetapi.CertIPSearchRequest{
				Dataset: viper.GetString("dset"),
				Ip:      ip,
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
				} else if err := utils.Output(cmd, v); err != nil {
					return err
				}
			}
		},
	}
	resume = cmd.Flags().StringP("resume", "r", "", "Resume retrieval from a specified token")
	return cmd
}
