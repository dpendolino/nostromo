package cmd

import (
	"github.com/pokanop/nostromo/task"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a config item from manifest",
	Long: `Get a config item from manifest.
Nostromo config items are saved in the manifest.

Use this command to set keys to configure these settings:
	verbose: true | false`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		task.GetConfig(args[0])
	},
}

func init() {
	manifestCmd.AddCommand(getCmd)
}