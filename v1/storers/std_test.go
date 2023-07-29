package storers

/*
var (
	l logger.Logger
)

func init() {
	var err error
	l, err = Logger()
	p(err)
}

func p(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_logger_Log(t *testing.T) {
	tests := []struct {
		l logger.Level
		m string
		a []any

		rv string
	}{
		{logger.Info, "q", []any{"w", "e", "r"}, "[Info] \"q\" \"w\" \"e\" \"r\"\n"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			assert.Equal(t, tt.rv, capture(func() {
				l.Log(tt.l, tt.m, tt.a...)
			}))
		})
	}
}
*/
