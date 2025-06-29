package repositories

import (
	"capsuler/pkg/capsule"
	"database/sql"
	"time"
)

type SqlCapsuleRepository struct {
	pool *sql.DB
}

func NewSqlCapsuleRepository(pool *sql.DB) *SqlCapsuleRepository {
	return &SqlCapsuleRepository{
		pool: pool,
	}
}

func (r *SqlCapsuleRepository) Save(capsule *capsule.Capsule) error {
	if _, err := r.pool.Exec(`
		INSERT INTO capsules (id, name, description, date_to_open, is_open, created_at, updated_at)
		VALUES (@id, @name, @description, @date_to_open, @is_open, @created_at, @updated_at)
		ON CONFLICT(id) DO UPDATE SET
		name = excluded.name,
		description = excluded.description,
		date_to_open = excluded.date_to_open,
		is_open = excluded.is_open,
		updated_at = excluded.updated_at
	`,
		sql.Named("id", capsule.GetId()),
		sql.Named("name", capsule.GetName()),
		sql.Named("description", capsule.GetDescription()),
		sql.Named("date_to_open", capsule.GetDateToOpen()),
		sql.Named("is_open", capsule.WasOpened()),
		sql.Named("created_at", capsule.GetCreatedAt()),
		sql.Named("updated_at", capsule.GetUpdatedAt()),
	); err != nil {
		return err
	}

	if len(capsule.GetMessages()) > 0 {
		message := capsule.GetMessages()[len(capsule.GetMessages())-1]

		if _, err := r.pool.Exec(`
			INSERT INTO messages (capsule_id, message, created_at)
			VALUES (@capsule_id, @message, @created_at)
		`,
			sql.Named("capsule_id", capsule.GetId()),
			sql.Named("message", message),
			sql.Named("created_at", time.Now())); err != nil {
			return err
		}
	}

	return nil
}

func (r *SqlCapsuleRepository) GetById(id string) (*capsule.Capsule, error) {
	var name string
	var description string
	var dateToOpen time.Time
	var isOpen bool
	var createdAt time.Time
	var updatedAt time.Time
	var messages []string

	if err := r.pool.QueryRow(`
	SELECT name, description, date_to_open, is_open, created_at, updated_at
	FROM capsules
	WHERE id = @id
	`, sql.Named("id", id)).Scan(&name, &description, &dateToOpen, &isOpen, &createdAt, &updatedAt); err != nil {
		return nil, err
	}

	rows, err := r.pool.Query(`
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

	capsule := capsule.Builder().
		WithId(id).
		WithName(name).
		WithIsOpen(isOpen).
		WithDescription(description).
		WithDateToOpen(dateToOpen).
		WithCreatedAt(createdAt).
		WithUpdatedAt(updatedAt).
		WithMessages(messages).
		Build()

	return capsule, nil
}
