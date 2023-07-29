package v1

import (
	"github.com/uladzislaudd-go/logger"
	"github.com/uladzislaudd-go/logger/v1/formaters"
	"github.com/uladzislaudd-go/logger/v1/storers"
)

type (
	l struct {
		level logger.Level

		formater logger.Formater
		storers  []logger.Storer
	}
)

func NewLogger(level logger.Level, formater logger.Formater, storers ...logger.Storer) logger.Logger {
	rv := &l{
		level:    level,
		formater: formater,
		storers:  storers,
	}

	return rv
}

func NewDefaultLogger(level logger.Level) logger.Logger {
	f, _ := formaters.NewDefaultFormater()
	s, _ := storers.NewStdStorer()
	return NewLogger(level, f, s)
}

func (l *l) log(level logger.Level, message string, anys []any) {
	if l.level < level {
		return
	}

	fields := make([]*logger.Field, 0, len(anys))
	var name *string

	for _, a := range anys {
		switch v := a.(type) {

		case logger.Field:
			fields = append(fields, &v)
		case *logger.Field:
			fields = append(fields, v)

		case string:
			if name == nil {
				name = &v
			} else {
				fields = append(fields, logger.String(*name, v))
				name = nil
			}
		case *string:
			if name == nil {
				name = v
			} else {
				fields = append(fields, logger.String(*name, *v))
				name = nil
			}
		}
	}

	if name != nil {
		fields = append(fields, logger.String(*name, "[!NO VALUE!]"))
	}

	log := l.formater.Format(level, message, fields)
	for _, storer := range l.storers {
		storer.Store(level, log)
	}
}

func (l *l) Log(level logger.Level, message string, a ...any) {
	l.log(level, message, a)
}

func (l *l) Emerg(message string, a ...any) {
	l.log(logger.Emerg, message, a)
}

func (l *l) Alert(message string, a ...any) {
	l.log(logger.Alert, message, a)
}

func (l *l) Crit(message string, a ...any) {
	l.log(logger.Crit, message, a)
}

func (l *l) Error(message string, a ...any) {
	l.log(logger.Error, message, a)
}

func (l *l) Warn(message string, a ...any) {
	l.log(logger.Warn, message, a)
}

func (l *l) Notice(message string, a ...any) {
	l.log(logger.Notice, message, a)
}

func (l *l) Info(message string, a ...any) {
	l.log(logger.Info, message, a)
}

func (l *l) Debug(message string, a ...any) {
	l.log(logger.Debug, message, a)
}

func (l *l) SetLevel(level logger.Level) {
	l.level = level
}

func (l *l) GetLevel() logger.Level {
	return l.level
}
