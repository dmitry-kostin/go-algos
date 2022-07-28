## Trie data structure

In computer science, a trie, also called digital tree or prefix tree, is a type of k-ary search tree, a tree data structure used for locating specific keys from within a set. These keys are most often strings, with links between nodes defined not by the entire key, but by individual characters.

### Usage

```go
func main() {
	tr := trie.New(trie.ArrayTrie)
	tr.Insert("hello")
	fmt.Println(tr.Search("hello"))
	fmt.Println(tr.Search("hell"))
	fmt.Println(tr.StartsWith("he"))
}
```

