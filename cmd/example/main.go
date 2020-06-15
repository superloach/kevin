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

	for i := range input {
		inp := input[:i+1]

		fmt.Println(inp, kevin.Suggest(
			inp, words, 3,
			kevin.KeyMapQWERTY,
		))
	}
}
