package main

import (
	"fmt"
)

func HashFunction(item int, size int) int {
	return item % size
}

func isUnique(element string) bool {
	size := len(element)
	if size == 0 {
		fmt.Printf("This is empty!\n")
		return false
	}
	m := make(map[int][]byte)
	for i := 0; i < size; i++ {
		index := HashFunction(int(element[i]), size)
		if len(m[index]) > 0 { // colision and if value exists
			for _, s := range m[index] {
				if s == element[i] {
					fmt.Printf("Not unique!\n")
					return false
				}
			}
			m[index] = append(m[index], element[i])
		} else {
			m[index] = append(m[index], element[i])
		}
	}
	fmt.Println(m)
	fmt.Printf("This is unique!\n")
	return true
}
func main() {
	example3 := "abcdeab"
	isUnique(example3)
	example1 := "abcde"
	isUnique(example1)
	example2 := ""
	isUnique(example2)
	example4 := "abcdefghijklmnopqrstuvxyz"
	fmt.Println(isUnique(example4))

}
