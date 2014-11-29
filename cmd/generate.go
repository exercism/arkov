package cmd

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"
)

// Generate uses a stored datastructure to output a markov chain.
func Generate(ctx *cli.Context) {
	markov, err := chain.FromFile(ctx.String("infile"))
	if err != nil {
		log.Fatal(err)
	}
	text := markov.Generate()
	fmt.Println(text)
}
