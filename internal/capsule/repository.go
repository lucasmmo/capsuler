package capsule

import (
	"database/sql"
	"time"
)

type Repository interface {
	Save(*Entity) error
	GetById(id string) (*Entity, error)
}

type repository struct {
	sql *sql.DB
}

func NewRepository(pool *sql.DB) *repository {
	return &repository{
		sql: pool,
	}
}

func (r *repository) Save(capsule *Entity) error {
	if _, err := r.sql.Exec(`
		INSERT INTO capsules (id, name, description, date_to_open, is_open, created_at, updated_at)
		VALUES (@id, @name, @description, @date_to_open, @is_open, @created_at, @updated_at)
		ON CONFLICT(id) DO UPDATE SET
		name = excluded.name,
		description = excluded.description,
		date_to_open = excluded.date_to_open,
		is_open = excluded.is_open,
		updated_at = excluded.updated_at
	`,
		sql.Named("id", capsule.Id),
		sql.Named("name", capsule.Name),
		sql.Named("description", capsule.Description),
		sql.Named("date_to_open", capsule.DateToOpen),
		sql.Named("is_open", capsule.IsOpen),
		sql.Named("created_at", capsule.CreatedAt),
		sql.Named("updated_at", capsule.UpdatedAt),
	); err != nil {
		return err
	}

	if len(capsule.Messages) > 0 {
		message := capsule.Messages[len(capsule.Messages)-1]

		if _, err := r.sql.Exec(`
			INSERT INTO messages (capsule_id, message, created_at)
			VALUES (@capsule_id, @message, @created_at)
		`,
			sql.Named("capsule_id", capsule.Id),
			sql.Named("message", message),
			sql.Named("created_at", time.Now())); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) GetById(id string) (*Entity, error) {
	var name string
	var description string
	var dateToOpen time.Time
	var isOpen bool
	var createdAt time.Time
	var updatedAt time.Time
	var messages []string

	if err := r.sql.QueryRow(`
	SELECT name, description, date_to_open, is_open, created_at, updated_at
	FROM capsules
	WHERE id = @id
	`, sql.Named("id", id)).Scan(&name, &description, &dateToOpen, &isOpen, &createdAt, &updatedAt); err != nil {
		return nil, err
	}

	rows, err := r.sql.Query(`
		SELECT message
		FROM messages
		WHERE capsule_id = @capsule_id
	`, sql.Named("capsule_id", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &Entity{
		Id:          id,
		Name:        name,
		Description: description,
		DateToOpen:  dateToOpen,
		IsOpen:      isOpen,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		Messages:    messages,
	}, nil
}
