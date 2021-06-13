# Effective GO

- [https://golang.org/doc/effective_go](https://golang.org/doc/effective_go)

## フォーマット

- Gofmtを使用することでフォーマットが可能
- インデント
  - タブ
- Line length
  - 制限なし

## コメント

- `/**/`、`//`
- godoc
  - パッケージコメント: パッケージ上位にコメント。1つのファイルにのみ存在

    ```bash
    go doc -all regexp | grep -i parse
    ```

## 命名

- パッケージ名
  - 短く、簡潔
  - 小文字の一語
  - パッケージ名 = ソースディレクトリのベース名
- Getter
  - 自動getter、setterはサポート無し
  - 自分で作成する
- interface name
  - メソッド名+`er`サフィックス

## セミコロン

- 自動的に挿入する
- ワンライナーで書く場合は明示必要

```go
// ok
if i < f() {
    g()
}

// no: セミコロンが余計な位置についてしまう
if i < f()
{
    g()
}
```

## 制御構造

### if

```go
if x > 0 {
    return y
}
```

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

```go
f, err := os.Open(name)
if err != nil {
    return err
}
condeUsing(f)
```

```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

### 再宣言と再割り当て

```go
f, err := os.Open(name)
d, err := f.Star()
```

- 変数がすでに宣言されていても上書き可（err）

### for

- cでいうfor, while, do-whileをfor1つで可能

```go
// Like a C for
for init; condition; post {
    ...
}

// Like a C while
for condition {
    ...
}

// Like a C for(;;)
for {
    ...
}
```

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

- 配列のfor

    ```go
    // oldMap := [3]string{"hoge", "fuga", "foo"}
    for key, value := range oldMap {
        newMap[key] = value
    }
    ```

    ```go
    for key := range m {
        if key.expired() {
            delete(m, key)
        }
    }

    sum := 0
    for _, value := range array {
        sum += value
    }
    ```

### Switch

- 定数である必要はない

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

- 複数ある場合はコンマ区切り

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

```go
// Loopラベル設定
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])
        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                // breakする場所指定
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```

### Type Switch

- 変数の型でswitch

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("未定義: %T", t)
case bool:
    fmt.Printf("bool型: %t ", t)
case int:
    fmt.Printf("int型: %t ", t)
case *bool:
    fmt.Printf("bool型ポインタ: %t ", t)
case *int:
    fmt.Printf("int型ポインタ: %t ", t)
}
```

## 関数

### 複数の戻り値

```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}

func main() {
    for i:=0; i < len(b); {
        x, i = nextInt(b, i)
        fmt.Println(x)
    }
}
```

### Defer

- 遅延実行
- その関数の最後に実行される
- .Close()などで使用

```go
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result append(result, buf[0:n]...)
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err
        }
    }
    return string(result), nil
}
```

## Data

### new

```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}

p := new(SyncedBuffer)   // type *SyncedBuffer
ver v SyncedBuffer       // type SyncedBuffer
```

### コンストラクタ、複合リテラル

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

### make

```go
var p *[]int = new([]int)
var v []int  = make([]int, 100)
```

#### makeとnewの違い

|        |            new(T)            |         make(T)         |
| :----: | :--------------------------: | :---------------------: |
|  対象  |           任意の型           | slice, map, channelのみ |
| 初期化 | 初期化しない（ゼロ値になる） |       初期化する        |
| 戻り値 |              *T              |            T            |

### Arrays

- CとGoの違い
  - 配列は値。ある配列を別の配列に割り当てるとすべてての要素がコピーされます。
  - 配列を関数に渡すと配列へのポインタでなく、配列のコピーを受け取ります。
  - 配列のサイズはその型の一部です。`[10]int`と`[20]int`は異なります。

```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}
```

### Slice

- 配列の参照を保持
- いわゆる可変長配列

```go
var hoge []string
```

### 2次元slice

```go
type Transform [3][3]float64
type LinesOfText [][]byte

text := LinesOfText {
    []byte("New iss the time"),
    []byte("for all good gophers"),
    []byte("to bring some fun to the party.")
}
```

```go
picture := make([][]uint8, YSize)
for i := range picture {
    picture[i] = make([]uint8, XSize)
}
```

### Map

- key-valueタイプのデータ構造

```go
var timeZone = map[string]int {
    "UTC": 0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}

offset := timeZone["EST"]
```

- keyが存在しない場合は0が返される

```go
attended := map[string]bool {
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] {
    fmt.Plintln(person, "was at the meeting")
}
```

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

```go
func offset(tz, string) int {
    if seconds, of := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

- mapの存在確認

```go
_, present := timeZone[tz]
```

- mapの要素の削除

```go
delete(timeZone, "PDT")
```

### Print

```go
fmt.Printf("Hello%d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

```go
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
```

## Append

```go
x := []int{1, 2, 3}
x = append(x, 4, 5, 6)
fmt.Println(x) // [1, 2, 3, 4, 5, 6]
```

- 配列x配列

```go
x := []int{1, 2, 3}
y := []int{4, 5, 6}
x = append(x, y...)
fmt.Println(x)
```

## 初期化

### 定数

- iotaを使用することで、連番からなる定数を定義可能

```go
const (
    _ = iota    // 0
    FAST        // 1
    NORMAL      // 2
    SLOW        // 3
)
```

```go
type BeteSize float64

const (
    _ = iota
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)

```

### 変数

```go
var (
    home = os.Getenv("HOME")
    user = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```

### 関数の初期化

- init関数はmainの前に実行される

```go
func init() {

}
```

## メソッド

### ポインタvs値

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    //...
}
```

## 空白の識別子

### The blank identifier in multiple assignment

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
    fmt.Printf("%s dose not exist\n", path)
}

// Bad
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
```

### 未使用のimport

- 未使用のimportはエラー
- エラーを回避

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

### 副作用のインポート

```go
import _ "net/http/pprof"
```

## Embedding

- Goはサブクラスの概念がないが、インターフェイスないいにタイプを埋め込むことで実装の一部を借用することができる

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}

type ReadWriter struct {
    reader *Reader
    writer *Witer
}
```

## 平行性

### 平行性での共有

- メモリを共有して通信してはいけない。通信してメモリを共有する。

### Goroutines

- Goroutines
  - スレッド
  - コルーチン
  - プロセス
- メソッドの呼び出し前に`go`キーワードをつけて、新しいゴルーチンを呼び出し実行する。
- 呼び出しが完了するとゴルーチンはサイレント終了します。

```go
go list.Sort()
```

```go
package main

import (
    "fmt"
    "time"
)

func printer(text string) {
    for i := 0; 100 > i; i++ {
        fmt.Printf("%s: %d\n", text, i)
        time.Sleep(100 * time.Microsecond)
    }

    fmt.Printf("End %s\n", text)
}

func main() {
    go printer("Hello")
    printer("world")
}

```

### Channels

- ゴルーチンと通信する

```go
c := make(chan int) // channelの割当

go func() {
    list.Sort()
    c <- 1 // Send a signal

    doSomethingForAWhile()
    <-c // list.Sort()が終わるまで待つ
}
```

### 並列化

[parallelization.go](../examples/parallelization.go)
