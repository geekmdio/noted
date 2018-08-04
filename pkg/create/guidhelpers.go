package create

import (
	"github.com/google/uuid"
	"log"
)

func GenerateGuidString() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Error creating new GUID: %v", err)
	}
	uuidString := uuid.String()
	return uuidString
}