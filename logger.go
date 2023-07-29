package logger

type (
	Logger interface {
		Log(level Level, message string, a ...any)
		Emerg(message string, a ...any)
		Alert(message string, a ...any)
		Crit(message string, a ...any)
		Error(message string, a ...any)
		Warn(message string, a ...any)
		Notice(message string, a ...any)
		Info(message string, a ...any)
		Debug(message string, a ...any)

		SetLevel(level Level)
		GetLevel() Level
	}

	Formater interface {
		Format(level Level, message string, fields []*Field) string
	}

	Storer interface {
		Store(level Level, message string) error
	}
)
