package chain

import (
	"bytes"
	"reflect"
	"sort"
	"testing"
)

type ByKey []*Node

func (n ByKey) Len() int           { return len(n) }
func (n ByKey) Less(i, j int) bool { return n[i].Key < n[j].Key }
func (n ByKey) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }

func TestChain(t *testing.T) {
	c := NewChain(2)
	s := "I am a free man. I am not a number."
	br := bytes.NewBufferString(s)
	expected := []*Node{
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

	c.Build(br)

	sort.Sort(ByKey(expected))
	sort.Sort(ByKey(c.Nodes))

	if !reflect.DeepEqual(expected, c.Nodes) {
		t.Fatalf("%v != %v", expected, c.Nodes)
	}
}
