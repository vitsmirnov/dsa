package main

import "fmt"

const alphabetLen = 'z' - 'a' + 1

type TrieNode struct {
	children [alphabetLen]*TrieNode
	wordEnd  bool
}

type Trie struct {
	root *TrieNode
}

func MakeTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}

func (this *Trie) Insert(word string) {
	// time: O(|word|), space: O(|word|)

	node := this.root
	for i := range word {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &TrieNode{}
		}
		node = node.children[c]
	}
	node.wordEnd = true
}

func (this *Trie) SearchNode(word string) *TrieNode {
	// time: O(|word|), space: O(1)

	node := this.root
	for i, l := 0, len(word); i < l && node != nil; i++ {
		node = node.children[word[i]-'a']
	}
	return node
}

func (this *Trie) Search(word string) bool {
	// time: O(|word|), space: O(1)

	node := this.SearchNode(word)
	return node != nil && node.wordEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	// time: O(|word|), space: O(1)

	return this.SearchNode(prefix) != nil
}

func (this *Trie) GetWords() []string {
	// time: O(n), space: O(h)

	words := []string{}

	var dfs func(node *TrieNode, prefix []byte)
	dfs = func(node *TrieNode, prefix []byte) {
		if node == nil {
			return
		}

		if node.wordEnd {
			words = append(words, string(prefix))
		}
		for char, child := range node.children {
			dfs(child, append(prefix, byte(char)+'a'))
		}
	}

	dfs(this.root, []byte{})
	return words
}

func (this *Trie) Print() {
	for _, word := range this.GetWords() {
		fmt.Println(word)
	}
}

func demo() {
	words := []string{"xy", "app", "apple", "beer", "add", "jam", "rental"}
	trie := MakeTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	fmt.Println("Words:")
	trie.Print()
	fmt.Println()
	word := "app"
	fmt.Printf("Word %q is in the trie: %v\n", word, trie.Search(word))
	word = "ren"
	fmt.Printf("Word with prefix %q is in the trie: %v\n", word, trie.StartsWith(word))
	word = "xy"
	fmt.Printf("Word with prefix %q is in the trie: %v\n", word, trie.StartsWith(word))
	word = "xyz"
	fmt.Printf("Word with prefix %q is in the trie: %v\n", word, trie.StartsWith(word))
}

func main() {
	demo()
}
