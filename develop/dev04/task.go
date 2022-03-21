package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func Unique(input string) bool {
	if utf8.RuneCountInString(input) < 2 {
		return true
	}

	input = strings.ToLower(input)
	m := make(map[string]int)
	for _, b := range input {
		m[string(b)] += 1
		if v, ok := m[string(b)]; ok {
			if v > 1 {
				return false
			}
		}
	}

	return true
}

func Set(sl []string) []string  {
	m := make(map[string]int)
	set := make([]string, 0)
	for _, v := range sl {
		m[v] += 1
	}

	for k := range m {
		if m[k] < 2 {
			set = append(set, k)
		}
	}

	return set
}

func IsAnagram(self, other string) bool {
	if utf8.RuneCountInString(self) != utf8.RuneCountInString(other) {
		return false
	}
	m := make(map[string]int)

	for _, v := range strings.Trim(strings.ToLower(self), " ") {
		m[string(v)]++
	}

	for _, v := range strings.Trim(strings.ToLower(other), " ")  {
		m[string(v)] = m[string(v)] - 1
		if m[string(v)] < 0 {
			return false
		}
	}

	return true
}

func SetAnagram(anagrams []string) map[string][]string {
	m := make(map[string][]string)
	for _, anagram := range anagrams {
		s := make([]string, 0)
		for _, word := range anagrams {
			if IsAnagram(anagram, word) {
				s = append(s, strings.ToLower(word))
				if len(s) > 1 {
					m[s[0]] = s[1:]
				}
			}
		}
	}

	for _, a := range m {
		sort.Strings(a)
	}

	return m
}

func main() {
	anagrams := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}
	fmt.Println(SetAnagram(anagrams))
}