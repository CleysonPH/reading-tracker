package dto

import (
	"encoding/json"
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

type NullInt32 struct {
	Value int32
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
