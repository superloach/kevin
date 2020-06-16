// +build example

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/superloach/kevin"
)

func main() {
	input := strings.Join(os.Args[1:], " ")
	words := kevin.Wordlist30k()

	for n, km := range map[string]kevin.KeyMap{
		"qwerty":  kevin.KeyMapQWERTY,
		"dvorak":  kevin.KeyMapDvorak,
		"colemak": kevin.KeyMapColemak,
	} {
		fmt.Println(n)

		for i := range input {
			inp := input[:i+1]

			fmt.Println(inp, kevin.Suggest(inp, 5, words, km))
		}
	}
}
