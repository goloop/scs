[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/scs)](https://goreportcard.com/report/github.com/goloop/scs) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/scs/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/scs/v2) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# scs — String Case Style for Go

Package `scs` converts identifiers between naming conventions:
`camelCase`, `PascalCase`, `snake_case`, `kebab-case`,
`SCREAMING_SNAKE_CASE`, `dot.case` and `Title Case`.

Every conversion is built on **one universal tokenizer**. `Split` breaks any
input into normalized words; each style is then a different rendering of those
words. Because a single tokenizer feeds every renderer, the converters are
**total** — they never return an error and never need to know the input's
original style.

```go
scs.ToSnake("HTTPServerID")  // "http_server_id"
scs.ToCamel("user_id")       // "userId"
scs.ToKebab("HelloWorld")    // "hello-world"
```

## Features

- Seven case styles from a single word model.
- **Total functions:** any string maps to a well-defined result, no errors.
- Predictable, documented rules for acronyms, digits and Unicode.
- Opt-in Go-style initialisms (`ID`, `URL`, `HTTP`) via a reusable, concurrency-safe `Caser`.
- `Detect` with an honest contract: it commits to a style only when the answer is unambiguous.
- Public tokenizer: `Split` (slice) and `Words` (`iter.Seq`).
- Zero dependencies.

## Installation

```bash
go get github.com/goloop/scs/v2
```

```go
import "github.com/goloop/scs/v2"
```

## Quick start

```go
package main

import (
    "fmt"

    "github.com/goloop/scs/v2"
)

func main() {
    fmt.Println(scs.ToCamel("hello_world"))         // helloWorld
    fmt.Println(scs.ToPascal("hello-world"))        // HelloWorld
    fmt.Println(scs.ToSnake("HelloWorld"))          // hello_world
    fmt.Println(scs.ToKebab("helloWorld"))          // hello-world
    fmt.Println(scs.ToScreamingSnake("userID"))     // USER_ID
    fmt.Println(scs.ToDot("HelloWorld"))            // hello.world
    fmt.Println(scs.ToTitle("hello_world"))         // Hello World

    // Style chosen at runtime (config, CLI flag, ...).
    style, _ := scs.ParseStyle("kebab")
    fmt.Println(scs.Convert(style, "HTTPServerID")) // http-server-id
}
```

## Conversion model

The package never guesses the input's style; it always reduces the input to a
list of words and renders that list:

```go
scs.Split("HTTPServerID") // ["http", "server", "id"]
scs.Split("user_id")      // ["user", "id"]
scs.Split("web2print")    // ["web2print"]

for w := range scs.Words("parseJSONResponse") {
    fmt.Println(w) // parse, json, response
}
```

Word boundaries are placed at separators, at lower→upper transitions, at the
end of an acronym followed by a lowercase word (`HTTPServer` → `HTTP|Server`),
and before a digit-led Title word (`v2Final` → `v2|Final`).

## Initialisms (acronyms)

By default words are Title-cased (`Id`, `Url`, `Http`), which always
round-trips. To follow the Go convention of all-caps initialisms, build a
`Caser`:

```go
c := scs.New(scs.WithAcronyms("ID", "URL", "HTTP", "API"))

c.ToPascal("user_id")        // "UserID"
c.ToCamel("http_url_builder") // "httpURLBuilder" (first word stays lowercase)
```

A `Caser` is immutable and safe for concurrent use. Initialisms are opt-in
because adjacent all-caps acronyms cannot always be split apart again
(`"HTTPAPI"` is ambiguous), so the round-trip-safe Title casing is the default.

## Numbers

Digits attach to neighboring letters and never split a word on their own, so
identifier fragments stay intact:

```go
scs.ToSnake("web2print") // "web2print"
scs.ToSnake("sha256sum") // "sha256sum"
scs.ToSnake("oauth2")    // "oauth2"
```

Only an explicit separator turns a number into its own word:

```go
scs.ToSnake("web 2 print") // "web_2_print"
```

## Detection

`Detect` returns a style only when the input is canonical for exactly one
style. Ambiguous inputs (such as the bare word `"api"`, valid in several
styles at once) return `(Unknown, false)`:

```go
scs.Detect("user_id") // Snake, true
scs.Detect("userId")  // Camel, true
scs.Detect("USER_ID") // ScreamingSnake, true
scs.Detect("api")     // Unknown, false
```

The per-style predicates `IsCamel`, `IsSnake`, … (and the generic
`Is(style, s)`) report whether a string is already canonical for that style.

## Guarantees and limits

- The converters never panic and always return valid UTF-8.
- `snake_case`, `kebab-case` and `dot.case` are **idempotent** and preserve the
  exact word sequence for any input, so routing a value through them is lossless.
- For ordinary multi-letter identifiers every style round-trips
  (`snake → camel → snake` returns the original).
- The cased joined styles (`camelCase`, `PascalCase`, `Title Case`) cannot be
  idempotent for *single-letter* or adjacent all-caps words, because such
  renderings are inherently ambiguous with acronyms — this is a property of any
  dictionary-free converter, not a defect.
- A dictionary-free tokenizer cannot tell `IPv6` apart from the
  "acronym + lowercase word" pattern, so `ToSnake("IPv6Address")` yields
  `i_pv6_address`. Use `WithAcronyms` or a separator when this matters.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE)
file for details.
