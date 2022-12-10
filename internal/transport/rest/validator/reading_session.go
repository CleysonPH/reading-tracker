package validator

import (
	"time"

	"github.com/CleysonPH/reading-tracker/internal/repository"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"
)

func NewReadingSessionValidator(
	bookRepository repository.BookRepository,
	readingSessionRepository repository.ReadingSessionRepository,
) ReadingSessionValidator {
	return &readingSessionValidator{
		bookRepository:           bookRepository,
		readingSessionRepository: readingSessionRepository,
	}
}

type readingSessionValidator struct {
	bookRepository           repository.BookRepository
	readingSessionRepository repository.ReadingSessionRepository
}

// ValidateReadingSessionCreate implements ReadingSessionValidator
func (v *readingSessionValidator) ValidateReadingSessionCreate(request *dto.ReadingSessionRequest) error {
	validationError := &ValidationError{}

	// Validate ReadPages
	validationError.AddErrorIf(request.ReadPages <= 0, "read_pages", "must be greater than 0")
	book, _ := v.bookRepository.Get(request.BookID)
	totalReadPages, _ := v.readingSessionRepository.GetTotalReadPagesByBookID(request.BookID)
	validationError.AddErrorIf(totalReadPages+request.ReadPages > book.Pages, "read_pages", "must not exceed the total number of pages of the book")

	// Validate Date
	validationError.AddErrorIf(request.Date.Value.IsZero(), "date", "must be a valid date")
	validationError.AddErrorIf(time.Now().Before(request.Date.Value), "date", "must be a date in the past")

	if validationError.HasErrors() {
		return validationError
	}

	return nil
}
