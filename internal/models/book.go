package models

import (
	"database/sql"
	"time"
)

type Book struct {
	ID            int64
	Title         string
	Subtitle      sql.NullString
	Isbn          sql.NullString
	Authors       string
	Categories    string
	Language      string
	Cover         sql.NullString
	Publisher     sql.NullString
	PublishedAt   sql.NullTime
	Pages         int32
	ReadPages     int32
	Description   sql.NullString
	ReadingStatus string
	Edition       sql.NullInt32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
