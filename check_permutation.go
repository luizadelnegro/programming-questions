package main

import "fmt"

func permutacao2(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	lens1 := len(s1)
	for i := 0; i < lens1; i++ {
		index1 := int(s1[i])
		index2 := int(s2[i])
		if m1[index1] == 0 {
			m1[index1] = 1
		} else {
			m1[index1] += 1
		}
		if m2[index2] == 0 {
			m2[index2] = 1
		} else {
			m2[index2] += 1
		}
	}

	for key, sum := range m1 {
		if m2[key] != sum {
			return false
		}
	}

	return true
}

func mergeSort(s string) []byte {
	if len(s) < 2 {
		return []byte(s)
	}
	firstPart := mergeSort(s[:len(s)/2])
	secondPart := mergeSort(s[len(s)/2:])
	return merge(firstPart, secondPart)
}

func merge(p1 []byte, p2 []byte) []byte { // merge two halves
	var result []byte
	i := 0
	j := 0
	for i < len(p1) && j < len(p2) {
		if p1[i] < p2[j] {
			result = append(result, p1[i])
			i++
		} else {
			result = append(result, p2[j])
			j++
		}
	}
	for ; i < len(p1); i++ {
		result = append(result, p1[i])
	}
	for ; j < len(p2); j++ {
		result = append(result, p2[j])
	}
	return result
}

func checkPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	ms1 := mergeSort(s1)
	ms2 := mergeSort(s2)
	for i := 0; i < len(s1); i++ {
		if ms1[i] != ms2[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := "zbac"
	s2 := "abaz"

	fmt.Println(checkPermutation(s1, s2))
}
