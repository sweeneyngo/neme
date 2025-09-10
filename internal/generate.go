package internal

import (
	"math/rand"
	"strings"
)

func pickNext(key string, chain Chain) string {
	transitions, ok := chain[key]
	if !ok || len(transitions) == 0 {
		return "$"
	}

	total := 0.0
	for _, t := range transitions {
		total += t.Weight
	}

	rnd := rand.Float64() * total
	for _, t := range transitions {
		rnd -= t.Weight
		if rnd <= 0 {
			return t.Next
		}
	}

	return transitions[len(transitions)-1].Next
}

func GenerateName(chain Chain, minLen, maxLen int, weightedStarterKeys []string) string {
	for {
		key := PickStarterKey(weightedStarterKeys)
		var name strings.Builder
		name.WriteString(key)

		for {
			next := pickNext(key, chain)
			if next == "$" || name.Len() >= maxLen {
				break
			}
			name.WriteString(next)
			key = key[1:] + next
		}

		if name.Len() >= minLen {
			return name.String()
		}
	}
}
