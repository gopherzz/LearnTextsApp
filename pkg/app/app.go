package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/KYCb2/LearnTextsApp/pkg/textwork"
	"github.com/KYCb2/LearnTextsApp/pkg/textwork/textutils"
	"github.com/KYCb2/LearnTextsApp/pkg/textwork/worderaser"
)

func Run() {
	filepath, _level := os.Args[1], os.Args[2]
	level, err := strconv.ParseInt(_level, 10, 32)
	if err != nil {
		log.Fatal("Level is not number!")
	}
	// Read source file
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	// Create new text object
	// for easiest work with text
	text, err := textwork.New(string(bytes))
	if err != nil {
		panic(err)
	}
	// Print Erased Text
	for str, ok := text.NextLine(); ok; str, ok = text.NextLine() {
		erasedLine := worderaser.EraseWordsByLevel(str, uint(level))
		fmt.Println(textutils.LineToString(erasedLine))
	}
}
