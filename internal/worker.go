package internal

func Generate(
	count int,
	chain Chain,
	minLen, maxLen int,
	weightedStarterKeys []string,
	numWorkers int,
) []string {
	results := make(chan string, count)
	jobs := make(chan int, count)

	worker := func() {
		for range jobs {
			name := GenerateName(chain, minLen, maxLen, weightedStarterKeys)
			results <- name
		}
	}

	for range numWorkers {
		go worker()
	}

	for i := range count {
		jobs <- i
	}
	close(jobs)

	var names []string
	for range count {
		r := <-results
		names = append(names, r)
	}

	return names
}
