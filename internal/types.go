package internal

type BigramMap [128][128]int

type CacheData struct {
	Chain               Chain
	WeightedStarterKeys []string
}

type Chain map[string][]Transition

type Transition struct {
	Next   string
	Weight float64
}
