package dao

import (
	"time"

	"example/todo-api/internal/model"
)

// TodoDAO — строка из БД. Domain/model не знает про sql.Null* и scan.
type TodoDAO struct {
	ID          int64
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *TodoDAO) ToModel() model.Todo {
	return model.Todo{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Completed:   t.Completed,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

type CreateTodoInputDAO struct {
	Title       string
	Description string
}

func (c *CreateTodoInputDAO) FromModel(in model.CreateTodoInput) {
	c.Title = in.Title
	c.Description = in.Description
}

type UpdateTodoInputDAO struct {
	Title       *string
	Description *string
	Completed   *bool
}

func (u *UpdateTodoInputDAO) FromModel(in model.UpdateTodoInput) {
	u.Title = in.Title
	u.Description = in.Description
	u.Completed = in.Completed
}
