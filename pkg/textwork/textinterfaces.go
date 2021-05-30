package textwork

// Iterator Interface for lines([][]string)
type TextLineIterator interface {
	NextLine() []string
}

// Iterator Interface for words([]string)
type TextWordIterator interface {
	NextWord() (string, uint)
}
