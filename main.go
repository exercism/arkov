package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"
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
				cli.StringFlag{
					Name:  "infile, f",
					Usage: "File containing original text",
				},
				cli.StringFlag{
					Name:  "outfile, o",
					Usage: "File to store markov chain",
				},
				cli.IntFlag{
					Name:  "prefix, p",
					Value: 2,
					Usage: "Prefix length",
				},
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
				cli.StringFlag{
					Name:  "infile, f",
					Usage: "File containing chain data",
				},
			},
			Action: func(c *cli.Context) {
				markov := chain.FromFile(c.String("infile"))
				text := markov.Generate()
				fmt.Println(text)
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
