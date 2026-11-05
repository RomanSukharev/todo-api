package todo_repo

import (
	"context"
	"errors"

	"example/todo-api/internal/errs"
	"example/todo-api/internal/model"
	"example/todo-api/internal/repository/dao"
	"example/todo-api/internal/repository/queries"
	"example/todo-api/pkg/database/postgres"

	"github.com/jackc/pgx/v5"
)

// ITodoRepository — контракт хранилища (как IUserRepository в EM).
type ITodoRepository interface {
	Create(ctx context.Context, input model.CreateTodoInput) (model.Todo, error)
	GetByID(ctx context.Context, id int64) (model.Todo, error)
	List(ctx context.Context) ([]model.Todo, error)
	Update(ctx context.Context, id int64, input model.UpdateTodoInput) (model.Todo, error)
	Delete(ctx context.Context, id int64) (model.Todo, error)
}

type todoRepo struct {
	q postgres.Querier
}

func New(q postgres.Querier) ITodoRepository {
	return &todoRepo{q: q}
}

func (r *todoRepo) Create(ctx context.Context, input model.CreateTodoInput) (model.Todo, error) {
	in := new(dao.CreateTodoInputDAO)
	in.FromModel(input)

	query, args, err := queries.BuildCreateTodoQuery(*in)
	if err != nil {
		return model.Todo{}, err
	}

	var row dao.TodoDAO
	err = r.q.QueryRow(ctx, query, args...).Scan(
		&row.ID,
		&row.Title,
		&row.Description,
		&row.Completed,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		return model.Todo{}, err
	}
	return row.ToModel(), nil
}

func (r *todoRepo) GetByID(ctx context.Context, id int64) (model.Todo, error) {
	query, args, err := queries.BuildGetTodoByIDQuery(id)
	if err != nil {
		return model.Todo{}, err
	}

	var row dao.TodoDAO
	err = r.q.QueryRow(ctx, query, args...).Scan(
		&row.ID,
		&row.Title,
		&row.Description,
		&row.Completed,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Todo{}, errs.ErrTodoNotFound
		}
		return model.Todo{}, err
	}
	return row.ToModel(), nil
}

func (r *todoRepo) List(ctx context.Context) ([]model.Todo, error) {
	query, args, err := queries.BuildListTodosQuery()
	if err != nil {
		return nil, err
	}

	rows, err := r.q.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var row dao.TodoDAO
		if err := rows.Scan(
			&row.ID,
			&row.Title,
			&row.Description,
			&row.Completed,
			&row.CreatedAt,
			&row.UpdatedAt,
		); err != nil {
			return nil, err
		}
		todos = append(todos, row.ToModel())
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todoRepo) Update(ctx context.Context, id int64, input model.UpdateTodoInput) (model.Todo, error) {
	in := new(dao.UpdateTodoInputDAO)
	in.FromModel(input)

	query, args, err := queries.BuildUpdateTodoQuery(id, *in)
	if err != nil {
		return model.Todo{}, err
	}

	var row dao.TodoDAO
	err = r.q.QueryRow(ctx, query, args...).Scan(
		&row.ID,
		&row.Title,
		&row.Description,
		&row.Completed,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Todo{}, errs.ErrTodoNotFound
		}
		return model.Todo{}, err
	}
	return row.ToModel(), nil
}

func (r *todoRepo) Delete(ctx context.Context, id int64) (model.Todo, error) {
	query, args, err := queries.BuildDeleteTodoQuery(id)
	if err != nil {
		return model.Todo{}, err
	}

	var row dao.TodoDAO
	err = r.q.QueryRow(ctx, query, args...).Scan(
		&row.ID,
		&row.Title,
		&row.Description,
		&row.Completed,
		&row.CreatedAt,
		&row.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Todo{}, errs.ErrTodoNotFound
		}
		return model.Todo{}, err
	}
	return row.ToModel(), nil
}
