package main

import (
	"fmt"
	"os"

	"github.com/maleo/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var maleoCmd = &cobra.Command{
		Use:   "maleo",
		Short: "The Maleo Blockchain CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	maleoCmd.AddCommand(cmd.VersionCmd)
	maleoCmd.AddCommand(cmd.BalancesCmd())
	maleoCmd.AddCommand(cmd.TxCmd())

	err := maleoCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
