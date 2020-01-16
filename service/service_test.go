package service

import (
	"reflect"
	"testing"
)

// AnagramTesting тестим разные кейсы
func TestAnagram(t *testing.T) {
	dict := []string{"foobar", "aabb", "baba", "boofar", "test"}
	expected := []string{"foobar", "boofar"}

	res1 := SearchAnagrams(dict, "foobar")
	if !reflect.DeepEqual(res1, expected) {
		t.Errorf("Failed, expected %s, got %s", expected, res1)
	}
}
