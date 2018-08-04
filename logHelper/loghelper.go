package logHelper

import log "github.com/sirupsen/logrus"

func GenerateLogFields(function string, expected string, got string, err error) log.Fields {
	 return log.Fields {
		"function": function,
		"expected": expected,
		"got": got,
		"error": err,
	}
}