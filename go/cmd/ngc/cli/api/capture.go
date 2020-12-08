package api

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
		Use:   "capture urls_filepath",
		Short: "Bulk request captures by urls from file (1 url per line). Assets expire in 24 hours",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			urlsFile, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer urlsFile.Close()

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

			var caps []*userapi.Capture
			scanner := bufio.NewScanner(urlsFile)
			for scanner.Scan() {
				url := strings.TrimSpace(scanner.Text())
				if url == "" {
					continue
				}
				cap := &userapi.Capture{
					Urls:         []string{url},
					Notification: *notification,
					Meta:         ms,
					Zone:         *zone,
					Extended:     *extended,
				}
				caps = append(caps, cap)
			}

			if err := scanner.Err(); err != nil {
				fmt.Printf("Error while reading from file: %s", err)
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
