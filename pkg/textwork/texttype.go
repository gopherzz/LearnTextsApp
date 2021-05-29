package textwork

import (
	"errors"
	"strings"
)

type Text struct {
	SourceText    string
	currentLineId uint
	currentWordId uint
	currentLine   []string
	textLines     [][]string
}

func New(source string) *Text {
	return &Text{
		SourceText:    source,
		currentLineId: 0,
		currentWordId: 0,
		textLines:     make([][]string, 1),
	}
}

func (t *Text) SetText(text string) {
	t.SourceText = strings.TrimSpace(text)
}

func (t *Text) SliceSourceText() error {
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

func (t *Text) NextLine() ([]string, bool) {
	if t.currentLineId == uint(len(t.textLines)) {
		return nil, false
	}
	tempLine := t.textLines[t.currentLineId]
	t.currentLineId++
	return tempLine, true
}

func (t *Text) NextWord() (string, bool) {
	if t.currentWordId == uint(len(t.textLines[t.currentLineId])) {
		return "", false
	}
	tempWord := t.textLines[t.currentLineId][t.currentWordId]
	t.currentWordId++
	return tempWord, true
}
