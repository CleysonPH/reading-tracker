package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID            int64
	Title         string
	Subtitle      sql.NullString
	Isbn          sql.NullString
	Authors       string
	Categories    string
	Language      string
	Cover         sql.NullString
	Publisher     sql.NullString
	PublishedAt   sql.NullTime
	Pages         int32
	ReadPages     int32
	Description   sql.NullString
	ReadingStatus string
	Edition       sql.NullInt32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type BookModel interface {
	All(q string) ([]*Book, error)
}

type bookMySQlModel struct {
	db *sql.DB
}

// All implements BookModel
func (m *bookMySQlModel) All(q string) ([]*Book, error) {
	stmt := `
		SELECT
			id,
			title,
			subtitle,
			isbn,
			authors,
			categories,
			language,
			cover,
			publisher,
			published_at,
			pages,
			read_pages,
			description,
			reading_status,
			edition,
			created_at,
			updated_at
		FROM
			books
		WHERE
			LOWER(title) LIKE CONCAT('%', LOWER(?), '%')
	`

	rows, err := m.db.Query(stmt, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []*Book{}
	for rows.Next() {
		b := &Book{}
		err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Subtitle,
			&b.Isbn,
			&b.Authors,
			&b.Categories,
			&b.Language,
			&b.Cover,
			&b.Publisher,
			&b.PublishedAt,
			&b.Pages,
			&b.ReadPages,
			&b.Description,
			&b.ReadingStatus,
			&b.Edition,
			&b.CreatedAt,
			&b.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func NewBookModel(db *sql.DB) BookModel {
	return &bookMySQlModel{db}
}
