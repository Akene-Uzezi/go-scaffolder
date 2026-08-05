package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-scaffold",
	Short: "A CLI tool meant to scaffold new Go projects",
	Long: `
  ________                   _________              _____  _____      .__       .___
 /  _____/  ____            /   _____/ ____ _____ _/ ____\/ ____\____ |  |    __| _/
/   \  ___ /  _ \   ______  \_____  \_/ ___\\__  \\   __\\   __\/  _ \|  |   / __ |
\    \_\  (  <_> ) /_____/  /        \  \___ / __ \|  |   |  | (  <_> )  |__/ /_/ |
 \______  /\____/          /_______  /\___  >____  /__|   |__|  \____/|____/\____ |
        \/                         \/     \/     \/                              \/

This is go scaffold. A CLI tool meant to scaffold new Go projects.`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("An Error occured: %s", err)
		os.Exit(1)
	}
}
