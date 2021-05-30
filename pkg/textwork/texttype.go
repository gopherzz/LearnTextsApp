package textwork

import (
	"errors"
	"strings"
)

// Text struct for simplest text manipulation
type Text struct {
	// Full text (poem) to learn
	SourceText string
	// Sliced Source Text by lines
	// And Lines sliced by words
	textLines [][]string
	// Current line from textLines
	currentLineId uint
	// Current word id from currentLine
	currentWordId uint
}

// Make new Text Object
// Return error if cannot slice SourceText
func New(source string) (*Text, error) {
	text := &Text{
		SourceText:    source,
		currentLineId: 0,
		currentWordId: 0,
		textLines:     make([][]string, 1),
	}
	err := text.sliceSourceText()
	if err != nil {
		return nil, errors.New("Error while slicing source text!")
	}
	return text, nil
}

// Setter func for SourceText
func (t *Text) SetText(text string) {
	t.SourceText = strings.TrimSpace(text)
}

// Slice SourceText Function by lines and
// lines by words.
// Return error if SourceText is empty
func (t *Text) sliceSourceText() error {
	if t.SourceText == "" {
		return errors.New("Source Text is empty!")
	}
	tempStringLines := strings.Split(t.SourceText, "\n")
	var tempLines [][]string
	for _, line := range tempStringLines {
		tempLines = append(tempLines, strings.Split(line, " "))
	}

	t.textLines = tempLines
	return nil
}

// Interator function for textLines[currentLineId]
func (t *Text) NextLine() ([]string, bool) {
	if t.currentLineId == uint(len(t.textLines)) {
		return nil, false
	}
	tempLine := t.textLines[t.currentLineId]
	t.currentLineId++
	return tempLine, true
}

// Iterator function for textLines[currentLineId][currentWordId]
func (t *Text) NextWord() (string, bool) {
	if t.currentWordId == uint(len(t.textLines[t.currentLineId])) {
		return "", false
	}
	tempWord := t.textLines[t.currentLineId][t.currentWordId]
	t.currentWordId++
	return tempWord, true
}
