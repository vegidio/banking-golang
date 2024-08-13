package main

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
	"template-golang/internal/ent"
)

func main() {
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUsername := os.Getenv("PG_USERNAME")
	pgPassword := os.Getenv("PG_PASSWORD")
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=countries sslmode=disable", pgHost,
		pgPort, pgUsername, pgPassword)

	client, err := ent.Open("postgres", dataSource)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()

	QueryUser(context.Background(), client)
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		First(ctx)

	if err != nil {
		slog.Error("failed querying user: %w", err)
	}

	log.Println("user returned: ", u)
	return u, nil
}
