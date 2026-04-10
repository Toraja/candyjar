---
jupyter:
  jupytext:
    text_representation:
      extension: .md
      format_name: markdown
      format_version: '1.3'
      jupytext_version: 1.19.1
  kernelspec:
    display_name: Go (gonb)
    language: go
    name: gonb
---

# Go Iterators with `iter.Seq` and `yield`

This notebook explores how Go's `iter` package enables generator-style iterators using `yield` functions.


## 1. Typical Generator: `numbers`

This is the canonical way to write a generator in Go.

The `yield` function:
- Returns `true` while the caller loop is still iterating
- Returns `false` when the caller loop has stopped (via `break` or `return`)
- **Blocks** until the next iteration of the caller loop

So the output interleaves pre-yield, yielded, and post-yield messages in lock-step.

```go
import (
	"fmt"
	"iter"
)

// numbers returns an iter.Seq[int] that yields integers from 0 to n-1.
// It checks the return value of yield so it can stop early when the caller breaks.
func numbers(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			fmt.Println("pre-yield:", i)
			if !yield(i) {
				return // caller broke out of the loop
			}
			fmt.Println("post-yield:", i)
		}
	}
}
```

```go
%%
// Break at n==4: we should see pre/post-yield for 0..3, then only pre-yield for 4.
for n := range numbers(7) {
	if n == 4 {
		break
		// `return` does the same, except it stops the entire enclosing function
		// instead of just the loop.
	}
	fmt.Println("yielded:", n)
}
```

## 2. Dangerous Generator: `numbersNoCheckYield`

Calling `yield` **without** checking its return value causes a panic when the caller breaks:

> `runtime error: range function continued iteration after function for loop body returned false`

The cell below defines the function but does **not** run it to avoid crashing the kernel.

```go
import "iter"

// numbersNoCheckYield panics if the caller breaks early,
// because yield is called even after the loop has stopped.
func numbersNoCheckYield(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			fmt.Println("pre-yield:", i)
			yield(i) // ⚠️ return value ignored — panics on early break!
			fmt.Println("post-yield:", i)
		}
	}
}
```

```go
%%
// panics after `pre-yield: 4`. You will see the output beyond that because generator is called asynchronously.
for n := range numbersNoCheckYield(7) {
	if n == 4 {
		break
	}
	fmt.Println("yielded:", n)
}
```

```go
%%
// It works fine if the caller loop does not exit early.
for n := range numbersNoCheckYield(7) {
	fmt.Println("yielded:", n)
}
```

## 3. No-op Generator: `numbersNoYield`

Not calling `yield` at all is perfectly safe — no panic, no error.

It simply means the caller loop body **never executes**: the generator runs its own loop
and exits without producing any values for the caller.

```go
import "iter"

// numbersNoYield never calls yield, so the caller loop body never runs.
// The generator still executes its own loop normally.
func numbersNoYield(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := range n {
			fmt.Println("pre-yield:", i)
			fmt.Println("post-yield:", i)
			// yield is never called — caller loop body will never execute
		}
	}
}
```

```go
%%
// The loop body (fmt.Println("yielded:", n)) will never be reached.
for n := range numbersNoYield(7) {
	if n == 4 {
		break
	}
	fmt.Println("yielded:", n)
}
```

## Summary

| Function | Checks `yield` return | Caller can `break` safely | Values produced |
|---|---|---|---|
| `numbers` | Yes | Yes | Yes |
| `numbersNoCheckYield` | No | **No — panics** | Yes |
| `numbersNoYield` | N/A | Yes | **No** |


## Extra

Iterator can be used outside loops as well.

```go
%%
yield := func(n int) bool {
    fmt.Println("yielded:", n)
    if n == 3 {
        return false
    }
    return true
}

gen := numbers(5)
gen(yield)
```
