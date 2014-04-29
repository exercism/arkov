package chain

import (
	"bytes"
	"reflect"
	"testing"
)

var chainTests = []struct {
	input string
	nodes []Node
}{
	{
		"I am a free man. I am not a number.",
		[]Node{
			{" ", []string{"I"}},
			{" I", []string{"am"}},
			{"I am", []string{"a", "not"}},
			{"a free", []string{"man."}},
			{"free man.", []string{"I"}},
			{"man. I", []string{"am"}},
			{"not a", []string{"number."}},
			{"am a", []string{"free"}},
			{"am not", []string{"a"}},
		},
	},
}

func TestChain(t *testing.T) {
	for _, tt := range chainTests {
		c := NewChain(2)
		br := bytes.NewBufferString(tt.input)
		c.Build(br)

		if len(tt.nodes) != len(c.Nodes) {
			t.Fatalf("The chain is the wrong size. Expected: %v, Got: %v", len(tt.nodes), len(c.Nodes))
		}

		for _, node := range tt.nodes {
			fragments := c.FindNode(node.Key).Fragments

			if !reflect.DeepEqual(node.Fragments, fragments) {
				t.Fatalf("%v != %v", node.Fragments, fragments)
			}
		}
	}
}
