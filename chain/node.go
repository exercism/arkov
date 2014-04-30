package chain

// Node connects a key with all the possible words that may follow it.
// A key may be one or more words. This number will be consistent for
// an entire chain (collection of nodes).
type Node struct {
	Key       string
	Fragments []string
}
