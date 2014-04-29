package chain

import (
	"bytes"
	"reflect"
	"testing"
)

var chainTests = []Node{
	{" ", []string{"I"}},
	{" I", []string{"am"}},
	{"I am", []string{"a", "not"}},
	{"a free", []string{"man."}},
	{"free man.", []string{"I"}},
	{"man. I", []string{"am"}},
	{"not a", []string{"number."}},
	{"am a", []string{"free"}},
	{"am not", []string{"a"}},
}

func TestChain(t *testing.T) {
	c := NewChain(2)
	s := "I am a free man. I am not a number."
	br := bytes.NewBufferString(s)
	c.Build(br)

	if len(chainTests) != len(c.Nodes) {
		t.Fatalf("The chain is the wrong size. Expected: %v, Got: %v", len(chainTests), len(c.Nodes))
	}

	for _, tt := range chainTests {
		node := c.FindNode(tt.Key)

		if !reflect.DeepEqual(tt.Fragments, node.Fragments) {
			t.Fatalf("%v != %v", tt.Fragments, node.Fragments)
		}
	}
}
