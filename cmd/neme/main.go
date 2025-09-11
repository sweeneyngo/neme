package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"neme/internal"
	"os"
	"runtime"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "\nGenerates random words for your in-game naming needs.")
		fmt.Fprintln(os.Stderr, "Options:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExample:")
		fmt.Fprintf(os.Stderr, "  %s --corpus /usr/share/dict/words --order 2 --min 3 --max 8 --count 10\n", os.Args[0])
	}

	corpusPath := flag.String("corpus", "/usr/share/dict/words", "Path to word corpus")
	order := flag.Int("order", 2, "Markov chain order (number of letters in key)")
	lengthMin := flag.Int("min", 3, "Minimum length of generated names")
	lengthMax := flag.Int("max", 8, "Maximum length of generated names")
	count := flag.Int("count", 10, "Number of names to generate")
	flag.Parse()

	file, err := os.Open(*corpusPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	title := cases.Title(language.English)
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		word = internal.CleanWord(word)
		if len(word) >= *lengthMin && len(word) <= *lengthMax {
			words = append(words, title.String(strings.ToLower(word)))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if len(words) == 0 {
		log.Fatal("No words found in corpus with given length constraints")
	}

	cachePath := internal.CacheFileName(*order)

	var chain internal.Chain
	var bigramCounts internal.BigramMap
	var weightedStarterKeys []string

	if cached, ok := internal.LoadCache(cachePath); ok {
		chain = cached.Chain
		weightedStarterKeys = cached.WeightedStarterKeys
	} else {
		bigramCounts = internal.BuildBigramCounts(words)
		chain = internal.BuildChain(words, *order, bigramCounts)
		weightedStarterKeys = internal.BuildStarterKeys(words, *order)
		data := internal.CacheData{
			Chain:               chain,
			WeightedStarterKeys: weightedStarterKeys,
		}
		if err := internal.SaveCache(cachePath, data); err != nil {
			log.Printf("Warning: failed to save cache: %v\n", err)
		}
	}

	// Collect results
	numWorkers := runtime.NumCPU()
	names := internal.Generate(*count, chain, *lengthMin, *lengthMax, weightedStarterKeys, numWorkers)
	for _, sn := range names {
		fmt.Println(sn)
	}
}
