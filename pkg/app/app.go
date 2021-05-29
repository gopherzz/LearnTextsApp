package app

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/KYCb2/LearnTextsApp/pkg/textwork"
	"github.com/KYCb2/LearnTextsApp/pkg/textwork/utils"
	"github.com/KYCb2/LearnTextsApp/pkg/textwork/worderaser"
)

func Run() {
	text := *textwork.New("Hello world\nWorld hello\nworld !")
	err := text.SliceSourceText()
	if err != nil {
		panic(err)
	}
	for str, ok := text.NextLine(); ok; str, ok = text.NextLine() {
		rand.Seed(time.Now().Unix())
		erasedLine := worderaser.EraseRandomWordsByLevel(str, 1)
		fmt.Println(utils.StringLine(erasedLine))
	}
}
