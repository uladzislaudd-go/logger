package logger

import (
	"fmt"
	"sync"
)

type (
	stub struct{}
)

var (
	s = &stub{}

	formats    = []string{"", "%#v"}
	formats_mu sync.RWMutex
)

func init() {
	format_expand(20)
}

func Stub() Logger {
	return s
}

func (*stub) Log(l Level, m string, a ...any) {
	if l == Emerg {
		s.Emerg(m, a)
	}
}
func (*stub) Emerg(m string, a ...any) {
	fmt.Printf(m+": "+format(len(a)), a)
}
func (*stub) Alert(m string, a ...any)  {}
func (*stub) Crit(m string, a ...any)   {}
func (*stub) Error(m string, a ...any)  {}
func (*stub) Warn(m string, a ...any)   {}
func (*stub) Notice(m string, a ...any) {}
func (*stub) Info(m string, a ...any)   {}
func (*stub) Debug(m string, a ...any)  {}

func (*stub) SetLevel(level Level) {}
func (*stub) GetLevel() Level      { return Emerg }

func format_expand(n int) {
	formats_mu.Lock()
	defer formats_mu.Unlock()
	for i := len(formats); i < n; i++ {
		formats = append(formats, formats[i-1]+" %#v")
	}
}

func format(i int) string {
	formats_mu.RLock()
	defer formats_mu.RUnlock()

	if i > len(formats) {
		format_expand(i)
	}

	return formats[i]
}
