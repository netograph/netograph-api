package api

import (
	"fmt"
	"io"
	"strings"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/userapi"
	"github.com/spf13/cobra"
)

func captureCommand() *cobra.Command {
	var notification *string
	var meta *[]string
	var zone *string
	var extended *bool
	cmd := &cobra.Command{
		Use:   "capture url [url...]",
		Short: "Bulk request captures. Assets expire in 24 hours",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectUser()
			if err != nil {
				return err
			}
			ms := make([]*userapi.Metadata, len(*meta))
			for i, v := range *meta {
				vals := strings.SplitN(v, "=", 2)
				if len(vals) != 2 {
					return fmt.Errorf("Invalid metadata specification: %s", v)
				}
				ms[i] = &userapi.Metadata{
					Key:   vals[0],
					Value: vals[1],
				}
			}
			for i, v := range args {
				args[i] = strings.TrimSpace(v)
			}

			caps := make([]*userapi.Capture, len(args))
			for i, arg := range args {
				caps[i] = &userapi.Capture{
					Urls:         []string{arg},
					Notification: *notification,
					Meta:         ms,
					Zone:         *zone,
					Extended:     *extended,
				}
			}

			capreq := userapi.CaptureRequest{Captures: caps}
			r, err := c.Capture(ctx, &capreq)
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
				err = utils.Output(cmd, v)
				if err != nil {
					return err
				}
			}
		},
	}
	notification = cmd.Flags().String(
		"notify", "", "Notify a URL after capture completes",
	)
	meta = cmd.Flags().StringArray(
		"meta",
		[]string{},
		"Metadata in the format key=value",
	)
	zone = cmd.Flags().String(
		"zone", "", "Capture zone (\"eu\" or \"us\")",
	)
	extended = cmd.Flags().Bool(
		"extended", false, "Extended capture",
	)
	return cmd
}
