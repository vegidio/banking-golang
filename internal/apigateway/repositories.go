package apigateway

import (
	"banking/internal/apigateway/ent"
	"banking/internal/apigateway/feat/users"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func SetupRepositories() (*ent.Client, error) {
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUsername := os.Getenv("PG_USERNAME")
	pgPassword := os.Getenv("PG_PASSWORD")
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=bank sslmode=disable", pgHost,
		pgPort, pgUsername, pgPassword)

	client, err := ent.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	initRepos(client)
	return client, nil
}

func initRepos(client *ent.Client) {
	users.SetupRepository(client)
}
