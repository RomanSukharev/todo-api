package main

import (
	"context"
	"fmt"
	"log"

	"example/todo-api/internal/config"
	"example/todo-api/internal/model"
	"example/todo-api/internal/repository"
	"example/todo-api/pkg/database/postgres"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	ctx := context.Background()

	db, err := postgres.New(ctx, cfg.Postgres.ConnStr)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	defer db.Close()

	repo, err := repository.New(repository.WithTodoRepo(db))
	if err != nil {
		log.Fatalf("repository: %v", err)
	}

	// Временная проверка слоя repository (потом заменим на HTTP).
	todo, err := repo.Create(ctx, model.CreateTodoInput{
		Title:       "первая задача",
		Description: "проверка repository",
	})
	if err != nil {
		log.Fatalf("create: %v", err)
	}

	fmt.Printf("todo-api ok | created id=%d title=%q on %s:%s\n",
		todo.ID, todo.Title, cfg.HTTP.Host, cfg.HTTP.Port)
}
