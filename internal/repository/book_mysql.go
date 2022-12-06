package repository

import (
	"database/sql"

	"github.com/CleysonPH/reading-tracker/internal/model"
)

func NewBookModel(db *sql.DB) BookRepository {
	return &bookMySQlRepository{db}
}

type bookMySQlRepository struct {
	db *sql.DB
}

// All implements BookModel
func (m *bookMySQlRepository) All(q string) ([]*model.Book, error) {
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

	books := []*model.Book{}
	for rows.Next() {
		b := &model.Book{}
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
