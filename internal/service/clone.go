package service

import (
	"context"
	"fmt"
	"strconv"

	"os"
	"os/exec"

	"github.com/google/go-github/v88/github"
	"github.com/pterm/pterm"
)

func SearchRepos(username string) []*github.Repository {
	client, err := github.NewClient()

	if err != nil {
		fmt.Println("Erro ao criar cliente:", err)
		return nil
	}

	opt := &github.RepositoryListByUserOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	repos, _, err := client.Repositories.ListByUser(
		context.Background(),
		username,
		opt,
	)

	if repos == nil || len(repos) == 0 {
		pterm.Warning.Println("Nenhum repositório encontrado.")
		return nil
	}

	var options []string
	for _, repo := range repos {
		options = append(options, repo.GetName())
	}

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Selecione um repositório")

	for _, repo := range repos {
		if repo.GetName() == selectedOption {
			var tableData pterm.TableData
			repoURL := repo.GetCloneURL()

			desc := repo.GetDescription()
			if desc == "" {
				desc = "Sem descrição"
			}
			tableData = append(tableData, []string{pterm.LightCyan("Nome:"), pterm.White(repo.GetName())})
			tableData = append(tableData, []string{pterm.LightCyan("Descrição:"), pterm.White(desc)})

			lang := repo.GetLanguage()
			if lang == "" {
				lang = "-"
			}
			tableData = append(tableData, []string{pterm.LightCyan("Linguagem:"), pterm.LightMagenta(lang)})
			tableData = append(tableData, []string{pterm.LightCyan("Stars:"), pterm.LightYellow(strconv.Itoa(repo.GetStargazersCount()))})
			tableData = append(tableData, []string{pterm.LightCyan("Forks:"), pterm.LightCyan(strconv.Itoa(repo.GetForksCount()))})
			tableData = append(tableData, []string{pterm.LightCyan("URL:"), pterm.LightBlue(repo.GetHTMLURL())})

			tableStr, err := pterm.DefaultTable.WithData(tableData).Srender()
			if err == nil {
				pterm.DefaultBox.
					WithTitle(pterm.LightGreen(" Detalhes do Repositório ")).
					WithTitleTopCenter().
					Println(tableStr)
			} else {
				fmt.Println("Erro ao formatar os dados:", err)
			}

			cmd := exec.Command("git", "clone", repoURL)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()

			break
		}
	}
	return nil
}
