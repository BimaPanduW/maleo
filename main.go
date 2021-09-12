package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var maleoCmd = &cobra.Command{
		Use:   "maleo",
		Short: "The Maleo Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	maleoCmd.AddCommand(versionCmd)
	maleoCmd.AddCommand(balancesCmd())

	err := maleoCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
