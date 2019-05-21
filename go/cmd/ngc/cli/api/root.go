package api

import (
	"github.com/spf13/cobra"
)

var APIRootCmd = &cobra.Command{
	Use:          "api",
	Short:        "Direct access to the Netograph API",
	SilenceUsage: true,
}

func init() {
	APIRootCmd.AddCommand(
		captureInfoCommand(),
		captureLogCommand(),
		certSearchCommand(),
		certDomainCommand(),
		certIPCommand(),
		certValidNamesCommand(),
		datasetsCommand(),
		domainHistoryCommand(),
		domainidDomainSearch(),
		domainidTagSearch(),
		domainSearchCommand(),
		domainsForIPCommand(),
		downloadCommand(),
		ipSearchCommand(),
		ipsForDomainCommand(),
		ipHistoryCommand(),
		ipLogSearchCommand(),
		metaForCaptureCommand(),
		metaSearchCommand(),
		policiesForRootCommand(),
		policyDomainCaptures(),
		policyStatsCommand(),
		policyURLCapturesCommand(),
		redirsByDestinationCommand(),
		redirsBySourceCommand(),
		rootLogSearchCommand(),
		rootsForSatelliteCommand(),
		satelliteLogSearchCommand(),
		satellitesForRootCommand(),
		submitCaptureCommand(),
		tempCaptureCommand(),
		tempDownloadCommand(),
		urlLogSearchCommand(),
	)
}
