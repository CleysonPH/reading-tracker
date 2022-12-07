package dto

import (
	"database/sql"
	"strings"
	"time"

	"github.com/CleysonPH/reading-tracker/internal/model"
)

type BookResponse struct {
	ID            int64      `json:"id"`
	Title         string     `json:"title"`
	Subtitle      NullString `json:"subtitle"`
	Isbn          NullString `json:"isbn"`
	Authors       []string   `json:"authors"`
	Categories    []string   `json:"categories"`
	Language      string     `json:"language"`
	Cover         NullString `json:"cover"`
	Publisher     NullString `json:"publisher"`
	PublishedAt   NullDate   `json:"published_at"`
	Pages         int32      `json:"pages"`
	ReadPages     int32      `json:"read_pages"`
	Description   NullString `json:"description"`
	ReadingStatus string     `json:"reading_status"`
	Edition       NullInt32  `json:"edition"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (b *BookResponse) FromBook(book *model.Book) {
	b.ID = book.ID
	b.Title = book.Title
	b.Subtitle = NullString{book.Subtitle.String}
	b.Isbn = NullString{book.Isbn.String}
	b.Authors = strings.Split(book.Authors, ",")
	b.Categories = strings.Split(book.Categories, ",")
	b.Language = book.Language
	b.Cover = NullString{book.Cover.String}
	b.Publisher = NullString{book.Publisher.String}
	b.PublishedAt = NullDate{book.PublishedAt.Time}
	b.Pages = book.Pages
	b.ReadPages = book.ReadPages
	b.Description = NullString{book.Description.String}
	b.ReadingStatus = book.ReadingStatus
	b.Edition = NullInt32{book.Edition.Int32}
	b.CreatedAt = book.CreatedAt
	b.UpdatedAt = book.UpdatedAt
}

type BookSummaryResponse struct {
	ID            int64      `json:"id"`
	Title         string     `json:"title"`
	Subtitle      NullString `json:"subtitle"`
	Authors       []string   `json:"authors"`
	Cover         NullString `json:"cover"`
	Pages         int32      `json:"pages"`
	ReadPages     int32      `json:"read_pages"`
	ReadingStatus string     `json:"reading_status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

func (b *BookSummaryResponse) FromBook(book *model.Book) {
	b.ID = book.ID
	b.Title = book.Title
	b.Subtitle = NullString{book.Subtitle.String}
	b.Authors = strings.Split(book.Authors, ",")
	b.Cover = NullString{book.Cover.String}
	b.Pages = book.Pages
	b.ReadPages = book.ReadPages
	b.ReadingStatus = book.ReadingStatus
	b.CreatedAt = book.CreatedAt
	b.UpdatedAt = book.UpdatedAt
}

type BookRequest struct {
	Title       string     `json:"title"`
	Subtitle    NullString `json:"subtitle"`
	Isbn        NullString `json:"isbn"`
	Authors     []string   `json:"authors"`
	Categories  []string   `json:"categories"`
	Language    string     `json:"language"`
	Publisher   NullString `json:"publisher"`
	PublishedAt NullDate   `json:"published_at"`
	Pages       int32      `json:"pages"`
	Description NullString `json:"description"`
	Edition     NullInt32  `json:"edition"`
}

func (b *BookRequest) ToBook() *model.Book {
	return &model.Book{
		Title: b.Title,
		Subtitle: sql.NullString{
			String: b.Subtitle.Value,
			Valid:  b.Subtitle.Valid(),
		},
		Isbn: sql.NullString{
			String: b.Isbn.Value,
			Valid:  b.Isbn.Valid(),
		},
		Authors:    strings.Join(b.Authors, ","),
		Categories: strings.Join(b.Categories, ","),
		Language:   b.Language,
		Publisher: sql.NullString{
			String: b.Publisher.Value,
			Valid:  b.Publisher.Valid(),
		},
		PublishedAt: sql.NullTime{
			Time:  b.PublishedAt.Value,
			Valid: b.PublishedAt.Valid(),
		},
		Pages: b.Pages,
		Description: sql.NullString{
			String: b.Description.Value,
			Valid:  b.Description.Valid(),
		},
		Edition: sql.NullInt32{
			Int32: b.Edition.Value,
			Valid: b.Edition.Valid(),
		},
	}
}
