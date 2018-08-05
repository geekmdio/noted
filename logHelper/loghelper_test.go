package logHelper

import (
	"testing"
)

func TestGenerateLogFields_ReturnsObject(t *testing.T) {
	lf := GenerateLogFields("function", "expected", "got", nil)
	if lf == nil {
		t.Errorf("Expected an object, got %v", lf)
	}
}