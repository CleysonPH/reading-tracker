package dto

import (
	"encoding/json"
	"strings"
	"time"
)

type NullString struct {
	Value string
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if n.Value != "" {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

func (n *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &n.Value)
}

func (n NullString) Valid() bool {
	return n.Value != ""
}

type NullInt32 struct {
	Value int32
}

func (n NullInt32) Valid() bool {
	return n.Value != 0
}

func (n NullInt32) MarshalJSON() ([]byte, error) {
	if n.Value != 0 {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

func (n *NullInt32) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &n.Value)
}

type NullTime struct {
	Value time.Time
}

func (n NullTime) Valid() bool {
	return !n.Value.IsZero()
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Value.IsZero() {
		return json.Marshal(n.Value)
	}
	return []byte("null"), nil
}

func (n *NullTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	return json.Unmarshal(b, &n.Value)
}

type NullDate struct {
	Value time.Time
}

func (n NullDate) Valid() bool {
	return !n.Value.IsZero()
}

func (n NullDate) MarshalJSON() ([]byte, error) {
	if !n.Value.IsZero() {
		return json.Marshal(n.Value.Format("2006-01-02"))
	}
	return []byte("null"), nil
}

func (n *NullDate) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	date, err := time.Parse("2006-01-02", strings.Replace(string(b), "\"", "", -1))
	if err != nil {
		return err
	}
	n.Value = date
	return nil
}
