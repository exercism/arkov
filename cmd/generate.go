package cmd

import (
	"fmt"

	"github.com/urfave/cli"
	"github.com/exercism/arkov/chain"
)

// Generate uses a stored datastructure to output a markov chain.
func Generate(ctx *cli.Context) error {
	markov, err := chain.FromFile(ctx.String("infile"))
	if err != nil {
		return err
	}
	text := markov.Generate()
	fmt.Println(text)
	return nil
}
