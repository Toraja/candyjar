# Go

## Requirements

- [GoNB](https://github.com/janpfeifer/gonb)
  - Go kernel for Jupyter Notebook
  - The code blocks do not automatically continue from the previous blocks so
    each block has its own scope and variables, functions and etc from the
    earlier blocks cannot be refernced in the later ones.  
    Use `%%` to denote the block is continuation of the previous block.
- [pkgsite](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite)
  - View snippets like Godoc
