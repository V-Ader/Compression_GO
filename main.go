package main

import (
	"fmt"

	"compression/file"
	"compression/tree"
)

var SOURCE_FILE = "data/data.txt"
var TARGET_FILE = "out/cyphered"
var LEFT_BYTE = 48
var RIGHT_BYTE = 49

func code(data []byte, tree tree.Tree) string {
	ret := ""
	preCalculated := map[byte]string{}
	for _, letter := range data {
		_, precalculated := preCalculated[letter]
		if !precalculated {
			newCode := tree.GetCode(letter)
			preCalculated[letter] = newCode
		}
		ret += preCalculated[letter]
	}
	return ret
}

func decode(data string, code tree.Tree) string {
	return code.GetByCode(data)
}

func textToBinary(input []byte) string {
	var result string
	for _, b := range input {
		result += fmt.Sprintf("%08b", b)
	}
	return result
}

func main() {
	data := file.Load(SOURCE_FILE, false)
	tree := tree.CreateTree(data)
	fmt.Printf("data: %s \n\n", data)

	cyphered := code(file.Load(SOURCE_FILE, false), tree)
	file.Save(TARGET_FILE, cyphered)
	fmt.Printf("cyphered data: %s \n\n", cyphered)

	decoded := decode(cyphered, tree)
	fmt.Printf("decoded data: %s \n\n", decoded)

	len_cyphered := len(cyphered)
	len_oryginal := len(textToBinary(data))
	fmt.Printf("Compression: %d/%d = %f%% \n", len_cyphered, len_oryginal, float32(len_cyphered)/float32(len_oryginal))
}
