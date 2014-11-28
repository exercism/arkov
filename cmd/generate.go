package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"
)

// Generate uses a stored datastructure to output a markov chain.
func Generate(ctx *cli.Context) {
	markov := chain.FromFile(ctx.String("infile"))
	text := markov.Generate()
	fmt.Println(text)
}
