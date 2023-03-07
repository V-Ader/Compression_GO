package tree

import (
	"math"
	"sort"
)

type Leaf struct {
	value       []byte
	propability float64
	codeElement byte
	left        *Leaf
	right       *Leaf
}

func (leaf *Leaf) GetByCode(code *string, depth int) (byte, int) {

	if leaf.left == nil && leaf.right == nil {
		return leaf.value[0], depth
	}

	if (*code)[depth] == byte('0') {
		if leaf.left != nil {
			return leaf.left.GetByCode(code, depth+1)
		} else {
			return leaf.value[0], depth
		}
	} else {
		if leaf.right != nil {
			return leaf.right.GetByCode(code, depth+1)
		} else {
			return leaf.value[0], depth
		}
	}
}

func (leaf *Leaf) GetCode(sign byte) []byte {

	if leaf == nil {
		return nil
	}

	if len(leaf.value) == 1 && leaf.value[0] == sign {
		return []byte{leaf.codeElement}
	}

	if leaf.left != nil {
		left := leaf.left.GetCode(sign)

		if len(left) > 0 {
			ret := []byte{leaf.codeElement}
			return append(ret, left...)
		}
	}

	if leaf.right != nil {
		right := leaf.right.GetCode(sign)
		if len(right) > 0 {
			ret := []byte{leaf.codeElement}
			return append(ret, right...)
		}
	}
	return nil
}

type Tree struct {
	head Leaf
}

func (tree Tree) GetByCode(code string) string {
	elements := []byte{}
	index := 0
	for index < len(code) {
		currentCode := code[index:]
		letter, increment := tree.head.GetByCode(&currentCode, 1)
		elements = append(elements, letter)
		index += increment
	}
	return string(elements[:])
}

func (tree Tree) GetCode(sign byte) string {
	return string(tree.head.GetCode(sign))
}

func PrintTree(leaf *Leaf, depth int) {
	if leaf == nil {
		return
	}
	println("\n\nNewNode")
	for _, el := range leaf.value {
		print(el, " ")
	}

	if leaf.left != nil {
		println("\n left:")
		for _, el := range leaf.left.value {
			print(el, " ")
		}
	}
	if leaf.right != nil {
		println("\n right:")
		for _, el := range leaf.right.value {
			print(el, " ")
		}
	}
	PrintTree(leaf.left, depth+2)
	PrintTree(leaf.right, depth+2)
}

func CreateTree(data []byte) Tree {
	props := getPropabiliMap(data)

	var coreLeaves []Leaf
	for key, value := range props {
		coreLeaves = append(coreLeaves, Leaf{
			value:       []byte{key},
			propability: value,
		})
	}

	leaves := coreLeaves[:]
	sort.SliceStable(leaves, func(i, j int) bool {
		if math.Abs(leaves[i].propability-leaves[j].propability) < 0.0000001 {
			return leaves[i].value[0] < leaves[j].value[0]
		}
		return leaves[i].propability < leaves[j].propability
	})

	for len(leaves) >= 2 {
		newLeaves := leaves[2:]
		newLeaf := combineLeaves(leaves[0], leaves[1])

		newIndex := findNewIndex(newLeaf.propability, newLeaves)
		leaves = insert(newLeaves, newIndex, newLeaf)
	}

	return Tree{leaves[0]}
}

func getPropabiliMap(data []byte) map[byte]float64 {
	count_all := float64(len(data))
	occurrences := make(map[byte]float64)
	for i := range data {
		occurrences[data[i]]++
	}
	for key := range occurrences {
		occurrences[key] /= count_all
	}
	return occurrences
}

func findNewIndex(value float64, currentLeaves []Leaf) int {
	for id, leaf := range currentLeaves {
		if value <= leaf.propability {
			return id
		}
	}
	return len(currentLeaves)
}

func insert(a []Leaf, index int, value Leaf) []Leaf {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func combineLeaves(left Leaf, right Leaf) Leaf {
	left.codeElement = '0'
	right.codeElement = '1'
	newValueFromLeft := []byte{}
	newValueFromLeft = append(newValueFromLeft, left.value...)
	newValueFromLeft = append(newValueFromLeft, right.value...)

	return Leaf{
		newValueFromLeft,
		left.propability + right.propability,
		0,
		&left,
		&right,
	}
}
