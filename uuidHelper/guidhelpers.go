package uuidHelper

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/geekmdio/noted/logHelper"
)

func GenerateGuidString() string {
	return generateGuidString(false)
}

// Do not use. This is only for testing purposes.
func XXX_GenerateGuidStringThrowsError() string {
	return generateGuidString(true)
}

func generateGuidString(throwError bool) string {
	uuid, err := uuid.NewUUID()
	if err != nil || !throwError {
		f := logHelper.GenerateLogFields("GenerateGuidString", "GUID", uuid.String(), err)
		log.WithFields(f).Error("Unable to generate a GUID.")
	}
	return uuid.String()
}