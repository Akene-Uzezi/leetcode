package main

func firstPalindrome(words []string) *string {
	return nil
}

func isPalindrome(s string) *bool {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	compare := s == string(runes)
	return &compare
}

func main() {}
