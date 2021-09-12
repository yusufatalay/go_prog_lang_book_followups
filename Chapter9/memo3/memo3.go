// Package memo provides a concurrency-unsafe
// memoization of a function of type Func
package memo3

import "sync"

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	// if there is a cache hit, then unlock the goroutine so others can use the Get()
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the two critical sections, several gorourines
		// may race to compute f(key) and update the map.
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
