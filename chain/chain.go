package chain

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
)

type Chain struct {
	Data      map[string][]string
	PrefixLen int `json:"prefix_len"`
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

func wordCount() int {
	numbers := []int{13, 21, 34, 55, 89, 144}
	return numbers[rand.Intn(len(numbers))]
}

func paragraphCount() int {
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	return numbers[rand.Intn(len(numbers))]
}

func (c *Chain) Generate() string {
	n := paragraphCount()
	paragraphs := make([]string, n, n)
	for i := 0; i < n; i++ {
		paragraphs = append(paragraphs, c.GenerateParagraph())
	}
	return strings.TrimLeft(strings.Join(paragraphs, "\n\n"), "\n ")
}

func (c *Chain) GenerateParagraph() string {
	p := make(prefix, c.PrefixLen)
	var words []string
	n := wordCount()
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

func (c Chain) ToFile(path string) {
	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Failed to marshal: %v\n", err)
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		fmt.Printf("Unable to write to %s: %v\n", path, err)
	}
}

func FromFile(path string) (c Chain) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Unable to read file: %s\n", path)
		return
	}

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		fmt.Printf("Cannot unmarshall: %v\n", err)
		return
	}
	return
}
