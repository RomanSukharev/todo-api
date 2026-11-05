package queries

import (
	"time"

	"example/todo-api/internal/repository/dao"
	sq "github.com/Masterminds/squirrel"
)

// Как в Effective Mobile: SQL собираем squirrel'ом, не конкатенацией строк.

func BuildCreateTodoQuery(in dao.CreateTodoInputDAO) (string, []interface{}, error) {
	return sq.Insert("todos").
		SetMap(map[string]interface{}{
			"title":       in.Title,
			"description": in.Description,
			"completed":   false,
		}).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, title, description, completed, created_at, updated_at").
		ToSql()
}

func BuildGetTodoByIDQuery(id int64) (string, []interface{}, error) {
	return sq.Select("id", "title", "description", "completed", "created_at", "updated_at").
		From("todos").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func BuildListTodosQuery() (string, []interface{}, error) {
	return sq.Select("id", "title", "description", "completed", "created_at", "updated_at").
		From("todos").
		OrderBy("id ASC").
		PlaceholderFormat(sq.Dollar).
		ToSql()
}

func BuildUpdateTodoQuery(id int64, in dao.UpdateTodoInputDAO) (string, []interface{}, error) {
	builder := sq.Update("todos").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Set("updated_at", time.Now())

	if in.Title != nil {
		builder = builder.Set("title", *in.Title)
	}
	if in.Description != nil {
		builder = builder.Set("description", *in.Description)
	}
	if in.Completed != nil {
		builder = builder.Set("completed", *in.Completed)
	}

	return builder.
		Suffix("RETURNING id, title, description, completed, created_at, updated_at").
		ToSql()
}

func BuildDeleteTodoQuery(id int64) (string, []interface{}, error) {
	return sq.Delete("todos").
		Where(sq.Eq{"id": id}).
		PlaceholderFormat(sq.Dollar).
		Suffix("RETURNING id, title, description, completed, created_at, updated_at").
		ToSql()
}
