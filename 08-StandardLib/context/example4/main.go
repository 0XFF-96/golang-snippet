package main

import (
	"context"
	"fmt"
)

// This different-struct-memory demonstrates how a value can be passed to the context
// and also how to retrieve it if it exists.
func ExampleWithValue() {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))

	// Output:
	// found value: Go
	// key not found: color
}
