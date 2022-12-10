package repository

import "github.com/CleysonPH/reading-tracker/internal/model"

type BookRepository interface {
	All(q string) ([]*model.Book, error)
	Get(id int64) (*model.Book, error)
	Delete(id int64) error
	Create(book *model.Book) (*model.Book, error)
	Exists(id int64) bool
	ExistsByIsbn(isbn string) bool
	ExistsByIsbnAndIdNot(isbn string, id int64) bool
	Update(book *model.Book) (*model.Book, error)
	UpdateReadPagesAndReadingStatus(bookID int64, readPages int32, readingStatus string) error
}

type ReadingSessionRepository interface {
	Get(id int64) (*model.ReadingSession, error)
	Create(readingSession *model.ReadingSession) (*model.ReadingSession, error)
	Delete(id int64) error
}
