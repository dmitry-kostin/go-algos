## Trie data structure

In computer science, a trie, also called digital tree or prefix tree, is a type of k-ary search tree, a tree data structure used for locating specific keys from within a set. These keys are most often strings, with links between nodes defined not by the entire key, but by individual characters.

There are two implementaions under the hood:
* array based tree
* map based tree

![a9bb8c2e7485693d5e4bce657e673bda](https://user-images.githubusercontent.com/1920678/181521199-38fed2c3-3433-48ee-a44b-61f4f7014bb0.png)

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

