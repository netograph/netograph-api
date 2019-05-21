package api

import (
	"fmt"
	"strings"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func submitCaptureCommand() *cobra.Command {
	var notification *string
	var skiprecent *int64
	var meta *[]string
	var zone *string
	var extended *bool
	cmd := &cobra.Command{
		Use:     "submitcapture url [url...]",
		Aliases: []string{"submit"},
		Short:   "Submit a capture request",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connectDset()
			if err != nil {
				return err
			}

			ms := make([]*dsetapi.Metadata, len(*meta))
			for i, v := range *meta {
				vals := strings.SplitN(v, "=", 2)
				if len(vals) != 2 {
					return fmt.Errorf("Invalid metadata specification: %s", v)
				}
				ms[i] = &dsetapi.Metadata{
					Key:   vals[0],
					Value: vals[1],
				}
			}

			for i, v := range args {
				args[i] = strings.TrimSpace(v)
			}

			cap := dsetapi.SubmitCaptureRequest{
				Dataset:      viper.GetString("dset"),
				Urls:         args,
				Notification: *notification,
				Skiprecent:   *skiprecent,
				Meta:         ms,
				Zone:         *zone,
				Extended:     *extended,
			}

			r, err := c.SubmitCapture(ctx, &cap)
			if err != nil {
				return err
			}

			if err := utils.Output(cmd, r); err != nil {
				return err
			}
			return nil
		},
	}
	notification = cmd.Flags().String(
		"notify", "", "Notify a URL after capture completes",
	)
	skiprecent = cmd.Flags().Int64(
		"skiprecent", 0, "Skip capture if we've seen a matching submission in the last N seconds",
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
