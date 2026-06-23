package cmd

import (
	"gitcli/internal/service"

	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone <username|id>",
	Short: "Use para clonar um repositório público de determinado usuário com base no seu username ou id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		user := service.SearchUser(username)

		if user == nil {
			return nil
		}

		service.SearchRepos(username)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}