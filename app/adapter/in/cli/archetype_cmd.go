package cli

import (
	"archetype/app/shared/archetype/cmd"
	"archetype/app/shared/archetype/container"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	container.InjectInboundAdapter(func() error {
		cmd.RootCmd.AddCommand(archetypeCmd)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

var archetypeCmd = &cobra.Command{
	Use:   "archetypeCmdUsage",
	Short: "short description of your command",
	Run:   runArchetypeCmd,
}

func runArchetypeCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Hello from archetypeCmdUsage")
}
