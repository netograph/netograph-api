package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:          "ngc",
	Short:        "A client for the Netograph API",
	SilenceUsage: true,
}

func init() {
	RootCmd.AddCommand(
		captureInfoCommand(),
		captureLogCommand(),
		datasetsCommand(),
		domainHistoryCommand(),
		domainSearchCommand(),
		domainsForIP(),
		downloadCommand(),
		ipSearchCommand(),
		ipsForDomain(),
		ipHistoryCommand(),
		ipLogSearchCommand(),
		metaForCaptureCommand(),
		metaSearchCommand(),
		rootLogSearchCommand(),
		rootsForSatelliteCommand(),
		satelliteLogSearchCommand(),
		satellitesForRoot(),
		submitCaptureCommand(),
		urlLogSearchCommand(),
	)
	viper.AutomaticEnv()
	viper.SetEnvPrefix("NGC")

	RootCmd.PersistentFlags().String("addr", "grpc.netograph.io:443", "API server address [NGC_ADDR]")
	viper.BindPFlag("addr", RootCmd.PersistentFlags().Lookup("addr"))

	RootCmd.PersistentFlags().String("dset", "netograph:social", "Dataset to query [NGC_DSET]")
	viper.BindPFlag("dset", RootCmd.PersistentFlags().Lookup("dset"))

	RootCmd.PersistentFlags().Bool(
		"cjson",
		false,
		`Compact JSON output format - one record per line, suitable for programmatic use`,
	)
	RootCmd.PersistentFlags().Bool(
		"json",
		false,
		`Indented JSON output format`,
	)
	RootCmd.PersistentFlags().Int64P(
		"limit", "n", 100,
		`Limit number of responses, 0 is unlimited. Please use this carefully,
the number of results returned can be very large.`,
	)

	RootCmd.PersistentFlags().String("token", "", "API authorization token [NGC_TOKEN]")
	viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))
}
