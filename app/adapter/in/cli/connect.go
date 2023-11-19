package cli

import (
	"archetype/app/shared/archetype"
	"archetype/app/shared/archetype/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "short description of your command",
	Run:   runconnect,
}

func runconnect(cmd *cobra.Command, args []string) {
	if err := archetype.Setup(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("einar cli connected.")
}
