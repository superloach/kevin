// +build example

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/superloach/kevin"
)

var (
	layout = flag.String("layout", "qwerty", "keyboard layout to use (supported: qwerty, dvorak, colemak)")
	num    = flag.Int("num", 5, "number of suggestions to give")
)

var layouts = map[string]kevin.Layout{
	"qwerty":  kevin.QWERTY,
	"dvorak":  kevin.Dvorak,
	"colemak": kevin.Colemak,
}

func main() {
	flag.Parse()

	fmt.Println("downloading wordlist")
	wordlist := wordlist30k()

	fmt.Printf("using layout %s\n", *layout)
	l := layouts[*layout]

	fmt.Printf("giving %d suggestions\n", *num)

	words := flag.Args()
	fmt.Printf("words: %s\n", words)

	for _, word := range words {
		fmt.Println(word, l.Suggest(word, *num, wordlist))
	}
}

func wordlist30k() []string {
	resp, err := http.Get("https://raw.githubusercontent.com/derekchuank/high-frequency-vocabulary/master/30k.txt")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		panic(err)
	}

	ws := make([]string, 0)
	for _, l := range strings.Split(string(bs), "\n") {
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			ws = append(ws, l)
		}
	}

	return ws
}
