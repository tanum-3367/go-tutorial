package main

const GlobalLimt = 100
const MaxCacheSize int = 10 * GlobalLimt
const (
	CachedKeyBook = "book_"
	CachedKeyCD   = "cd_"
)

var cache map[string]string

func cacheGet(key string) string {
	return cache[key]
}

func cacheSet(key, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}
	cache[key] = val
}

func main() {
	cache = make(map[string]string)
}
