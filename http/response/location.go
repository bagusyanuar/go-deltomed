package response

import "github.com/google/uuid"

type Location struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
