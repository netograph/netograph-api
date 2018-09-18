package cli

import (
	"fmt"
	"io"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var timeFormats = []string{
	time.RFC3339,
	"2006-01-02",
	"2006/01/02",
}

func parseTime(s string) (*timestamp.Timestamp, error) {
	if s == "" {
		return &timestamp.Timestamp{}, nil
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}

	var t time.Time
	for _, f := range timeFormats {
		t, err = time.ParseInLocation(f, s, loc)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("could not parse date specification: %s", s)
	}
	return ptypes.TimestampProto(t)
}

func captureLogCommand() *cobra.Command {
	var resume *string
	var start *string
	var end *string
	cmd := &cobra.Command{
		Use:     "capturelog",
		Aliases: []string{"caplog"},
		Short:   "Log of captures in reverse chronological order",
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
			c, ctx, err := connectDset()
			if err != nil {
				return err
			}
			limit, err := cmd.Parent().PersistentFlags().GetInt64("limit")
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
				} else if err := Output(cmd, v); err != nil {
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
