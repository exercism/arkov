package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/cmd"
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
			Action: cmd.Build,
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
			Action: cmd.Generate,
		},
		{
			Name:      "seed",
			ShortName: "s",
			Usage:     "store seeds into exercism seed database",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dir, d",
					Usage: "Directory containing files of chain data",
				},
			},
			Action: cmd.Seed,
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
