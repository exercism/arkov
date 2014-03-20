package chain

import "strings"

type prefix []string

func (p prefix) key() string {
	return strings.Join(p, " ")
}

func (p prefix) shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}
