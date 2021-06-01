# How To Use

## Create a project

```bash
go mod init [package name]
```

## Hello World

- `hello.go`

  ```go
  package main
  import "fmt"

  func main() {
    fmt.Println("Hello World")
  }
  ```

```bash
# Quick run
go run .

# Build and run
go build
./hello
```
