package cli

import (
	"archetype/app/shared/archetype"
	"archetype/app/shared/archetype/cmd"
	"fmt"
	"os"
	"os/exec"

	"archetype/app/shared/config"

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

func init() {
	cmd.RootCmd.AddCommand(connectCmd)
}

func runconnect(cmd *cobra.Command, args []string) {
	if len(args) > 0 && args[0] == "setup-child" {
		// This is the child process for setup
		os.Setenv(string(config.PORT), "5556")
		if err := archetype.Setup(); err != nil {
			fmt.Fprintf(os.Stderr, "Setup failed: %s\n", err)
			os.Exit(1)
		}
		fmt.Println("Setup completed")
		os.Exit(0)
	}

	// Start the setup in a separate child process
	childProcess := exec.Command(os.Args[0], "connect", "setup-child")
	childProcess.Stdout = os.Stdout
	childProcess.Stderr = os.Stderr

	if err := childProcess.Start(); err != nil {
		fmt.Printf("Error starting setup child process: %s\n", err)
		return
	}

	fmt.Printf("Setup child process started, PID: %d\n", childProcess.Process.Pid)
}
