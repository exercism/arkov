package main

import (
	"bufio"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	app := cli.NewApp()
	app.Name = "exercismarkov"
	app.Usage = "Create fake exercism nitpicks"

	app.Commands = []cli.Command{
		{
			Name:      "build",
			ShortName: "b",
			Usage:     "build a markov chain",
			Flags: []cli.Flag{
				cli.StringFlag{"infile, f", "", "File containing original text"},
				cli.StringFlag{"outfile, o", "", "File to store markov chain"},
				cli.IntFlag{"prefix, p", 2, "Prefix length"},
			},
			Action: func(c *cli.Context) {
				markov := chain.NewChain(c.Int("prefix"))

				file, err := os.Open(c.String("infile"))
				if err != nil {
					println(err)
					return
				}
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					markov.Build(strings.NewReader(scanner.Text()))
				}

				markov.ToFile(c.String("outfile"))
			},
		},
		{
			Name:      "generate",
			ShortName: "g",
			Usage:     "generate a new comment",
			Flags: []cli.Flag{
				cli.StringFlag{"infile, f", "", "File containing chain data"},
			},
			Action: func(c *cli.Context) {
				markov := chain.FromFile(c.String("infile"))
				text := markov.Generate()
				fmt.Println(text)
			},
		},
	}

	app.Run(os.Args)
}
