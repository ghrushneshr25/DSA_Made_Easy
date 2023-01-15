package Backtracking

import (
	"strings"
	"fmt"
)

type KAryString struct {
	Array []string
}

func InitiateKaryString(n int, K []string){
	var kary KAryString
	kary.Array = make([]string, n)
	kary.GenerateString(n, K)
}

func (kary *KAryString) GenerateString(n int, K []string){
	if n <= 0 {
		fmt.Println(strings.Join(kary.Array, ""))
		return
	} else {
		for _, val := range K{
			kary.Array[n-1] = val
			kary.GenerateString(n-1, K)
		}
	}
}

