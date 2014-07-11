package main

// this passes on 1.3 if changed to var OR the length of
// the chars string is 1,2,4,7,8  fails being var with len 3,5,6,9
const chars = "ABCD"

// makeString generates a string consisting only of letters from chars
func makeString(length int) string {
	bytes := make([]byte, length)
	for i, _ := range bytes {
		bytes[i] = byte(i)
	}
	for i, b := range bytes {
		bytes[i] = chars[b%byte(len(chars))]
	}
	return string(bytes)
}

func main() {
}
