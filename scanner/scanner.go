package scanner

type Scanner interface {
	ScanTokens() []string
}

func New(source string) Scanner {
	return &_Scanner{}
}

type _Scanner struct{}

func (s _Scanner) ScanTokens() []string {
	return make([]string, 0)
}
