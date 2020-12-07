package api

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/buger/jsonparser"
	"github.com/spf13/cobra"
)

func tempDownloadCommand() *cobra.Command {
	var quiet *bool
	cmd := &cobra.Command{
		Use:     "tempdownload path url",
		Aliases: []string{"tget"},
		Short:   "Retrieve all assets of a temporary capture given the capture base URL",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			outdir := args[0]
			url := args[1]
			skip := 0
			for i := 1; i < len(args); i++ {
				if _, err := os.Stat(outdir); os.IsNotExist(err) {
					err := os.MkdirAll(outdir, 0700)
					if err != nil {
						return err
					}
				} else if err != nil {
					return err
				}
				for _, a := range coreAssets {
					err := downloadProgress(
						path.Join(outdir, a),
						fmt.Sprintf("%s/%s", url, a),
						*quiet,
					)
					if err != nil {
						skip++
					}
				}
			}
			data, err := ioutil.ReadFile(path.Join(outdir, "starmap.json"))
			if err != nil {
				return err
			}
			if len(data) == 0 {
				return fmt.Errorf("Could not download - perhaps the capture is still pending?")
			}

			cnt := 0
			jsonparser.ArrayEach(
				data,
				func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					cnt++
				},
				"plan", "urls",
			)

			for i := 0; i < cnt; i++ {
				screenshot := fmt.Sprintf("screenshot%.3d.jpg", i)
				err := downloadProgress(
					path.Join(outdir, screenshot),
					fmt.Sprintf("%s/%s", url, screenshot),
					*quiet,
				)
				if err != nil {
					break
				}
			}
			return nil
		},
	}
	quiet = cmd.Flags().BoolP(
		"quiet", "q", false, "Silence progress output",
	)
	return cmd
}
