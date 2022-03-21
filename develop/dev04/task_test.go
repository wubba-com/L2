package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetAnagrams(t *testing.T)  {
	anagrams := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "Лунь", "нуль", "горечь"}

	testM := make(map[string][]string)
	testM["листок"] = append(testM["листок"], "слиток", "столик")
	testM["лунь"] = append(testM["лунь"], "нуль")
	testM["пятак"] = append(testM["пятак"], "пятка", "тяпка")

	m := SetAnagram(anagrams)
	fmt.Println(m)
	for k := range m {
		if ok := reflect.DeepEqual(testM[k], m[k]); !ok {
			t.Error("result != value")
		}
	}
}
