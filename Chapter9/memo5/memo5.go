// Package memo provides a concurrency-unsafe
// memoization of a function of type Func
// contains exercise 9.3
package memo5

import "log"

// A request is message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result // the client wants a single result
}

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// done channel will be used for cancellling
// Func is the type of the function to memoize.
type Func func(key string, done chan struct{}) (interface{}, error)

type Memo struct{ requests chan request }

// New returns a memoization of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	select {
	case <-done:
		log.Print("Cancelled")
	default:
		memo.requests <- request{key, response}
	}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Evaluate the function
	e.res.value, e.res.err = f(key)
	// Broadcast the ready condition
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client
	response <- e.res
}
func (memo *Memo) Close() { close(memo.requests) }
