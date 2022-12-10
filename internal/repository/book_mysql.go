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

// UpdateReadPages implements BookRepository
func (r *bookMySQlRepository) UpdateReadPagesAndReadingStatus(bookID int64, readPages int32, readingStatus string) error {
	stmt := `
		UPDATE
			books
		SET
			read_pages = ?,
			reading_status = ?
		WHERE
			id = ?
	`

	_, err := r.db.Exec(stmt, readPages, readingStatus, bookID)
	return err
}

// Exists implements BookRepository
func (r *bookMySQlRepository) Exists(id int64) bool {
	stmt := `
		SELECT
			id
		FROM
			books
		WHERE
			id = ?
		LIMIT 1
	`

	var bookId int64
	err := r.db.QueryRow(stmt, id).Scan(&bookId)
	return err == nil && id > 0
}

// ExistsByIsbnAndIdNot implements BookRepository
func (r *bookMySQlRepository) ExistsByIsbnAndIdNot(isbn string, id int64) bool {
	stmt := `
		SELECT
			id
		FROM
			books
		WHERE
			isbn = ?
			AND id != ?
		LIMIT 1
	`

	var bookId int64
	err := r.db.QueryRow(stmt, isbn, id).Scan(&bookId)
	return err == nil && id > 0
}

// Update implements BookRepository
func (r *bookMySQlRepository) Update(book *model.Book) (*model.Book, error) {
	stmt := `
		UPDATE
			books
		SET
			title = ?,
			subtitle = ?,
			isbn = ?,
			authors = ?,
			categories = ?,
			language = ?,
			publisher = ?,
			published_at = ?,
			pages = ?,
			description = ?,
			edition = ?
		WHERE
			id = ?
	`

	_, err := r.db.Exec(
		stmt,
		book.Title,
		book.Subtitle,
		book.Isbn,
		book.Authors,
		book.Categories,
		book.Language,
		book.Publisher,
		book.PublishedAt,
		book.Pages,
		book.Description,
		book.Edition,
		book.ID,
	)
	if err != nil {
		return nil, err
	}

	return r.Get(book.ID)
}

// ExistsByIsbn implements BookRepository
func (r *bookMySQlRepository) ExistsByIsbn(isbn string) bool {
	stmt := `
		SELECT
			id
		FROM
			books
		WHERE
			isbn = ?
		LIMIT 1
	`

	var id int64
	err := r.db.QueryRow(stmt, isbn).Scan(&id)
	return err == nil && id > 0
}

// Create implements BookRepository
func (r *bookMySQlRepository) Create(book *model.Book) (*model.Book, error) {
	stmt := `
		INSERT INTO books (
			title,
			subtitle,
			isbn,
			authors,
			categories,
			language,
			publisher,
			published_at,
			pages,
			read_pages,
			description,
			reading_status,
			edition
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, 0, ?, 'to-read', ?)
	`

	res, err := r.db.Exec(
		stmt,
		book.Title,
		book.Subtitle,
		book.Isbn,
		book.Authors,
		book.Categories,
		book.Language,
		book.Publisher,
		book.PublishedAt,
		book.Pages,
		book.Description,
		book.Edition,
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

// Delete implements BookRepository
func (r *bookMySQlRepository) Delete(id int64) error {
	stmt := `
		DELETE FROM
			books
		WHERE
			id = ?
	`

	_, err := r.db.Exec(stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Get implements BookRepository
func (r *bookMySQlRepository) Get(id int64) (*model.Book, error) {
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
			id = ?
		LIMIT 1
	`

	row := r.db.QueryRow(stmt, id)

	book := &model.Book{}
	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Subtitle,
		&book.Isbn,
		&book.Authors,
		&book.Categories,
		&book.Language,
		&book.Cover,
		&book.Publisher,
		&book.PublishedAt,
		&book.Pages,
		&book.ReadPages,
		&book.Description,
		&book.ReadingStatus,
		&book.Edition,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrBookNotFound
		}
		return nil, err
	}

	return book, nil
}

// All implements BookModel
func (r *bookMySQlRepository) All(q string) ([]*model.Book, error) {
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

	rows, err := r.db.Query(stmt, q)
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
