package kevin

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Wordlist []string

func Wordlist30k() Wordlist {
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

	ws := make(Wordlist, 0)
	for _, l := range strings.Split(string(bs), "\n") {
		l = strings.TrimSpace(l)
		if len(l) > 0 {
			ws = append(ws, l)
		}
	}

	return ws
}
