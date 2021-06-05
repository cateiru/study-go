# Go Tutorial

## Getting Started

### Install Go

pass

[goenv](https://github.com/syndbg/goenv)使用

### Create a Go module

```bash
go mod init [name]
```

### Writing Web App

- [source](../gowiki/wiki.go)

```bash
cd gowiki
go build wiki.go
./wiki
```

#### Data Structure

```go
type Page struct {
    Title string
}

func (p* Page) save() error {
    title := p.Title
}
```
