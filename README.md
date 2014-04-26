# exercism/arkov

A markov chain generator for exercism nitpicks.

**This project is based on the markov chain generator written for the [golang.org codewalk](http://golang.org/doc/codewalk/markov/).**

## Getting Started

Get the dependencies:

```bash
$ go get
```

Build the binary:

```bash
$ go build
```

## Usage

The program takes a newline-delimited text file as input, and generates a json file containing the markov chain.

```bash
$ ./arkov build --infile=data/seuss.txt --outfile=data/seuss.json
```

To generate output, use the json file:

```bash
$ ./arkov generate --infile=data/seuss.json
```

## License

BSD

Copyright (c) 2012 The Go Authors. All rights reserved.
