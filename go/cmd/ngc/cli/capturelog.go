package cli

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/netograph/netograph-api/go/proto/ngapi"
)

func captureLogCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:     "capturelog",
		Aliases: []string{"caplog"},
		Short:   "Log of captures in reverse chronological order",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connect()
			if err != nil {
				return err
			}
			limit, err := cmd.Parent().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}
			r, err := c.CaptureLog(ctx, &ngapi.CaptureLogRequest{
				Dataset: viper.GetString("dset"),
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
