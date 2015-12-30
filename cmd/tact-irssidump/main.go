package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// A tool for dumping a text corpus from irssi logs. Assumes the log
// format I use.

var ignoreflag = flag.String("ignorenicks", "", "space-separated nicks to ignore")

func main() {
	flag.Parse()

	ignore := strings.Fields(strings.ToLower(*ignoreflag))

	for _, fn := range flag.Args() {
		f, err := os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}

		dump(f, ignore)
		f.Close()
	}
}

var logline = regexp.MustCompile(
	`^[0-9][0-9]:[0-9][0-9] <(.*?)> ([A-Za-z0-9_]+[:,] )?(.*)`)

func dump(r io.Reader, ignores []string) error {
	s := bufio.NewScanner(r)
	for s.Scan() {
		m := logline.FindStringSubmatch(s.Text())
		if m == nil {
			continue
		}

		from, _, text := m[1], m[2], m[3]
		if ignore(from, ignores) {
			continue
		}

		text = strings.TrimSpace(text)
		if text != "" {
			fmt.Println(text)
		}
	}

	return s.Err()
}

func ignore(nick string, ignores []string) bool {
	nick = strings.ToLower(nick)
	for _, pat := range ignores {
		if strings.HasPrefix(nick, pat) {
			return true
		}
	}
	return false
}
