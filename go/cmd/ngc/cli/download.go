package cli

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/briandowns/spinner"
	"github.com/netograph/netograph-api/go/proto/ngapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func downloadCommand() *cobra.Command {
	var direct *bool
	var quiet *bool
	cmd := &cobra.Command{
		Use:     "download path id [id...]",
		Aliases: []string{"get"},
		Short:   "Retrieve all assets for a capture to a directory",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := connect()
			if err != nil {
				return err
			}
			if len(args) > 2 && *direct {
				return fmt.Errorf("Cannot combine -direct with multiple download IDs")
			}

			for i := 1; i < len(args); i++ {
				id := args[i]
				s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
				if !*quiet {
					fmt.Printf("Fetching %s...\n", id)
					s.Suffix = fmt.Sprintf("\tfetching metadata")
					s.FinalMSG = "   "
					s.Color("green")
					s.Start()
				}
				r, err := c.CaptureInfo(
					ctx,
					&ngapi.CaptureInfoRequest{
						Dataset: viper.GetString("dset"),
						Id:      id,
					},
				)
				if !*quiet {
					s.Stop()
					fmt.Println()
				}
				if err != nil {
					return err
				}
				outdir := args[0]
				if !*direct {
					outdir = path.Join(outdir, id)
				}
				if _, err := os.Stat(outdir); os.IsNotExist(err) {
					err := os.MkdirAll(outdir, 0700)
					if err != nil {
						return err
					}
				} else if err != nil {
					return err
				}
				ca := coreAssets[:]
				for i := range r.Capsummary.Roots {
					screenshot := fmt.Sprintf("screenshot%.3d.jpg", i)
					ca = append(ca, screenshot)
				}
				for _, a := range ca {
					s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
					if !*quiet {
						s.FinalMSG = " "
						s.Suffix = fmt.Sprintf("\t%s  ", a)
						s.Color("green")
						s.Start()
					}
					err := downloadFile(
						path.Join(outdir, a),
						fmt.Sprintf("%s/%s", r.Capsummary.Assets, a),
					)
					if err != nil {
						fmt.Fprintf(os.Stderr, "error: %s", err)
					}
					if !*quiet {
						s.Stop()
						fmt.Println()
					}
				}
			}
			return nil
		},
	}
	direct = cmd.Flags().BoolP(
		"direct", "d", false, "Save directly to the path, without prefixing with the capture ID ",
	)
	quiet = cmd.Flags().BoolP(
		"quiet", "q", false, "Silence progress output",
	)
	return cmd
}
