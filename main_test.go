// rand_test.go
package main

import (
	"testing"
	"unicode"
)

func TestLetters(t *testing.T) {
	v1 := makeString(10)
	for _, char := range v1 {
		if !unicode.IsLetter(char) {
			t.Fail()
		}
	}
}
