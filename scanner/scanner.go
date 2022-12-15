package scanner

func New(source string) *Scanner {
	return &Scanner{}
}

type Scanner struct{}

func (s Scanner) ScanTokens() []string {
	return make([]string, 0)
}
