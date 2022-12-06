package repository

import "github.com/CleysonPH/reading-tracker/internal/model"

type BookRepository interface {
	All(q string) ([]*model.Book, error)
	Get(id int64) (*model.Book, error)
	Delete(id int64) error
}
