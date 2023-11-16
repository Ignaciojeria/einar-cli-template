package cli

import (
	"archetype/app/shared/archetype/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(archetypeCmd)
}

var archetypeCmd = &cobra.Command{
	Use:   "_example",
	Short: "short description of your command",
	Run:   runArchetypeCmd,
}

func runArchetypeCmd(cmd *cobra.Command, args []string) {
	fmt.Println("_example command not implemented yet")
}
