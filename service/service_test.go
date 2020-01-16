package service

import (
	"reflect"
	"testing"
)

func TestAnagram(t *testing.T) {
	standartDict := []string{"foobar", "aabb", "baba", "boofar", "test", "живу", "вижу"}
	expected := []string{"foobar", "boofar"}

	case1 := SearchAnagrams(standartDict, "foobar")
	if !reflect.DeepEqual(case1, expected) {
		t.Errorf("Failed, expected %s, got %s", expected, case1)
	}

	lowerCaseTest := []string{"FOOBAR", "barfoo"}
	DictToLowerCase(lowerCaseTest)

	expected2 := []string{"foobar", "barfoo"}
	case2 := SearchAnagrams(lowerCaseTest, "barfoo")

	if !reflect.DeepEqual(case2, expected2) {
		t.Errorf("Failed, expected %s, got %s", expected2, case2)
	}

	single := []string{"test"}
	expected3 := []string{"test"}

	case3 := SearchAnagrams(single, "test")
	if !reflect.DeepEqual(case3, expected3) {
		t.Errorf("Failed, expected %s, got %s", expected3, case3)
	}
}

func TestIsAnagram(t *testing.T) {
	item := ConvertStrToMap("silent")
	word := ConvertStrToMap("listen")

	if !IsAnagram(item, word) {
		t.Errorf("Failed, expected %t, got %t", true, false)
	}
}

func TestDictToLowerCase(t *testing.T) {
	dict := []string{"FOOBAR"}
	DictToLowerCase(dict)

	for _, v := range dict {
		if v == "FOOBAR" {
			t.Errorf("Failed, expected %s, got %s", "foobar", v)
		}
	}
}
