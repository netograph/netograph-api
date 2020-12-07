package api

import (
	"fmt"
	"io"
	"strings"

	"github.com/netograph/netograph-api/go/cmd/ngc/cli/utils"
	"github.com/netograph/netograph-api/go/proto/ngapi/dsetapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var validFields = map[string]bool{
	"subj.cn":      true,
	"subj.org":     true,
	"subj.ou":      true,
	"subj.loc":     true,
	"subj.prov":    true,
	"subj.country": true,

	"iss.cn":      true,
	"iss.org":     true,
	"iss.ou":      true,
	"iss.loc":     true,
	"iss.prov":    true,
	"iss.country": true,
}

func certSearchCommand() *cobra.Command {
	var resume *string
	cmd := &cobra.Command{
		Use:     "certsearch field text",
		Aliases: []string{"csearch"},
		Short:   "Certificates matching a text query",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("Usage: %s", cmd.Use)
			}
			if _, ok := validFields[args[0]]; !ok {
				fields := []string{}
				for k := range validFields {
					fields = append(fields, k)
				}
				opts := strings.Join(fields, ", ")
				return fmt.Errorf("Field must be one of: %s", opts)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c, ctx, err := utils.ConnectDset()
			if err != nil {
				return err
			}
			limit, err := cmd.Root().PersistentFlags().GetInt64("limit")
			if err != nil {
				return err
			}

			r, err := c.CertSearch(ctx, &dsetapi.CertSearchRequest{
				Dataset: viper.GetString("dset"),
				Field:   args[0],
				Text:    args[1],
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
				} else if err := utils.Output(cmd, v); err != nil {
					return err
				}
			}
		},
	}
	resume = cmd.Flags().StringP("resume", "r", "", "Resume retrieval from a specified token")
	return cmd
}
