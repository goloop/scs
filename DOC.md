# scs — reference

The full reference for the `scs` package: the mental model, the tokenizer, every
style and converter, the `Caser` and initialisms, detection, the guarantees and
limits, and practical recipes.

Ukrainian version: **[DOC.UK.md](DOC.UK.md)**.

## Contents

- [Mental model](#mental-model)
- [The tokenizer: Split and Words](#the-tokenizer-split-and-words)
- [Styles and converters](#styles-and-converters)
- [Choosing a style at runtime](#choosing-a-style-at-runtime)
- [Initialisms and the Caser](#initialisms-and-the-caser)
- [Numbers](#numbers)
- [Detection and predicates](#detection-and-predicates)
- [Guarantees and limits](#guarantees-and-limits)
- [Recipes and tips](#recipes-and-tips)

## Mental model

`scs` converts identifiers between naming conventions: `camelCase`,
`PascalCase`, `snake_case`, `kebab-case`, `SCREAMING_SNAKE_CASE`, `dot.case`,
`Title Case` and `Sentence case`.

The whole package rests on one idea: **there is a single universal tokenizer.**
`Split` reduces any input to a list of normalized words, and each style is just
a different rendering of that list. Because one tokenizer feeds every renderer:

- The converters are **total** — they never return an error and never need to
  know the input's original style.
- Adding a style means adding a renderer, not another N×N pairwise converter.

```go
scs.ToSnake("HTTPServerID") // "http_server_id"
scs.ToCamel("user_id")      // "userId"
scs.ToKebab("HelloWorld")   // "hello-world"
```

```go
import "github.com/goloop/scs/v2"
```

## The tokenizer: Split and Words

```go
func Split(s string) []string
func Words(s string) iter.Seq[string]
```

`Split` returns the normalized words; `Words` yields them lazily as an
`iter.Seq[string]`:

```go
scs.Split("HTTPServerID")   // ["http", "server", "id"]
scs.Split("user_id")        // ["user", "id"]
scs.Split("web2print")      // ["web2print"]

for w := range scs.Words("parseJSONResponse") {
    fmt.Println(w) // parse, json, response
}
```

Word boundaries are placed:

- at separators (`_`, `-`, `.`, space);
- at lower→upper transitions (`helloWorld` → `hello|World`);
- at the end of an acronym followed by a lowercase word (`HTTPServer` →
  `HTTP|Server`);
- before a digit-led Title word (`v2Final` → `v2|Final`).

## Styles and converters

Each style has a total `To…` function:

| Function | Example output |
|----------|----------------|
| `ToCamel`          | `helloWorld` |
| `ToPascal`         | `HelloWorld` |
| `ToSnake`          | `hello_world` |
| `ToKebab`          | `hello-world` |
| `ToScreamingSnake` | `HELLO_WORLD` |
| `ToDot`            | `hello.world` |
| `ToTitle`          | `Hello World` |
| `ToSentence`       | `Hello world` |

```go
scs.ToCamel("hello_world")      // helloWorld
scs.ToPascal("hello-world")     // HelloWorld
scs.ToScreamingSnake("userID")  // USER_ID
scs.ToSentence("hello_world")   // Hello world
```

## Choosing a style at runtime

When the target style comes from config, a CLI flag or a database, use the
`Style` enum and `Convert`:

```go
func Convert(to Style, s string) string
func ParseStyle(name string) (Style, bool)
func (s Style) String() string
func (s Style) Valid() bool
```

```go
style, ok := scs.ParseStyle("kebab")     // Kebab, true
out := scs.Convert(style, "HTTPServerID") // "http-server-id"
```

The `Style` constants are `Unknown`, `Camel`, `Pascal`, `Snake`, `Kebab`,
`ScreamingSnake`, `Dot`, `Title` and `Sentence`. `Unknown` is the zero value and
the "no single style" answer from `Detect`.

## Initialisms and the Caser

By default words are Title-cased (`Id`, `Url`, `Http`), which always
round-trips. To follow the Go convention of all-caps initialisms, build a
`Caser`:

```go
func New(opts ...Option) *Caser
func WithAcronyms(words ...string) Option
```

```go
c := scs.New(scs.WithAcronyms("ID", "URL", "HTTP", "API"))

c.ToPascal("user_id")         // "UserID"
c.ToCamel("http_url_builder") // "httpURLBuilder" (first word stays lowercase)
```

A `Caser` exposes the same `To…` and `Convert` methods as the package functions.
It is **immutable and safe for concurrent use** — build it once and share it.

Initialisms are opt-in because adjacent all-caps acronyms cannot always be split
apart again (`"HTTPAPI"` is ambiguous), so the round-trip-safe Title casing is
the default.

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

## Detection and predicates

```go
func Detect(s string) (Style, bool)
func Is(style Style, s string) bool
func IsCamel(s string) bool // IsSnake, IsKebab, IsPascal, IsScreamingSnake,
                            // IsDot, IsTitle, IsSentence
```

`Detect` returns a style **only when the input is canonical for exactly one
style**. Ambiguous inputs (such as the bare word `"api"`, valid in several
styles at once) return `(Unknown, false)`:

```go
scs.Detect("user_id") // Snake, true
scs.Detect("userId")  // Camel, true
scs.Detect("USER_ID") // ScreamingSnake, true
scs.Detect("api")     // Unknown, false
```

The per-style predicates (and the generic `Is(style, s)`) report whether a
string is already canonical for that style — useful to skip a no-op conversion.

## Guarantees and limits

- The converters **never panic** and always return valid UTF-8.
- `snake_case`, `kebab-case` and `dot.case` are **idempotent** and preserve the
  exact word sequence for any input, so routing a value through them is lossless.
- For ordinary multi-letter identifiers every style round-trips
  (`snake → camel → snake` returns the original).
- The cased joined styles (`camelCase`, `PascalCase`, `Title Case`) cannot be
  idempotent for *single-letter* or adjacent all-caps words, because such
  renderings are inherently ambiguous with acronyms — a property of any
  dictionary-free converter, not a defect.
- A dictionary-free tokenizer cannot tell `IPv6` apart from the "acronym +
  lowercase word" pattern, so `ToSnake("IPv6Address")` yields `i_pv6_address`.
  Use `WithAcronyms` or a separator when this matters.

## Recipes and tips

**Normalise to a lossless style first.** When you need a stable key, convert to
`snake_case`/`kebab-case`/`dot.case` — they are idempotent and preserve the word
sequence, so the value survives repeated processing.

**Share one Caser.** Build `scs.New(scs.WithAcronyms(...))` once at startup and
reuse it everywhere; it is immutable and concurrency-safe.

**Drive the style from data.** Store or receive a style name, resolve it with
`ParseStyle`, and render with `Convert` — no switch statement over eight
functions.

**Skip no-op conversions.** Guard with the `Is…` predicate (or `Is(style, s)`)
when you only want to convert values that are not already canonical.

**Disambiguate mixed-case tech terms.** For tokens like `IPv6` or `OAuth2`,
supply a separator in the source or register the acronym with `WithAcronyms`.
