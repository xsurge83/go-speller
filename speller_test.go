package speller_test

import (
	"fmt"
	"testing"

	"github.com/sajari/fuzzy"
)

func TestSpelling(t *testing.T) {
	// dictionary := []string{
	// 	"apple", "banana", "cherry", "date", "elderberry",
	// 	"fig", "grape", "honeydew", "kiwi", "lemon",
	// }
	dictionary := []string{"bob", "your", "uncle", "dynamite", "delicate", "biggest", "big", "bigger", "aunty", "you're"}

	// correctWords := speller.CheckSpelling("bi", dictionary)
	model := fuzzy.NewModel()
	model.SetThreshold(2)

	// model.SetDepth()
	model.Train(dictionary)

	correctWords := model.Suggestions("bo", true)
	fmt.Println("	Swap test (uncel) : ", model.SpellCheck("uncel"))

	if len(correctWords) == 1 {
		t.Errorf("Expected %d", len(correctWords))
	}

}
