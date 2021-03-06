package worderaser_test

import (
	"testing"

	test "github.com/KYCb2/LearnTextsApp/pkg/testutils"
	"github.com/KYCb2/LearnTextsApp/pkg/textwork/worderaser"
)

func TestEraseWordsByLevel(t *testing.T) {
	test.AssertEqual([]string{"█████", "World"}, worderaser.EraseWordsByLevel([]string{"Hello", "World"}, 1))
	test.AssertEqual([]string{"███████████", "Worlds"}, worderaser.EraseWordsByLevel([]string{"Hello World", "Worlds"}, 1))
	test.AssertEqual([]string{"█████", "█████", "!"}, worderaser.EraseWordsByLevel([]string{"Hello", "World", "!"}, 2))
}

func TestEraseWord(t *testing.T) {
	test.AssertEqual("█████", worderaser.EraseWord("hello"))
	test.AssertEqual("███████████", worderaser.EraseWord("hello world"))
	test.AssertEqual("█", worderaser.EraseWord("!"))
}
