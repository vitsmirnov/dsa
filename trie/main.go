package main

import "fmt"

const alphabetLen = 'z' - 'a' + 1

type TrieNode struct {
	children  []*TrieNode
	isWordEnd bool
}

func MakeTrieNode() *TrieNode {
	return &TrieNode{
		children:  make([]*TrieNode, alphabetLen),
		isWordEnd: false}
}

type Trie struct {
	root *TrieNode
}

func MakeTrie() *Trie {
	return &Trie{root: MakeTrieNode()}
}

func (this *Trie) Add(word string) {
	// time: O(|word|), space: O(|word|)

	node := this.root
	for i := range word {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = MakeTrieNode()
		}
		node = node.children[c]
	}
	node.isWordEnd = true
}

func (this *Trie) FindNode(word string) *TrieNode {
	// time: O(|word|), space: O(1)

	node := this.root
	for i, l := 0, len(word); i < l && node != nil; i++ {
		node = node.children[word[i]-'a']
	}
	return node
}

func (this *Trie) Exists(word string) bool {
	// time: O(|word|), space: O(1)

	node := this.FindNode(word)
	return node != nil && node.isWordEnd
}

func (this *Trie) Exists2(word string, maxHammingDist int) bool {
	wordLen := len(word)

	var dfs func(node *TrieNode, pos int, hamDist int) bool
	dfs = func(node *TrieNode, pos int, hamDist int) bool {
		if node == nil || hamDist < 0 {
			return false
		}
		if pos == wordLen {
			return node.isWordEnd
		}

		c := int(word[pos] - 'a')
		if dfs(node.children[c], pos+1, hamDist) {
			return true
		}
		for i, child := range node.children {
			if i != c && dfs(child, pos+1, hamDist-1) {
				return true
			}
		}
		return false
	}

	return dfs(this.root, 0, maxHammingDist)
}

func (this *Trie) StartsWith(prefix string) bool {
	// time: O(|word|), space: O(1)

	return this.FindNode(prefix) != nil
}

func (this *Trie) GetWords() []string {
	// time: O(n), space: O(h)

	words := []string{}

	var dfs func(node *TrieNode, prefix []byte)
	dfs = func(node *TrieNode, prefix []byte) {
		if node == nil {
			return
		}

		if node.isWordEnd {
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
		trie.Add(word)
	}
	fmt.Println("Words:")
	trie.Print()
	fmt.Println()
	word := "app"
	fmt.Printf("Word %q is in the trie: %v\n", word, trie.Exists(word))
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
