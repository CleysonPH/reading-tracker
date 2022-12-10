package dto

import (
	"github.com/CleysonPH/reading-tracker/internal/model"
)

type ReadingSessionRequest struct {
	ReadPages int32 `json:"read_pages"`
	Date      Date  `json:"date"`
	BookID    int64 `json:"-"`
}

func (r *ReadingSessionRequest) ToReadingSession() *model.ReadingSession {
	return &model.ReadingSession{
		BookID:    r.BookID,
		ReadPages: r.ReadPages,
		Date:      r.Date.Value,
	}
}

type ReadingSessionResponse struct {
	ID        int64 `json:"id"`
	ReadPages int32 `json:"read_pages"`
	Date      Date  `json:"date"`
	BookID    int64 `json:"book_id"`
	CreatedAt Date  `json:"created_at"`
	UpdatedAt Date  `json:"updated_at"`
}

func (r *ReadingSessionResponse) FromReadingSession(readingSession *model.ReadingSession) {
	r.ID = readingSession.ID
	r.ReadPages = readingSession.ReadPages
	r.Date = Date{Value: readingSession.Date}
	r.BookID = readingSession.BookID
	r.CreatedAt = Date{Value: readingSession.CreatedAt}
	r.UpdatedAt = Date{Value: readingSession.UpdatedAt}
}
