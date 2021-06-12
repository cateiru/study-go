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

