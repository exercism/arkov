# exercism/arkov

A markov chain generator for exercism nitpicks.

This project is based on the markov chain generator written for the [golang.org codewalk](http://golang.org/doc/codewalk/markov/).

## Usage

The program takes a text file. Each line is built into the markov chain independently.

```bash
go run main.go -file=comments.dat
```

## License

BSD

Copyright (c) 2012 The Go Authors. All rights reserved.
