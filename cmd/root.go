package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitcli",
	Short: "Uma CLI simples para uso diário do Git",
}

func Execute() error {
	return rootCmd.Execute()
}
