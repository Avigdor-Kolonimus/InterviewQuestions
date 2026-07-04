package main

import (
	"context"
)

// Implement a function that performs
// searching for `query` across all provided SearchFunc functions.
// As soon as we get the first successful result,
// return it immediately.
// If all SearchFunc calls fail,
// return the last received error.

type Result struct{}

type SearchFunc func(ctx context.Context, query string) (Result, error)

func MultiSearch(ctx context.Context, query string, sfs []SearchFunc) (Result, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	type resp struct {
		res Result
		err error
	}

	ch := make(chan resp, len(sfs))

	for _, sf := range sfs {
		sf := sf
		go func() {
			res, err := sf(ctx, query)
			ch <- resp{res: res, err: err}
		}()
	}

	var lastErr error

	for i := 0; i < len(sfs); i++ {
		select {
		case <-ctx.Done():
			return Result{}, ctx.Err()
		case r := <-ch:
			if r.err == nil {
				cancel()
				return r.res, nil
			}
			lastErr = r.err
		}
	}

	return Result{}, lastErr
}
