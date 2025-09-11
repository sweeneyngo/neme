package internal

import (
	"encoding/gob"
	"fmt"
	"os"
)

func CacheFileName(order int) string {
	return fmt.Sprintf(".cache/cache-order%d.gob", order)
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
