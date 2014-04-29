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

// Chain contains sentence fragments that can be recombined.
type Chain struct {
	Nodes     []*Node
	PrefixLen int `json:"prefix_len"`
}

// Build creates a new chain from newline delimited text.
func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	p := prefix(make([]string, c.PrefixLen))
	var s string
	for {
		if _, err := fmt.Fscan(br, &s); err != nil {
			break
		}
		key := p.key()
		c.appendFragment(key, s)
		p.shift(s)
	}
}

// Generate creates multiple paragraphs.
func (c *Chain) Generate() string {
	n := paragraphCount()
	paragraphs := make([]string, n, n)
	for i := 0; i < n; i++ {
		paragraphs = append(paragraphs, c.GenerateParagraph())
	}
	return strings.TrimLeft(strings.Join(paragraphs, "\n\n"), "\n ")
}

// GenerateParagraph creates a single paragraph.
func (c *Chain) GenerateParagraph() string {
	p := make(prefix, c.PrefixLen)
	n := wordCount()

	var words []string
	for {
		node := c.findNode(p.key())
		if node == nil {
			break
		}
		choices := node.Fragments
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.shift(next)
		if len(choices) > n && completesSentence(next) {
			break
		}
	}
	return strings.Join(words, " ")
}

// ToFile marshalls a chain to a file in JSON format.
func (c *Chain) ToFile(path string) {
	bytes, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Failed to marshal: %v\n", err)
	}

	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		fmt.Printf("Unable to write to %s: %v\n", path, err)
	}
}

func (c *Chain) findNode(key string) *Node {
	for _, n := range c.Nodes {
		if n.Key == key {
			return n
		}
	}
	return nil
}

func (c *Chain) appendFragment(key, fragment string) {
	node := c.findNode(key)
	if node == nil {
		node = new(Node)
		node.Key = key
		c.Nodes = append(c.Nodes, node)
	}
	node.Fragments = append(node.Fragments, fragment)
}

// NewChain creates an empty chain.
// The prefix of each node will be prefixLen words long.
func NewChain(prefixLen int) *Chain {
	return &Chain{[]*Node{}, prefixLen}
}

// FromFile unmarshalls a stored JSON chain.
func FromFile(path string) *Chain {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Unable to read file: %s\n", path)
		return nil
	}

	c := new(Chain)

	err = json.Unmarshal(bytes, c)
	if err != nil {
		fmt.Printf("Cannot unmarshall: %v\n", err)
		return nil
	}
	return c
}

func wordCount() int {
	numbers := []int{13, 21, 34, 55, 89, 144}
	return numbers[rand.Intn(len(numbers))]
}

func paragraphCount() int {
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 5}
	return numbers[rand.Intn(len(numbers))]
}

func completesSentence(s string) bool {
	return strings.LastIndexAny(s, "?!.") == len(s)-1
}
