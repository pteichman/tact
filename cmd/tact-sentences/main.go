package main

// tact-sentences takes each line of stdin, splits its sentences, and
// outputs those sentences.

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/neurosnap/sentences"
	"github.com/neurosnap/sentences/english"
)

func nonempty(sents []*sentences.Sentence) []string {
	var ret []string
	for _, sent := range sents {
		text := strings.TrimSpace(sent.Text)
		if text != "" {
			ret = append(ret, text)
		}
	}
	return ret
}

func main() {
	tok, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		input := strings.TrimSpace(s.Text())

		sents := tok.Tokenize(input)
		if len(sents) == 0 {
			continue
		}

		texts := nonempty(sents)
		if len(texts) == 1 {
			continue
		}

		for _, text := range texts {
			fmt.Println(text)
		}
	}
}
