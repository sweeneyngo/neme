package internal

import "strings"

// Character sets
var vowels = "aeiouAEIOU"
var consonants = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"

func isVowel(r rune) bool {
	return strings.ContainsRune(vowels, r)
}

func isConsonant(r rune) bool {
	return strings.ContainsRune(consonants, r)
}
