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
func generateGuidStringThrowsError() string {
	return generateGuidString(true)
}

func generateGuidString(throwError bool) string {
	uuid, err := uuid.NewUUID()
	if throwError {
		f := logHelper.GenerateLogFields("GenerateGuidString", "GUID", "Directive to throw error.", err)
		log.WithFields(f).Error("This is an intentional error that should only arise in testing cases.")
		return ""
	}
	if err != nil {
		f := logHelper.GenerateLogFields("GenerateGuidString", "GUID", uuid.String(), err)
		log.WithFields(f).Error("Unable to generate a GUID.")
		return ""
	}
	return uuid.String()
}