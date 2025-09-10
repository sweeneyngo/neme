package main

import "math/rand"

type BigramMap [128][128]int
type Chain map[string][]Transition

type Transition struct {
	Next   string
	Weight float64
}

var P_CC = 0.3868
var P_CV = 0.6132
var P_VC = 0.8588
var P_VV = 0.1412

// Constructs bigram counts for all words in the corpus.
func BuildBigramCounts(words []string) BigramMap {
	var counts BigramMap
	for _, word := range words {
		runes := []rune(word + "$")
		for i := 0; i < len(runes)-1; i++ {
			prev, next := runes[i], runes[i+1]
			if prev < 128 && next < 128 { // bounds check
				counts[prev][next]++
			}
		}
	}
	return counts
}

// Builds a Markov chain with precomputed transition weights.
// words: corpus, order: chain order (2–3 supported easily).
func BuildChain(words []string, order int, bigramCounts BigramMap) Chain {
	raw := make(map[string]map[string]int)

	// Step 1: Build raw counts
	for _, word := range words {
		runes := []rune(word + "$")
		for i := 0; i <= len(runes)-order; i++ {
			key := string(runes[i : i+order])
			next := string(runes[i+order : i+order+1])
			if raw[key] == nil {
				raw[key] = make(map[string]int)
			}
			raw[key][next]++
		}
	}

	// Step 2: Convert to optimized format with precomputed weights
	chain := make(Chain)
	for key, nextMap := range raw {
		last := rune(key[len(key)-1])

		maxCount := 1
		for _, cnt := range bigramCounts[last] {
			if cnt > maxCount {
				maxCount = cnt
			}
		}

		var transitions []Transition
		for nextStr := range nextMap {
			r := []rune(nextStr)[0]

			cvWeight := 0.3
			if nextStr == "$" {
				cvWeight = 1.0
			} else if isConsonant(last) && isConsonant(r) {
				cvWeight = P_CC
			} else if isConsonant(last) && isVowel(r) {
				cvWeight = P_CV
			} else if isVowel(last) && isConsonant(r) {
				cvWeight = P_VC
			} else if isVowel(last) && isVowel(r) {
				cvWeight = P_VV
			}

			bigramCnt := 1
			if bigramCounts[last][r] > 0 {
				bigramCnt = bigramCounts[last][r]
			}

			w := (float64(bigramCnt) / float64(maxCount)) * cvWeight
			transitions = append(transitions, Transition{Next: nextStr, Weight: w})
		}

		chain[key] = transitions
	}

	return chain
}

// Computes weighted starter keys based on word frequency.
func BuildStarterKeys(words []string, order int) []string {
	starterFreq := make(map[string]int)
	for _, word := range words {
		if len(word) >= order {
			key := word[:order]
			starterFreq[key]++
		}
	}

	var weightedKeys []string
	for key, freq := range starterFreq {
		if freq < 2 {
			continue
		}
		for i := 0; i < freq; i++ {
			weightedKeys = append(weightedKeys, key)
		}
	}

	return weightedKeys
}

func PickStarterKey(weightedKeys []string) string {
	if len(weightedKeys) == 0 {
		return ""
	}
	return weightedKeys[rand.Intn(len(weightedKeys))]
}
