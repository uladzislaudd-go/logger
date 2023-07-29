package formaters

import (
	"fmt"

	"github.com/uladzislaudd-go/logger"
)

type (
	d struct{}
)

func NewDefaultFormater() (logger.Formater, error) {
	return &d{}, nil
}

func (d *d) Format(level logger.Level, message string, fields []*logger.Field) string {
	f := "[%s] %s:"
	v := []any{level.String(), message}
	for _, field := range fields {
		f += " " + field.Format()
		v = append(v, field.Name, field.Value)
	}

	return fmt.Sprintf(f, v...)
}
