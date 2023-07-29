package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/uladzislaudd-go/logger"
	"github.com/uladzislaudd-go/logger/test"
)

func Test_l_log(t *testing.T) {
	l := NewDefaultLogger(logger.Debug)

	tests := []struct {
		level   logger.Level
		message string
		anys    []any

		rv string
	}{
		{logger.Info, "q", []any{"w", "e", "r"}, "[Info] q: w: \"e\"; r: \"[!NO VALUE!]\";\n"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			rv := test.CaptureStdout(func() { l.Log(tt.level, tt.message, tt.anys...) })
			assert.Equal(t, tt.rv, rv)
		})
	}
}
