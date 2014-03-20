package main

import (
	"bufio"
	"strings"
	"fmt"
	"flag"
	"github.com/exercism/arkov/chain"
	"os"
	"time"
	"math/rand"
)

func main() {
	numWords := flag.Int("words", 100, "maximum number of words to print")
	prefixLen := flag.Int("prefix", 2, "prefix length in words")
	filename := flag.String("file", "", "name of file to read from")

	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	c := chain.NewChain(*prefixLen)

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Build(strings.NewReader(scanner.Text()))
	}
	text := c.Generate(*numWords)
	fmt.Println(text)
}
