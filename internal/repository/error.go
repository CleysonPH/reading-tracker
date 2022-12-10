package repository

import "errors"

var ErrBookNotFound = errors.New("book not found")
var ErrReadingSessionNotFound = errors.New("reading session not found")
