# neme

A **Markov-chain based name generator**.
Create unique names for characters, pets, guilds, or worlds — without needing to brainstorm for hours.

neme uses Markov chains + weighted generation to produce a natural stream of names.

Benchmarking was performed using a corpus of ~4M characters (e.g., `/usr/share/dict/words`) on a modern CPU, with chain order 2. The generator was run with 16 goroutines (parallel workers) to measure average generation times. We run it for 100 iterations.

| Count     | Goroutines | Total Time (µs) | Avg Time per Name (µs) | StdDev per Name (µs) |
|-----------|------------|----------------|------------------------|---------------------|
| 100       | 16         | 115.19          | 1.15                   | 0.35               |
| 1,000     | 16          | 131.58           | 1.32                   | 0.23                |
| 10,000    | 16          | 116.37          | 1.16                   | 0.15                |

**Notes:**

- **Count** = number of names generated in a single batch.
- **Goroutines** = number of concurrent workers used for generation.
- **Total Time** = wall-clock time measured for the batch.
- **Avg Time per Name** = `Total Time ÷ Count`.

---

## Installation

Download the latest release from [GitHub Releases](https://github.com/sweeneyngo/neme/releases)
or build from source:

```bash
git clone https://github.com/sweeneyngo/neme.git
cd neme
go build ./cmd/neme
./neme --help
```

## Usage

```bash
Usage: ./neme [options]

Generates random words for your in-game naming needs.

Options:
  -corpus string
        Path to word corpus (default "/usr/share/dict/words")
  -count int
        Number of names to generate (default 10)
  -max int
        Maximum length of generated names (default 8)
  -min int
        Minimum length of generated names (default 3)
  -order int
        Markov chain order (number of letters in key) (default 2)

Example:
  ./neme --corpus /usr/share/dict/words --order 2 --min 3 --max 8 --count 10
```

This produces the `neme` binary in the current directory.

```bash
$ ./neme --count 5
Ralven
Quira
Monel
Tirath
Velora
```

## Building

### Prerequisites

- [Go 1.24](https://go.dev/dl/) or later.

```bash
git clone https://github.com/sweeneyngo/neme.git
cd neme
go build ./cmd/neme
```
