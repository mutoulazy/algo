package main

import "fmt"

// tire node
type trieNode struct {
	char     string             // Unicode 字符
	isEnding bool               // 是否是单词结尾
	children map[rune]*trieNode // 该节点的子节点
}

// new Trie node
func NewTrieNode(char string) *trieNode {
	return &trieNode{
		char:     char,
		isEnding: false,
		children: make(map[rune]*trieNode),
	}
}

// trie tree
type Trie struct {
	root *trieNode
}

func NewTrieTree() *Trie {
	rootNode := NewTrieNode("/")
	return &Trie{rootNode}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, code := range word {
		childNode, ok := node.children[code]
		if !ok {
			// new child node
			childNode = NewTrieNode(string(code))
			node.children[code] = childNode
		}
		node = childNode
	}
	node.isEnding = true
}

func (t *Trie) Find(word string) bool {
	node := t.root
	for _, code := range word {
		childNode, ok := node.children[code]
		if !ok {
			// not find
			return false
		}
		node = childNode
	}
	if false == node.isEnding {
		// not full compare
		return false
	}
	return true
}

func main() {
	trie := NewTrieTree()
	words := []string{"Golang", "Language", "Trie", "Go", "木头人"}
	for _, word := range words {
		trie.Insert(word)
	}
	term := "木头"
	if trie.Find(term) {
		fmt.Printf("contain \"%s\"\n", term)
	} else {
		fmt.Printf("not contain \"%s\"\n", term)
	}
}
