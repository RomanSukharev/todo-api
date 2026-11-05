package model

import "time"

// Todo — основная сущность приложения.
// Эту структуру используют:
//   - handler  → JSON запрос/ответ (теги json)
//   - service  → бизнес-логика
//   - repository → чтение/запись в БД (через dao)
//
// Сюда НЕ кладём SQL, HTTP и валидацию с побочными эффектами —
// только данные и теги сериализации.
type Todo struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateTodoInput — данные для создания (как CreateUserInput в EM).
type CreateTodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateTodoInput — частичное обновление (nil = поле не трогаем).
type UpdateTodoInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
}
