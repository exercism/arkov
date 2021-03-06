package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/exercism/arkov/chain"
)

// Build creates and stores a markov datastructure.
func Build(ctx *cli.Context) error {
	markov := chain.NewChain(ctx.Int("prefix"))

	file, err := os.Open(ctx.String("infile"))
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		markov.Build(strings.NewReader(scanner.Text()))
	}

	if err := markov.ToFile(ctx.String("outfile")); err != nil {
		return err
	}
	return nil
}
