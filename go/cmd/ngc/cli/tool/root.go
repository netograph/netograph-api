package tool

import (
	"github.com/spf13/cobra"
)

var ToolRootCmd = &cobra.Command{
	Use:          "tool",
	Short:        "Higher level tools that use the Netograph API",
	SilenceUsage: true,
}

func init() {
	ToolRootCmd.AddCommand(
		domainidRelated(),
	)
}
