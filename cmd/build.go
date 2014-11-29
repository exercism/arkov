package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"
)

// Build creates and stores a markov datastructure.
func Build(ctx *cli.Context) {
	markov := chain.NewChain(ctx.Int("prefix"))

	file, err := os.Open(ctx.String("infile"))
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		markov.Build(strings.NewReader(scanner.Text()))
	}

	if err := markov.ToFile(ctx.String("outfile")); err != nil {
		log.Fatal(err)
	}
}
