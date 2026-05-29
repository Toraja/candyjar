# inline

## About

`go fix inline` replaces function calls with their body if the functions are annotated with `//go:fix inline`.  
A use case is to replace deprecated functions with new ones automatically.  
See [the blog page](https://go.dev/blog/inliner) for more detail.

## Example

Run `go fix inline` to see how the inlining works.  
You can also run `go fix inline -diff` to see the diff of the changes without modifying the source code.
