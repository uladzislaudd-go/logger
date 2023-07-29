package test

import (
	"bytes"
	"io"
	"os"
)

func Capture(file **os.File, f func()) string {
	old := *file
	r, w, _ := os.Pipe()
	defer r.Close()
	*file = w

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
	file = &old
	return <-outC
}

func CaptureStdout(f func()) string {
	return Capture(&os.Stdout, f)
}
