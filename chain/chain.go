package chain

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type Chain struct {
	Data      map[string][]string
	PrefixLen int `json:"submission_path"`
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(prefix, c.PrefixLen)
	for {
		var s string
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.key()
		c.Data[key] = append(c.Data[key], s)
		p.shift(s)
	}
}

func (c *Chain) Generate(n int) string {
	p := make(prefix, c.PrefixLen)
	var words []string
	for i := 0; i < n; i++ {
		choices := c.Data[p.key()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.shift(next)
	}
	return strings.Join(words, " ")
}
