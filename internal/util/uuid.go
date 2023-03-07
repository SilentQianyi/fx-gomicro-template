package util

import (
	"github.com/google/uuid"
	"strings"
)

func GetUUID() string {
	id, _ := uuid.NewUUID()
	return strings.ToLower(strings.ReplaceAll(id.String(), "-", ""))
}
