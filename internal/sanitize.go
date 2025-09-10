package internal

import "strings"

func CleanWord(word string) string {
	var sb strings.Builder
	for _, r := range word {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
