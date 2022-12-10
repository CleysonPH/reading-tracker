package repository

import (
	"database/sql"

	"github.com/CleysonPH/reading-tracker/internal/model"
)

func NewReadingSessionModel(db *sql.DB) ReadingSessionRepository {
	return &readingSessionRepository{db}
}

type readingSessionRepository struct {
	db *sql.DB
}

// Create implements ReadingSessionRepository
func (r *readingSessionRepository) Create(
	readingSession *model.ReadingSession,
) (*model.ReadingSession, error) {
	stmt := `
		INSERT INTO reading_sessions (
			read_pages,
			date,
			book_id
		) VALUES (?, ?, ?)
	`

	res, err := r.db.Exec(
		stmt,
		readingSession.ReadPages,
		readingSession.Date,
		readingSession.BookID,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return r.Get(id)
}

// Get implements ReadingSessionRepository
func (r *readingSessionRepository) Get(id int64) (*model.ReadingSession, error) {
	stmt := `
		SELECT
			id,
			read_pages,
			date,
			book_id,
			created_at,
			updated_at
		FROM
			reading_sessions
		WHERE
			id = ?
		LIMIT 1
	`

	row := r.db.QueryRow(stmt, id)

	rs := &model.ReadingSession{}
	err := row.Scan(
		&rs.ID,
		&rs.ReadPages,
		&rs.Date,
		&rs.BookID,
		&rs.CreatedAt,
		&rs.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrReadingSessionNotFound
		}
		return nil, err
	}

	return rs, nil
}
