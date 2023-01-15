package Backtracking

import (
	"fmt"
	"strings"
)

type BinaryString struct {
	Array []string
}

func InitiateBinaryString(n int){
	var binary BinaryString
	binary.Array = make([]string,n)
	binary.GenerateStrings(n);
}

func (b *BinaryString) GenerateStrings(n int){
	if n <= 0 {
		fmt.Println(strings.Join(b.Array, ""))
	} else {
		b.Array[n-1] = "0"
		b.GenerateStrings(n-1)
		b.Array[n-1] = "1"
		b.GenerateStrings(n-1)
	}
}
