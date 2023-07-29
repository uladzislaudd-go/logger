package logger

type (
	Field struct {
		Name      string
		FieldType FieldType
		Value     any
	}

	FieldType uint32
)

const (
	TypeAny FieldType = iota
	TypeString
)

func (f *Field) Format() string {
	switch f.FieldType {
	case TypeString:
		return "%s: \"%s\";"
	}

	return "%s: \"%#v\";"
}

func String(name, value string) *Field {
	rv := &Field{
		Name:      name,
		FieldType: TypeString,
		Value:     value,
	}

	return rv
}
