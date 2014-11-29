package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/exercism/arkov/chain"

	// do not expose exported names of postgresql driver,
	// but make it available to database/sql
	_ "github.com/lib/pq"
)

const connInfo = `user=exercism dbname=exercism_seeds sslmode=disable`

var errUsage = errors.New(`USAGE: arkov seed --dir=/path/to/dir/containing/json/files`)

// Seed stores markov chains as comments in the exercism seed database.
func Seed(ctx *cli.Context) {
	if ctx.String("dir") == "" {
		log.Fatal(errUsage)
	}

	db, err := sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// get a list of the languages available
	var languages []string
	var s string
	rows, err := db.Query(`SELECT DISTINCT(language) FROM submissions`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&s); err != nil {
			log.Fatal(err)
		}
		languages = append(languages, s)
	}

	stmt, err := db.Prepare(`UPDATE comments SET body = $1 WHERE id = $2`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, language := range languages {
		path := filepath.Join(ctx.String("dir"), fmt.Sprintf("%s.json", language))
		markov, err := chain.FromFile(path)
		if err != nil {
			log.Println(fmt.Errorf("WARN: cannot do %s because %s", language, err))
		}

		// get the ids of the comments
		var ids []int
		rows, err = db.Query(`SELECT c.id FROM comments c INNER JOIN submissions s ON c.submission_id=s.id AND s.language = $1`, language)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var i int
		for rows.Next() {
			if err := rows.Scan(&i); err != nil {
				log.Fatal(err)
			}
			ids = append(ids, i)
		}
		log.Println("generating nitpicks for", len(ids), "comments in", language)

		for _, id := range ids {
			text := markov.Generate()
			_, err := stmt.Exec(text, id)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
