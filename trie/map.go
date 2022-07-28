package trie

type Map struct {
	nodes map[byte]*Map
	end   bool
}

func newMap() *Map {
	return &Map{
		nodes: make(map[byte]*Map),
	}
}

func (r *Map) Insert(word string) {
	p := r
	for i := 0; i < len(word); i++ {
		if p.nodes[word[i]] == nil {
			p.nodes[word[i]] = &Map{
				nodes: make(map[byte]*Map),
			}
		}
		p = p.nodes[word[i]]
	}
	p.end = true
}

func (r *Map) Search(word string) bool {
	p := r
	for i := 0; i < len(word); i++ {
		if p.nodes[word[i]] == nil {
			return false
		}
		p = p.nodes[word[i]]
	}
	return p.end
}

func (r *Map) StartsWith(prefix string) bool {
	p := r
	for i := 0; i < len(prefix); i++ {
		if p.nodes[prefix[i]] == nil {
			return false
		}
		p = p.nodes[prefix[i]]
	}
	return true
}
