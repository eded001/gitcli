package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gitcli",
	Short: "Uma CLI simples para uso do Git Cli",
}

func Execute() error {
	return rootCmd.Execute()
}