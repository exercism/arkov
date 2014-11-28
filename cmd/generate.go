package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"
)

func Generate(ctx *cli.Context) {
	markov := chain.FromFile(ctx.String("infile"))
	text := markov.Generate()
	fmt.Println(text)
}
