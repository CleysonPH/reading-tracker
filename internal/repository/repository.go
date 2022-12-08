package repository

import "github.com/CleysonPH/reading-tracker/internal/model"

type BookRepository interface {
	All(q string) ([]*model.Book, error)
	Get(id int64) (*model.Book, error)
	Delete(id int64) error
	Create(book *model.Book) (int64, error)
	Exists(id int64) bool
	ExistsByIsbn(isbn string) bool
	ExistsByIsbnAndIdNot(isbn string, id int64) bool
	Update(book *model.Book) (*model.Book, error)
}
