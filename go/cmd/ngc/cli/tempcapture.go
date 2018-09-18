package cli

import (
	"fmt"
	"strings"

	"github.com/netograph/netograph-api/go/proto/ngapi/userapi"
	"github.com/spf13/cobra"
)

func tempCaptureCommand() *cobra.Command {
	var notification *string
	var meta *[]string
	cmd := &cobra.Command{
		Use:     "tempcapture url [url...]",
		Aliases: []string{"submit"},
		Short:   "Submit a capture request for a temp capture (no dataset, assets expire in 24 hours)",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connectUser()
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
			cap := userapi.TempCaptureRequest{
				Urls:         args,
				Notification: *notification,
				Meta:         ms,
			}

			r, err := c.TempCapture(ctx, &cap)
			if err != nil {
				return err
			}

			if err := Output(cmd, r); err != nil {
				return err
			}
			return nil
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
	return cmd
}
