package helpers

import(
	"github.com/google/uuid"
)

// IsValidUUID helper function takes in a string and confirms is a UUID and returns bool
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
