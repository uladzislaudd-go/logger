package storers

import (
	"fmt"
	"os"

	"github.com/uladzislaudd-go/logger"
)

type (
	std struct{}
)

var (
	s = &std{}
)

func NewStdStorer() (logger.Storer, error) {
	return s, nil
}

func (s *std) Store(level logger.Level, message string) error {
	if level <= logger.Error {
		_, err := fmt.Fprintf(os.Stderr, "%s\n", message)
		return err
	} else {
		_, err := fmt.Fprintf(os.Stdout, "%s\n", message)
		return err
	}
}
