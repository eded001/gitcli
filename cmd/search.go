package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitcli/internal/service"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var outputJSON bool

var searchCmd = &cobra.Command{
	Use:   "search <username|id>",
	Short: "Use para procurar dados de determinado usuário com base no seu username ou id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		username := args[0]
		user := service.SearchUser(username)

		if user != nil {
			if outputJSON {
				jsonData, err := json.MarshalIndent(user, "", "  ")
				if err != nil {
					fmt.Println("Erro ao converter dados para JSON:", err)
					return err
				}
				fmt.Println(string(jsonData))
				return nil
			}

			var tableData pterm.TableData

			if name := user.GetName(); name != "" {
				tableData = append(tableData, []string{pterm.LightCyan("Nome:"), pterm.White(name)})
			}

			if id := user.GetID(); id != 0 {
				tableData = append(tableData, []string{pterm.LightCyan("ID:"), pterm.White(strconv.FormatInt(id, 10))})
			}

			if login := user.GetLogin(); login != "" {
				tableData = append(tableData, []string{pterm.LightCyan("Username:"), pterm.White(login)})
			}

			if company := user.GetCompany(); company != "" {
				tableData = append(tableData, []string{pterm.LightCyan("Empresa:"), pterm.LightMagenta(company)})
			}

			if location := user.GetLocation(); location != "" {
				tableData = append(tableData, []string{pterm.LightCyan("Local:"), pterm.LightGreen(location)})
			}

			if email := user.GetEmail(); email != "" {
				tableData = append(tableData, []string{pterm.LightCyan("Email:"), pterm.LightMagenta(email)})
			}

			if blog := user.GetBlog(); blog != "" {
				tableData = append(tableData, []string{pterm.LightCyan("Site/Blog:"), pterm.LightBlue(blog)})
			}

			if twitter := user.GetTwitterUsername(); twitter != "" {
				tableData = append(tableData, []string{pterm.LightCyan("Twitter:"), pterm.LightBlue("@" + twitter)})
			}

			tableData = append(tableData, []string{pterm.LightCyan("Repositórios:"), pterm.LightRed(strconv.Itoa(user.GetPublicRepos()))})
			tableData = append(tableData, []string{pterm.LightCyan("Seguidores:"), pterm.LightRed(strconv.Itoa(user.GetFollowers()))})
			tableData = append(tableData, []string{pterm.LightCyan("Seguindo:"), pterm.LightRed(strconv.Itoa(user.GetFollowing()))})

			if url := user.GetHTMLURL(); url != "" {
				tableData = append(tableData, []string{pterm.LightCyan("URL:"), pterm.LightBlue(url)})
			}

			tableStr, err := pterm.DefaultTable.WithData(tableData).Srender()
			if err == nil {
				pterm.DefaultBox.
					WithTitle(pterm.LightGreen(" GitHub User Info ")).
					WithTitleTopCenter().
					Println(tableStr)
			} else {
				fmt.Println("Erro ao formatar os dados:", err)
			}
		}
		return nil
	},
}

func init() {
	searchCmd.Flags().BoolVarP(&outputJSON, "json", "j", false, "Retorna a saída no formato JSON")
	rootCmd.AddCommand(searchCmd)
}
