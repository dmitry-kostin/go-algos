package trie

type Array struct {
	chars []*Array
	end   bool
}

func newArray() *Array {
	return &Array{
		chars: make([]*Array, 26),
	}
}

func (r *Array) Insert(word string) {
	p := r
	for i := 0; i < len(word); i++ {
		ind := word[i] - 'a'
		if p.chars[ind] == nil {
			p.chars[ind] = &Array{
				chars: make([]*Array, 26),
				end:   false,
			}
		}
		p = p.chars[ind]
	}
	p.end = true
}

func (r *Array) Search(word string) bool {
	p := r
	for i := 0; i < len(word); i++ {
		ind := word[i] - 'a'
		if p.chars[ind] == nil {
			return false
		}
		p = p.chars[ind]
	}
	return p.end
}

func (r *Array) StartsWith(prefix string) bool {
	p := r
	for i := 0; i < len(prefix); i++ {
		ind := prefix[i] - 'a'
		if p.chars[ind] == nil {
			return false
		}
		p = p.chars[ind]
	}
	return true
}
