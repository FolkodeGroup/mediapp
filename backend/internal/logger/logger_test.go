package logger

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	Init()
	if L() == nil {
		t.Error("Logger no fue inicializado correctamente")
	}
}
