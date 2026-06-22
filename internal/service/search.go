package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v88/github"
)

func SearchUser(usernameOrID string) *github.User {
	client, err := github.NewClient()
	if err != nil {
		fmt.Println("Erro ao criar cliente:", err)
		return nil
	}

	var user *github.User
	if id, err := strconv.ParseInt(usernameOrID, 10, 64); err == nil {
		// É um número, buscar por ID
		user, _, err = client.Users.GetByID(context.Background(), id)
		if err != nil {
			fmt.Println("Erro ao buscar usuário por ID:", err)
			return nil
		}
	} else {
		// Não é um número, buscar por Username
		user, _, err = client.Users.Get(context.Background(), usernameOrID)
		if err != nil {
			fmt.Println("Erro ao buscar usuário por username:", err)
			return nil
		}
	}

	return user
}
