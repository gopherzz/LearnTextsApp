package textwork

type TextLineIterator interface {
	NextLine() []string
}

type TextWordIterator interface {
	NextWord() (string, uint)
}
