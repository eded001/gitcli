# Git CLI

CLI construída em **Go** para experimentar automação e interação com Git/GitHub diretamente pelo terminal.

## Stack

- Go
- Cobra
- go-github
- PTerm

## Estrutura

```text
.
├── cmd/
├── internal/
├── main.go
├── go.mod
└── go.sum
```

- `cmd/`: comandos e composição da CLI;
- `internal/`: lógica interna e integrações;
- `main.go`: ponto de entrada.

## Executando

Pré-requisito: Go instalado.

```bash
git clone https://github.com/eded001/gitcli.git
cd gitcli
go mod download
go run .
```

Para gerar um binário:

```bash
go build -o gitcli .
```

Depois:

```bash
./gitcli
```

## Dependências principais

- **Cobra** para estrutura de comandos;
- **go-github** para integração com a API do GitHub;
- **PTerm** para uma experiência de terminal mais rica.

## Objetivo

Este projeto funciona como laboratório para desenvolvimento de ferramentas de linha de comando em Go, integração com APIs e organização modular de aplicações CLI.
