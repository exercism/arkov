package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"
)

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

	markov.ToFile(ctx.String("outfile"))
}
