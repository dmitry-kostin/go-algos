package trie

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		t Type
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test array tree",
			args: args{t: ArrayTrie},
		},
		{
			name: "Test map tree",
			args: args{t: MapTrie},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := New(tt.args.t)
			word := "apple"
			searchWords := []string{"a", "ap", "app", "appl"}
			tr.Insert(word)
			for _, w := range searchWords {
				if !tr.StartsWith(w) {
					t.Errorf("startsWith fail: %v not found in %v", w, word)
				}
			}
			if !tr.Search(word) {
				t.Errorf("Search fail: %v not found in tree", word)
			}
			word2 := "app"
			tr.Insert(word2)
			if !tr.Search(word2) {
				t.Errorf("Search fail: %v not found in tree", word2)
			}
			if !tr.StartsWith(word2) {
				t.Errorf("Search fail: %v not found in tree", word2)
			}
			if !tr.Search(word) {
				t.Errorf("Search fail: %v not found in tree", word)
			}
			word3 := "ape"
			if tr.Search(word3) || tr.StartsWith(word3) {
				t.Errorf("found not existing word: %v in tree", word3)
			}
		})
	}
}
