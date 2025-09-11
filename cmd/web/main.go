package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"neme/internal"
)

func main() {
	count := 20
	order := 2
	minLen, maxLen := 3, 8

	cachePath := internal.CacheFileName(order)
	cached, ok := internal.LoadCache(cachePath)
	if !ok {
		log.Fatalf("❌ Cache file %s not found. Please run the CLI tool to build it first.", cachePath)
	}

	chain := cached.Chain
	weightedStarterKeys := cached.WeightedStarterKeys

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numWorkers := runtime.NumCPU()
		names := internal.Generate(count, chain, minLen, maxLen, weightedStarterKeys, numWorkers)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(names)
	})

	log.Println("🌐 Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
