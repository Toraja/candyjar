# Go

## Requirements

- [GoNB](https://github.com/janpfeifer/gonb)
  - Go kernel for Jupyter Notebook
  - GoNB requires `main` function to be present in each cell
    - `%%` is a special commands that means "insert a `func main {...}` here", so below two codes are equivalent
      ```go
      func main() {
          fmt.Println("hi")
      }
      ```
      ```go
      %%
      fmt.Println("hi")
      ```
  - It guesses `import` statements whenever possible
    ```go
    %%
    fmt.Println("No import statement is necessary")
    ```
  - See [tutorial](https://github.com/janpfeifer/gonb/blob/main/examples/tutorial.ipynb) for more details
- [pkgsite](https://pkg.go.dev/golang.org/x/pkgsite/cmd/pkgsite)
  - View snippets like Godoc
