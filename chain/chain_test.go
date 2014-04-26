package chain

import (
	"bytes"
	"reflect"
	"testing"
)

func TestChain(t *testing.T) {
	c := NewChain(2)
	s := "I am a free man. I am not a number."
	br := bytes.NewBufferString(s)
	expected := map[string][]string{
		" ":         []string{"I"},
		" I":        []string{"am"},
		"I am":      []string{"a", "not"},
		"a free":    []string{"man."},
		"free man.": []string{"I"},
		"man. I":    []string{"am"},
		"not a":     []string{"number."},
		"am a":      []string{"free"},
		"am not":    []string{"a"},
	}

	c.Build(br)
	if !reflect.DeepEqual(expected, c.Data) {
		t.Fatalf("%v != %v", expected, c.Data)
	}
}
