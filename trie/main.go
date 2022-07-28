package trie

type Type int8

const (
	ArrayTrie Type = iota
	MapTrie
)

type Trie interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}

func New(t Type) Trie {
	if t == ArrayTrie {
		return newArray()
	}
	if t == MapTrie {
		return newMap()
	}
	return nil
}
