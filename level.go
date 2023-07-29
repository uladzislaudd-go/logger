package logger

type (
	Level uint32
)

const (
	Emerg Level = iota
	Alert
	Crit
	Error
	Warn
	Notice
	Info
	Debug
)

var (
	Levels = []Level{Emerg, Alert, Crit, Error, Warn, Notice, Info, Debug}

	levelStrings = map[Level]string{
		Emerg:  "Emerg",
		Alert:  "Alert",
		Crit:   "Crit",
		Error:  "Error",
		Warn:   "Warn",
		Notice: "Notice",
		Info:   "Info",
		Debug:  "Debug",
	}
)

func (l Level) Valid() (rv bool) {
	_, rv = levelStrings[l]
	return
}

func (l Level) String() string {
	if v, ok := levelStrings[l]; ok {
		return v
	}

	return "Invalid"
}
