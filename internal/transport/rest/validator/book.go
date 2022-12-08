package validator

import (
	"fmt"
	"time"

	"github.com/CleysonPH/reading-tracker/internal/repository"
	"github.com/CleysonPH/reading-tracker/internal/transport/rest/dto"
)

func NewBookValidator(bookRepository repository.BookRepository) BookValidator {
	return &bookValidator{bookRepository: bookRepository}
}

type bookValidator struct {
	bookRepository repository.BookRepository
}

// validateBook validates common fields for create and update
func (*bookValidator) validateBook(request *dto.BookRequest) *ValidationError {
	validationError := &ValidationError{}

	// validate title
	validationError.AddErrorIf(request.Title == "", "title", "is required")
	validationError.AddErrorIf(len(request.Title) < 3, "title", "must be at least 3 characters")
	validationError.AddErrorIf(len(request.Title) > 255, "title", "must be less than 255 characters")

	// validate subtitle
	if request.Subtitle.Valid() {
		validationError.AddErrorIf(len(request.Subtitle.Value) < 3, "subtitle", "must be at least 3 characters")
		validationError.AddErrorIf(len(request.Subtitle.Value) > 255, "subtitle", "must be less than 255 characters")
	}

	// validate isbn
	if request.Isbn.Valid() {
		has10Or13Digits := len(request.Isbn.Value) == 10 || len(request.Isbn.Value) == 13
		validationError.AddErrorIfNot(has10Or13Digits, "isbn", "must have 10 or 13 digits")
		for _, char := range request.Isbn.Value {
			if char < '0' || char > '9' {
				validationError.AddError("isbn", "must contain only digits")
				break
			}
		}
	}

	// validate authors
	validationError.AddErrorIf(len(request.Authors) == 0, "authors", "is required")
	for i, author := range request.Authors {
		validationError.AddErrorIf(len(author) < 3, fmt.Sprintf("authors[%d]", i), "must be at least 3 characters")
		validationError.AddErrorIf(len(author) > 255, fmt.Sprintf("authors[%d]", i), "must be less than 255 characters")
	}

	// validate categories
	validationError.AddErrorIf(len(request.Categories) == 0, "categories", "is required")
	for i, category := range request.Categories {
		validationError.AddErrorIf(len(category) < 3, fmt.Sprintf("categories[%d]", i), "must be at least 3 characters")
		validationError.AddErrorIf(len(category) > 255, fmt.Sprintf("categories[%d]", i), "must be less than 255 characters")
	}

	// validate language
	validationError.AddErrorIf(request.Language == "", "language", "is required")
	validationError.AddErrorIf(len(request.Language) != 2, "language", "must have 2 characters")
	validationError.AddErrorIfNot(stringIn(request.Language, "en", "pt"), "language", "must be 'en' or 'pt'")

	// validate publisher
	if request.Publisher.Valid() {
		validationError.AddErrorIf(len(request.Publisher.Value) < 3, "publisher", "must be at least 3 characters")
		validationError.AddErrorIf(len(request.Publisher.Value) > 255, "publisher", "must be less than 255 characters")
	}

	// validate publishedAt
	if request.PublishedAt.Valid() {
		validationError.AddErrorIf(request.PublishedAt.Value.After(time.Now()), "published_at", "must be in the past")
	}

	// validate pages
	validationError.AddErrorIf(request.Pages <= 0, "pages", "must be greater than 0")

	// validate description
	if request.Description.Valid() {
		validationError.AddErrorIf(len(request.Description.Value) < 3, "description", "must be at least 3 characters")
		validationError.AddErrorIf(len(request.Description.Value) > 1000, "description", "must be less than 1000 characters")
	}

	// validate edition
	if request.Edition.Valid() {
		validationError.AddErrorIf(request.Edition.Value <= 0, "edition", "must be greater than 0")
	}

	return validationError
}

// ValidateBookUpdate implements BookValidator
func (v *bookValidator) ValidateBookUpdate(id int64, request *dto.BookRequest) error {
	validationError := v.validateBook(request)

	// validate isbn
	if request.Isbn.Valid() {
		alreadyInUse := v.bookRepository.ExistsByIsbnAndIdNot(request.Isbn.Value, id)
		validationError.AddErrorIf(alreadyInUse, "isbn", "is already in use")
	}

	if validationError.HasErrors() {
		return validationError
	}

	return nil
}

// ValidateBookCreate implements BookValidator
func (v *bookValidator) ValidateBookCreate(request *dto.BookRequest) error {
	validationError := v.validateBook(request)

	// validate isbn
	if request.Isbn.Valid() {
		alreadyInUse := v.bookRepository.ExistsByIsbn(request.Isbn.Value)
		validationError.AddErrorIf(alreadyInUse, "isbn", "is already in use")
	}

	if validationError.HasErrors() {
		return validationError
	}

	return nil
}
