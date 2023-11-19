package main

import (
	_ "archetype/app/adapter/in/cli"
	"archetype/app/shared/archetype/cmd"
)

func main() {
	cmd.RootCmd.Execute()
}
