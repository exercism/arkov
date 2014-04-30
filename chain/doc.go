/*
Package chain is used to generate fake comments.

The chain is built using a corpus of existing data, preferably in the style
of the comments that you wish to generate.

A chain consists of a prefix followed by all the possible words that appeared
after the prefix in the training data.  A prefix may consist of one or more
words, and is consistent for an entire chain.

There are helper methods which try to optimize the output for seeming
real-ish. A comment will, by default, have a small number of paragraphs, and
the paragraphs will be neither particularly short, nor extremely long. Where
possible, the generated output will end with punctuation.
*/
package chain
