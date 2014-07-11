// rand_test.go
package main

import (
	"testing"
	"unicode"
)

func TestV1vsV2(t *testing.T) {
	for i := 0; i < 10; i++ {
		v1 := makeRandomMessageV1(10)
		v2 := makeRandomMessageV2(10)
		for _, char := range v1 {
			if !unicode.IsUpper(char) {
				t.FailNow()
			}
		}
		for _, char := range v2 {
			if !unicode.IsUpper(char) {
				t.FailNow()
			}
		}
	}
}
