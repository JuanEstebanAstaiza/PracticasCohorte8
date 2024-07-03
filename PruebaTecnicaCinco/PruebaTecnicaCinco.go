package main

/*import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func newTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) SearchPrefix(prefix string) []string {
	node := t.root
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			return nil
		}
		node = node.children[char]
	}
	return t.collectWords(node, prefix)
}

func (t *Trie) collectWords(node *TrieNode, prefix string) []string {
	var words []string
	if node.isEnd {
		words = append(words, prefix)
	}
	for char, childNode := range node.children {
		words = append(words, t.collectWords(childNode, prefix+string(char))...)
	}
	return words
}

func main() {
	trie := newTrie()
	words := []string{"apple", "app", "apricot", "banana", "bat", "batch", "batman"}
	for _, word := range words {
		trie.Insert(word)
	}

	prefix := "app"
	fmt.Printf("Words with prefix '%s': %v\n", prefix, trie.SearchPrefix(prefix))

	prefix = "bat"
	fmt.Printf("Words with prefix '%s' : %v\n", prefix, trie.SearchPrefix(prefix))
}*/

/*type SegmentTree struct {
	tree []int
	arr  []int
	n    int
}

func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)
	tree := make([]int, 4*n)
	segTree := &SegmentTree{tree: tree, arr: arr, n: n}
	segTree.build(0, 0, n-1)
	return segTree
}

func (st *SegmentTree) build(node, start, end int) {
	if start == end {
		st.tree[node] = st.arr[start]
	} else {
		mid := (start + end) / 2
		leftChild := 2*node + 1
		rightChild := 2*node + 2
		st.build(leftChild, start, mid)
		st.build(rightChild, mid+1, end)
		st.tree[node] = st.tree[leftChild] + st.tree[rightChild]
	}
}

func (st *SegmentTree) Update(index, value int) {
	st.updateHelper(0, 0, st.n-1, index, value)
}

func (st *SegmentTree) updateHelper(node, start, end, index, value int) {
	if start == end {
		st.arr[index] = value
		st.tree[node] = value

	} else {
		mid := (start + end) / 2
		leftChild := 2*node + 1
		rightChild := 2*node + 2
		if start <= index && index <= mid {
			st.updateHelper(leftChild, start, mid, index, value)
		} else {
			st.updateHelper(rightChild, mid+1, end, index, value)
		}
		st.tree[node] = st.tree[leftChild] + st.tree[rightChild]
	}
}

func (st *SegmentTree) QueryRange(L, R int) int {
	return st.queryRangeHelper(0, 0, st.n-1, L, R)
}

func (st *SegmentTree) queryRangeHelper(node, start, end, L, R int) int {
	if R < start || end < L {
		return 0
	}
	if L <= start && end <= R {
		return st.tree[node]
	}
	mid := (start + end) / 2
	leftChild := 2*node + 1
	rightChild := 2*node + 2
	leftSum := st.queryRangeHelper(leftChild, start, mid, L, R)
	rightSum := st.queryRangeHelper(rightChild, mid+1, end, L, R)
	return leftSum + rightSum

}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}
	segTree := NewSegmentTree(arr)

	fmt.Println("Initial array;", arr)
	fmt.Println("Sum of range [1,3]:", segTree.QueryRange(1, 3))

	segTree.Update(1, 10)
	fmt.Println("Updated array:", segTree.arr)
	fmt.Println("Sum of range [1,3]:", segTree.QueryRange(1, 3))
}*/

import (
	"fmt"
	"strings"
)

type JuegoAhorcado struct {
	palabra       string
	letras        []string
	intentos      int
	intentosMax   int
	palabraOculta string
}

func NuevoJuegoAhorcado(palabra string, intentosMax int) *JuegoAhorcado {
	palabra = strings.ToUpper(palabra)
	palabraOculta := strings.Repeat("_ ", len(palabra))
	return &JuegoAhorcado{

		palabra:       palabra,
		letras:        make([]string, 0),
		intentos:      intentosMax,
		intentosMax:   intentosMax,
		palabraOculta: palabraOculta,
	}
}

func (j *JuegoAhorcado) AdivinarLetra(letra string) {

	letra = strings.ToUpper(letra)
	if !strings.Contains(j.palabra, letra) {
		j.intentos--
	} else {
		j.letras = append(j.letras, letra)
		palabraOculta := ""
		for _, char := range j.palabra {
			if contains(j.letras, string(char)) {
				palabraOculta += string(char) + " "
			} else {
				palabraOculta += "_ "
			}
		}

		j.palabraOculta = palabraOculta
	}
}

func (j *JuegoAhorcado) HaGanado() bool {
	return j.palabraOculta == j.palabra
}

func (j *JuegoAhorcado) HaPerdido() bool {
	return j.intentos <= 0
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	juego := NuevoJuegoAhorcado("GOL", 5)

	for {
		fmt.Println("Intentos restantes: ", juego.intentos)
		fmt.Println("Palabra: ", juego.palabraOculta)

		var letra string
		fmt.Print("Adivina una letra: ")
		fmt.Scanln(&letra)

		juego.AdivinarLetra(letra)

		if juego.HaGanado() {
			fmt.Println("Has ganado! La palabra era: ", juego.palabra)
			break
		} else if juego.HaPerdido() {
			fmt.Println("Has perdido! La palabra era: ", juego.palabra)
			break
		} else if juego.HaGanado() {
			fmt.Println("Has ganado! La palabra era: ", juego.palabra)
			break
		}
	}
}
