package logger

import (
	"fmt"
)

type (
	Logger interface {
		Log(level Level, a ...any)

		Emerg(a ...any)
		Alert(a ...any)
		Crit(a ...any)
		Error(a ...any)
		Warning(a ...any)
		Notice(a ...any)
		Info(a ...any)
		Debug(a ...any)

		Logf(level Level, format string, a ...any)

		Emergf(format string, a ...any)
		Alertf(format string, a ...any)
		Critf(format string, a ...any)
		Errorf(format string, a ...any)
		Warningf(format string, a ...any)
		Noticef(format string, a ...any)
		Infof(format string, a ...any)
		Debugf(format string, a ...any)
	}

	logger struct{}
)

func New() (Logger, error) {
	return &logger{}, nil
}

func (l *logger) log(level Level, format string, a ...any) {
	msg := fmt.Sprintf("[%s] "+format, level.String(), a)
	fmt.Printf("%s\n", msg)
}

func (l *logger) Log(level Level, a ...any) {
	l.log(level, format(len(a)), a...)
}

func (l *logger) Emerg(a ...any) {
	l.log(Emerg, format(len(a)), a...)
}

func (l *logger) Alert(a ...any) {
	l.log(Alert, format(len(a)), a...)
}

func (l *logger) Crit(a ...any) {
	l.log(Crit, format(len(a)), a...)
}

func (l *logger) Error(a ...any) {
	l.log(Error, format(len(a)), a...)
}

func (l *logger) Warning(a ...any) {
	l.log(Warning, format(len(a)), a...)
}

func (l *logger) Notice(a ...any) {
	l.log(Notice, format(len(a)), a...)
}

func (l *logger) Info(a ...any) {
	l.log(Info, format(len(a)), a...)
}

func (l *logger) Debug(a ...any) {
	l.log(Debug, format(len(a)), a...)
}

func (l *logger) Logf(level Level, format string, a ...any) {
	l.log(level, format, a...)
}

func (l *logger) Emergf(format string, a ...any) {
	l.log(Emerg, format, a...)
}

func (l *logger) Alertf(format string, a ...any) {
	l.log(Alert, format, a...)
}

func (l *logger) Critf(format string, a ...any) {
	l.log(Crit, format, a...)
}

func (l *logger) Errorf(format string, a ...any) {
	l.log(Error, format, a...)
}

func (l *logger) Warningf(format string, a ...any) {
	l.log(Warning, format, a...)
}

func (l *logger) Noticef(format string, a ...any) {
	l.log(Notice, format, a...)
}

func (l *logger) Infof(format string, a ...any) {
	l.log(Info, format, a...)
}

func (l *logger) Debugf(format string, a ...any) {
	l.log(Debug, format, a...)
}
