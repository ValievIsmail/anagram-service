package service

import (
	"strings"
)

// SearchAnagrams ищем анаграму в словаре
func SearchAnagrams(dict []string, word string) (res []string) {
	for i := 0; i < len(dict); i++ {
		if IsAnagram(dict[i], word) {
			res = append(res, dict[i])
		}
	}
	return res
}

// IsAnagram проверяем слово на признак анаграмы
func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int, len(s1))
	m2 := make(map[rune]int, len(s2))

	for r := range m1 {
		m1[r]++
	}

	for r := range m2 {
		m2[r]++
	}

	for i := range m1 {
		if m1[i] != m2[i] {
			return false
		}
	}

	return true
}

// DictToLowerCase переводим массив в lowercase
func DictToLowerCase(dict []string) {
	for i, v := range dict {
		dict[i] = strings.ToLower(v)
	}
}
