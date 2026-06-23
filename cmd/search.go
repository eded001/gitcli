package cmd

import (
	"gitcli/internal/service"
	"gitcli/internal/ui"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <username|id>",
	Short: "Use para procurar dados de determinado usuário com base no seu username ou id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		user := service.SearchUser(username)

		if user == nil {
			return nil
		}

		return ui.PrintUserInfo(user)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
