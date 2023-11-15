package cli

import (
	"archetype/app/shared/archetype/cobra_cli"
	"archetype/app/shared/archetype/container"
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	container.InjectInboundAdapter(func() error {
		cobra_cli.RootCmd.AddCommand(cliArchetypeCmd)
		return nil
	}, container.InjectionProps{
		DependencyID: uuid.NewString(),
	})
}

var cliArchetypeCmd = &cobra.Command{
	Use:   "cliArchetype",
	Short: "short description of your command",
	Run:   runCliArchetypeCmd,
}

func runCliArchetypeCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Hello from cliArchetypeCmd")
}
