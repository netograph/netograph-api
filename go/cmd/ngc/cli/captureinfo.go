package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/netograph/netograph-api/go/proto/ngapi"
)

func captureInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "captureinfo id",
		Aliases: []string{"cap"},
		Short:   "Retrieve summary info for a specified capture",
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
			r, err := c.CaptureInfo(
				ctx,
				&ngapi.CaptureInfoRequest{
					Dataset: viper.GetString("dset"),
					Id:      args[0],
				},
			)
			if err != nil {
				return err
			}

			if err := Output(cmd, r); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
