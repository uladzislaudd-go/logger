package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	l Logger
)

func init() {
	var err error
	l, err = New()
	p(err)
}

func p(err error) {
	if err != nil {
		panic(err)
	}
}

func capture(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	defer r.Close()
	os.Stdout = w

	f()

	outC := make(chan string)
	defer close(outC)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = old // restoring the real stdout
	return <-outC
}

func Test_logger_Log(t *testing.T) {
	tests := []struct {
		level Level
		a     []any
		rv    string
	}{
		{Info, []any{"q", "w", "e", "r"}, "[Info] \"q\" \"w\" \"e\" \"r\"\n"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tt.rv, capture(func() {
				l.Log(tt.level, tt.a...)
			}))
		})
	}
}
