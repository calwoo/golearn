package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, Ø¯Ù†ÙŠØ§!")
	fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))
	fmt.Println(Anagram("windfury", "furydinw"))
	fmt.Println(Anagram("cat", "dog"))
	fmt.Println(FindEvens([]int{2, 3, 5, 6, 7, 8}))
	fmt.Println(SliceProduct([]int{1, 2, 3}))
	fmt.Println(Unique([]int{1, 1, 1, 3, 5, 3, 1}))
	fmt.Println(InvertMap(map[string]int{"henry": 4, "jim": 5}))
	fmt.Println(TopCharacters("mississippi", 2))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	cleanPhone := strings.ReplaceAll(phone, " ", "")
	cleanPhone = strings.ReplaceAll(cleanPhone, "-", "")

	fmtPhone := fmt.Sprintf("(%s) %s-%s",
		cleanPhone[:3],
		cleanPhone[3:6],
		cleanPhone[6:])
	return fmtPhone
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	s1Slice := strings.Split(s1, "")
	s2Slice := strings.Split(s2, "")

	sort.Strings(s1Slice)
	sort.Strings(s2Slice)
	return strings.Join(s1Slice, "") == strings.Join(s2Slice, "")
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	// TODO
	var justEvens []int
	for _, i := range e {
		if i%2 == 0 {
			justEvens = append(justEvens, i)
		}
	}
	return justEvens
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	var product int = 1
	for _, i := range e {
		product *= i
	}
	return product
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var uniqueElems []int
	for _, i := range e {
		var flag bool = true
		for _, j := range uniqueElems {
			if i == j {
				flag = false
			}
		}

		if flag {
			uniqueElems = append(uniqueElems, i)
		}
	}
	return uniqueElems
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	var inverted map[int]string = make(map[int]string)
	for k, v := range kv {
		inverted[v] = k
	}
	return inverted
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	var counts map[rune]int = make(map[rune]int)
	for _, s := range s {
		c, ok := counts[s]
		if !ok {
			counts[s] = 1
		} else {
			counts[s] = c + 1
		}
	}

	// filter out
	var topCounts map[rune]int = make(map[rune]int)
	for r, c := range counts {
		if c > k {
			topCounts[r] = c
		}
	}
	return topCounts
}
