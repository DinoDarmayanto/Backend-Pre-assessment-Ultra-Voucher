package main

import (
	"fmt"
	"sort"
)

// fungsi untuk mengurutkan karakter pada string
func sortString(s string) string {
	characters := []rune(s)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}

// fungsi untuk mengelompokkan array of strings menjadi kumpulan anagram
func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		// mengurutkan karakter dalam string untuk mencari anagram
		sortedStr := sortString(str)

		// tambahkan string ke dalam map dengan kunci berupa sortedStr
		anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
	}

	// konversi map menjadi slice dari slice untuk hasil akhir
	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func main() {
	// array of strings yang diberikan
	input := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}

	// panggil fungsi untuk mengelompokkan anagram
	result := groupAnagrams(input)

	// print hasil
	fmt.Println(result)
}
