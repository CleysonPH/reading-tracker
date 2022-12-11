package repository

import "github.com/CleysonPH/reading-tracker/internal/model"

type BookRepository interface {
	Exists(id int64) bool
	Delete(id int64) error
	ExistsByIsbn(isbn string) bool
	Get(id int64) (*model.Book, error)
	All(q string) ([]*model.Book, error)
	Create(book *model.Book) (*model.Book, error)
	Update(book *model.Book) (*model.Book, error)
	ExistsByIsbnAndIdNot(isbn string, id int64) bool
	UpdateReadPagesAndReadingStatus(bookID int64, readPages int32, readingStatus string) error
}

type ReadingSessionRepository interface {
	Delete(id int64) error
	Get(id int64) (*model.ReadingSession, error)
	AllByBookID(bookID int64) ([]*model.ReadingSession, error)
	Create(readingSession *model.ReadingSession) (*model.ReadingSession, error)
}
