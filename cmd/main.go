package main

import (
	"fmt"
	"speller"

	"github.com/manifoldco/promptui"
)

func main() {
	const DICT_PATH = "/usr/share/dict/words"
	dictionary, err := speller.LoadDictionary(DICT_PATH)
	if err != nil {
		fmt.Printf("Failed to load dictionary: %s\n", err)
		return
	}

	prompt := promptui.Prompt{
		Label: "Type a word",
	}

	word, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	suggestions := speller.Suggestions(word, dictionary)

	selectPrompt := promptui.Select{
		Label: "Select an option",
		Items: suggestions,
	}

	_, result, err := selectPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Selected option: %s\n", result)
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("Enter a word (or q to quit): ")
	// 	word, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Printf("Failed to read input: %s\n", err)
	// 		break
	// 	}
	// 	word = strings.TrimSpace(word)
	// 	if word == "q" {
	// 		break
	// 	}
	// 	speller.Suggestions(word, dictionary)
	// }
}
