package repository

import (
	todo_repo "example/todo-api/internal/repository/todo"
	"example/todo-api/pkg/database/postgres"
)

// Repository — агрегат репозиториев (как в Effective Mobile).
// Сейчас один todo; позже можно добавить другие сущности через With*.
type Repository struct {
	todo_repo.ITodoRepository
}

type repoOption func(r *Repository) error

func New(opts ...repoOption) (*Repository, error) {
	r := &Repository{}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}
	return r, nil
}

func WithTodoRepo(q postgres.Querier) repoOption {
	return func(r *Repository) error {
		r.ITodoRepository = todo_repo.New(q)
		return nil
	}
}
