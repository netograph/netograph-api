package api

import (
	"io"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func captureLogCommand() *cobra.Command {
	var resume *string
	var start *string
	var end *string
	cmd := &cobra.Command{
		Use:     "capturelog",
		Aliases: []string{"caplog"},
		Short:   "Captures in reverse chronological order",
		Long: `
Log of captures in reverse chronological order

Since the order is reverse chronological, the start time is the most recent
time in the range, and the end time is the least recent time in the range.

Times can be specified in the following formats:

	2006-01-02T15:04:05Z07:00 (RFC3339)
	yyyy/mm/dd
	yyyy-mm-dd
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectDset()
			if err != nil {
				return err
			}
			limit, err := cmd.Root().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}
			st, err := parseTime(*start)
			if err != nil {
				return err
			}
			et, err := parseTime(*end)
			if err != nil {
				return err
			}

			r, err := c.CaptureLog(ctx, &dsetapi.CaptureLogRequest{
				Dataset: viper.GetString("dset"),
				Limit:   limit,
				Resume:  *resume,
				Start:   st,
				End:     et,
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
	start = cmd.Flags().StringP("start", "s", "", "Start time (most recent time in the range)")
	end = cmd.Flags().StringP("end", "e", "", "End time (least recent time in the range)")
	return cmd
}
