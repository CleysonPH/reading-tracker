package validator

import "github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"

type BookValidator interface {
	ValidateBookCreate(request *dto.BookRequest) error
	ValidateBookUpdate(id int64, request *dto.BookRequest) error
}

type ReadingSessionValidator interface {
	ValidateReadingSessionCreate(request *dto.ReadingSessionRequest) error
}
