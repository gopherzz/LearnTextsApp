package worderaser_test

import (
	"testing"

	test "github.com/KYCb2/LearnTextsApp/pkg/testutils"
	"github.com/KYCb2/LearnTextsApp/pkg/textwork/worderaser"
)

func TestEraseWordsByLevel(t *testing.T) {
	test.Assert([]string{"Hello", "█████"}, worderaser.EraseWordsByLevel([]string{"Hello", "World"}, 1))
	test.Assert([]string{"Hello World", "██████"}, worderaser.EraseWordsByLevel([]string{"Hello World", "Worlds"}, 1))
	test.Assert([]string{"Hello", "█████", "!"}, worderaser.EraseWordsByLevel([]string{"Hello", "World", "!"}, 1))
}

func TestEraseWord(t *testing.T) {
	test.Assert("█████", worderaser.EraseWord("hello"))
	test.Assert("███████████", worderaser.EraseWord("hello world"))
	test.Assert("█", worderaser.EraseWord("!"))
}
