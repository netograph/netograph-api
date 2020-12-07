package api

import (
	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func captureInfoCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "captureinfo id",
		Aliases: []string{"cap"},
		Short:   "Info for a specified capture",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectDset()
			if err != nil {
				return err
			}
			r, err := c.CaptureInfo(
				ctx,
				&dsetapi.CaptureInfoRequest{
					Dataset: viper.GetString("dset"),
					Id:      args[0],
				},
			)
			if err != nil {
				return err
			}

			if err := utils.Output(cmd, r); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
