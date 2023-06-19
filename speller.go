package speller

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sajari/fuzzy"
)

func Suggestions(word string, dictionary []string) []string {
	fmt.Println("Searching...")
	model := fuzzy.NewModel()
	model.SetThreshold(3)

	model.SetDepth(1)
	model.Train(dictionary)

	correctWords := model.Suggestions(word, true)

	if len(correctWords) > 0 {
		fmt.Println("Found words", correctWords)
	} else {
		fmt.Println("No words found for", word)
	}

	return correctWords
}

func SimpleSuggestions(word string, dictionary []string) []string {
	filteredWords := []string{}
	for _, w := range dictionary {
		if strings.HasPrefix(w, word) {
			filteredWords = append(filteredWords, w)
		}
	}
	sort.Strings(filteredWords)
	return filteredWords
}
