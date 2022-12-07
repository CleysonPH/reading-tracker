package validator

import "github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"

type BookValidator interface {
	ValidateBookCreate(request *dto.BookRequest) error
}
