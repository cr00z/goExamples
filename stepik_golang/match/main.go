// $ ./match -p is 'go is awesome'
// go is awesome
// $ ./match -p was 'go is awesome'
// (empty)

package main

import (
	"fmt"
	"errors"
	"flag"
	"os"
	"strings"
	// "github.com/cr00z/match/glob"
	"github.com/sahilm/fuzzy"
)

func main() {
    pattern, src, err := readInput()
    if err != nil {
        fail(err)
    }
	// isMatch, err := glob.Match(pattern, src)
    // if err != nil {
    //     fail(err)
    // }
	matches := fuzzy.Find(pattern, []string{src})
    isMatch := len(matches) > 0
    if isMatch {
        fmt.Println(src)
    }
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (pattern, src string, err error) {
	flag.StringVar(&pattern, "p", "", "pattern to match against")
	flag.Parse()
	if pattern == "" {
		return pattern, src, errors.New("missing pattern")
	}
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return pattern, src, errors.New("missing string to match")
	}
	return pattern, src, nil
}

// fail prints the error and exits.
func fail(err error) {
    fmt.Println(err)
	os.Exit(1)
}