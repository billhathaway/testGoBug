package main

import "fmt"
import "crypto/rand"

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func makeRandomMessageV1(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		index := b % byte(len(chars)) // this line is a no-op
		_ = index                     // since we throw it away
		bytes[i] = chars[b%byte(len(chars))]
	}
	return string(bytes)
}
func makeRandomMessageV2(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	return string(bytes)
}

func main() {
	fmt.Printf("v1=[%x]\n", makeRandomMessageV1(10))
	fmt.Printf("v2=[%x]\n", makeRandomMessageV2(10))
}
