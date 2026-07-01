[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/scs/v2)](https://goreportcard.com/report/github.com/goloop/scs/v2) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/scs/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/scs/v2) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# scs — String Case Style for Go

`scs` converts identifiers between naming conventions: `camelCase`,
`PascalCase`, `snake_case`, `kebab-case`, `SCREAMING_SNAKE_CASE`, `dot.case`,
`Title Case` and `Sentence case`.

Every conversion is built on **one universal tokenizer**. `Split` breaks any
input into normalized words; each style is then a different rendering of those
words. Because a single tokenizer feeds every renderer, the converters are
**total** — they never return an error and never need to know the input's
original style.

```go
scs.ToSnake("HTTPServerID") // "http_server_id"
scs.ToCamel("user_id")      // "userId"
scs.ToKebab("HelloWorld")   // "hello-world"
```

## Features

- Eight case styles from a single word model.
- **Total functions** — any string maps to a well-defined result, no errors.
- Predictable, documented rules for acronyms, digits and Unicode.
- Opt-in Go-style initialisms (`ID`, `URL`, `HTTP`) via a reusable,
  concurrency-safe `Caser`.
- `Detect` with an honest contract — it commits to a style only when the answer
  is unambiguous.
- Public tokenizer: `Split` (slice) and `Words` (`iter.Seq`).
- Zero dependencies.

## Installation

```bash
go get github.com/goloop/scs/v2
```

```go
import "github.com/goloop/scs/v2"
```

Requires Go 1.24 or newer. The package has no third-party dependencies.

## Quick start

```go
package main

import (
    "fmt"

    "github.com/goloop/scs/v2"
)

func main() {
    fmt.Println(scs.ToCamel("hello_world"))     // helloWorld
    fmt.Println(scs.ToPascal("hello-world"))    // HelloWorld
    fmt.Println(scs.ToSnake("HelloWorld"))      // hello_world
    fmt.Println(scs.ToKebab("helloWorld"))      // hello-world
    fmt.Println(scs.ToScreamingSnake("userID")) // USER_ID
    fmt.Println(scs.ToSentence("hello_world"))  // Hello world

    // Style chosen at runtime (config, CLI flag, ...).
    style, _ := scs.ParseStyle("kebab")
    fmt.Println(scs.Convert(style, "HTTPServerID")) // http-server-id

    // Go-style all-caps initialisms are opt-in via a reusable Caser.
    c := scs.New(scs.WithAcronyms("ID", "URL", "HTTP"))
    fmt.Println(c.ToPascal("user_id")) // UserID
}
```

## Documentation

- Full reference and recipes: [DOC.md](DOC.md) · [DOC.UK.md](DOC.UK.md)
- Package API: [pkg.go.dev/github.com/goloop/scs/v2](https://pkg.go.dev/github.com/goloop/scs/v2)
- Changes between versions: [CHANGELOG.md](CHANGELOG.md)

## Contributing

Contributions are welcome. Please run `go test ./...`, `go vet ./...` and
`gofmt -l .` before submitting a pull request.

## License

`scs` is released under the MIT License. See [LICENSE](LICENSE).
