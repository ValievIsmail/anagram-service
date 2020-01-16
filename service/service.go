package service

import (
	"reflect"
	"strings"
)

// SearchAnagrams ищем анаграму
func SearchAnagrams(dict []string, word string) (res []string) {
	wordRunes := ConvertStrToMap(word)

	for i := 0; i < len(dict); i++ {
		item := ConvertStrToMap(dict[i])

		if IsAnagram(item, wordRunes) {
			res = append(res, dict[i])
		}
	}

	return res
}

// IsAnagram проверяем слово на признак анаграмы
func IsAnagram(item, word map[rune]int) bool {
	return reflect.DeepEqual(item, word)
}

// ConvertStrToMap конвертим строку в мапу
func ConvertStrToMap(s string) (m map[rune]int) {
	m = make(map[rune]int, len(s))

	for _, r := range s {
		m[r]++
	}

	return m
}

// DictToLowerCase переводим массив в lowercase
func DictToLowerCase(dict []string) {
	for i, v := range dict {
		dict[i] = strings.ToLower(v)
	}
}
