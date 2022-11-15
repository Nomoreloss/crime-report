package id

import (
	"github.com/twinj/uuid"
)

// GenerateNewUniqueCode for struct
func GenerateNewUniqueCode() string {
	return uuid.NewV4().String()
}
