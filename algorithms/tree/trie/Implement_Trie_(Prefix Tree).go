/*
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.
*/

type TrieNode struct{
    children [26]*TrieNode
    isLeaf bool
}

type Trie struct {
    root *TrieNode
}


func Constructor() Trie {
    return Trie{
        root: &TrieNode{
            children: [26]*TrieNode{},
            isLeaf: false,
        },
    }
}


func (this *Trie) Insert(word string)  {
    curr := this.root
    for i := 0; i < len(word); i++{
        idx := (word[i]-'a')
        if curr.children[idx] == nil{
            curr.children[idx] = &TrieNode{
                children: [26]*TrieNode{},
                isLeaf: false,
            }
        }
        curr = curr.children[idx]
    }
    curr.isLeaf = true
}


func (this *Trie) Search(word string) bool {
    curr := this.root
    for i := 0; i < len(word); i++{
        idx := (word[i]-'a')
        if curr.children[idx] == nil{
            return false
        }
        curr = curr.children[idx]
    }
    if !curr.isLeaf {return false} 
    return true
}


func (this *Trie) StartsWith(word string) bool {
    curr := this.root
    for i := 0; i < len(word); i++{
        idx := (word[i]-'a')
        if curr.children[idx] == nil{
            return false
        }
        curr = curr.children[idx]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */