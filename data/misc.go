package data

import uuid "github.com/satori/go.uuid"

// NewUUID generates a new UUID.
func NewUUID() string {
	return uuid.Must(uuid.NewV4()).String()
}
