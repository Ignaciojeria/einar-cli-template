package cli

import (
	"archetype/app/adapter/out/client"
	"archetype/app/shared/archetype/cmd"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(shutdownCmd)
}

var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "short description of your command",
	Run:   runshutdown,
}

func runshutdown(cmd *cobra.Command, args []string) {
	client.Shutdown(cmd.Context())
}
