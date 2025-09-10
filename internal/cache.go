package internal

import (
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CacheFileName(corpusPath string, order int) string {
	data, err := os.ReadFile(corpusPath)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(data)
	shortHash := hex.EncodeToString(hash[:8]) // just 8 bytes for readability
	cacheDir := ".cache"

	// Ensure cache dir exists
	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		log.Fatal(err)
	}

	return filepath.Join(cacheDir, fmt.Sprintf("cache-order%d-%s.gob", order, shortHash))
}

func LoadCache(path string) (CacheData, bool) {
	f, err := os.Open(path)
	if err != nil {
		return CacheData{}, false
	}
	defer f.Close()

	var data CacheData
	dec := gob.NewDecoder(f)
	if err := dec.Decode(&data); err != nil {
		return CacheData{}, false
	}
	return data, true
}

func SaveCache(path string, data CacheData) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	return enc.Encode(data)
}
