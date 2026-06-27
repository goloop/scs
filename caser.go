package scs

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Caser is a reusable, immutable converter configuration. The package-level
// functions (ToCamel, ToSnake, ...) use a zero-configuration Caser that always
// renders Title-cased words; build a custom one with New when you need extra
// behavior such as preserved initialisms.
//
// A Caser is safe for concurrent use by multiple goroutines: it is never
// mutated after New returns.
type Caser struct {
	// acronyms maps a lowercased word to its preferred all-caps rendering.
	// nil means "no initialisms", the canonical default.
	acronyms map[string]string
}

// Option configures a Caser. Options are applied by New in order.
type Option func(*Caser)

// WithAcronyms makes the Caser render the given words as all-caps initialisms
// in the camelCase, PascalCase and Title styles, matching the Go convention of
// writing "ID", "URL" or "HTTP" instead of "Id", "Url" or "Http". Matching is
// case-insensitive on word boundaries detected by the tokenizer.
//
//	c := scs.New(scs.WithAcronyms("ID", "URL", "HTTP"))
//	c.ToPascal("user_id_url") // "UserIDURL"
//	c.ToCamel("http_server")  // "httpServer"  (first word stays lowercase)
//
// Initialisms are an opt-in convenience, not the default, because adjacent
// all-caps acronyms cannot always be split back apart: "HTTPAPI" is ambiguous
// between [HTTP, API] and other partitions. The zero-configuration Caser keeps
// Title casing, which always round-trips.
func WithAcronyms(words ...string) Option {
	return func(c *Caser) {
		for _, w := range words {
			key := strings.ToLower(strings.TrimSpace(w))
			if key == "" {
				continue
			}
			if c.acronyms == nil { // allocate only once a real key appears
				c.acronyms = make(map[string]string, len(words))
			}
			c.acronyms[key] = strings.ToUpper(key)
		}
	}
}

// New returns a Caser configured by the given options. With no options it is
// equivalent to the package-level functions.
func New(opts ...Option) *Caser {
	c := &Caser{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// defaultCaser backs the package-level functions. It has no initialisms.
var defaultCaser = &Caser{}

// cap1 capitalizes the first rune of an already-lowercased word, honoring the
// Caser's initialism set. ASCII words take an allocation-free fast path.
func (c *Caser) cap1(w string) string {
	if w == "" {
		return w
	}
	if c.acronyms != nil {
		if up, ok := c.acronyms[w]; ok {
			return up
		}
	}

	r0 := w[0]
	if r0 < utf8.RuneSelf {
		if r0 >= 'a' && r0 <= 'z' {
			return string(r0-('a'-'A')) + w[1:]
		}
		return w
	}

	r, size := utf8.DecodeRuneInString(w)
	return string(unicode.ToUpper(r)) + w[size:]
}

// ToCamel renders s as camelCase.
func (c *Caser) ToCamel(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	first := true
	for w := range Words(s) {
		if first {
			b.WriteString(w) // first word stays lowercase, even if an acronym
			first = false
			continue
		}
		b.WriteString(c.cap1(w))
	}
	return b.String()
}

// ToPascal renders s as PascalCase.
func (c *Caser) ToPascal(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for w := range Words(s) {
		b.WriteString(c.cap1(w))
	}
	return b.String()
}

// ToTitle renders s as Title Case (capitalized words joined by spaces).
func (c *Caser) ToTitle(s string) string {
	var b strings.Builder
	b.Grow(len(s) + 4)
	first := true
	for w := range Words(s) {
		if !first {
			b.WriteByte(' ')
		}
		b.WriteString(c.cap1(w))
		first = false
	}
	return b.String()
}

// ToSnake renders s as snake_case.
func (c *Caser) ToSnake(s string) string { return joinLower(s, '_') }

// ToKebab renders s as kebab-case.
func (c *Caser) ToKebab(s string) string { return joinLower(s, '-') }

// ToDot renders s as dot.case.
func (c *Caser) ToDot(s string) string { return joinLower(s, '.') }

// ToScreamingSnake renders s as SCREAMING_SNAKE_CASE.
func (c *Caser) ToScreamingSnake(s string) string {
	return strings.ToUpper(joinLower(s, '_'))
}

// Convert renders s in the requested style. An Unknown or out-of-range style
// returns s unchanged, so Convert is always total.
func (c *Caser) Convert(to Style, s string) string {
	switch to {
	case Camel:
		return c.ToCamel(s)
	case Pascal:
		return c.ToPascal(s)
	case Snake:
		return c.ToSnake(s)
	case Kebab:
		return c.ToKebab(s)
	case ScreamingSnake:
		return c.ToScreamingSnake(s)
	case Dot:
		return c.ToDot(s)
	case Title:
		return c.ToTitle(s)
	default:
		return s
	}
}

// joinLower joins the words of s with the delimiter. The lowercase separated
// styles do not depend on the acronym set, so this is a free function shared
// by every Caser.
func joinLower(s string, delim byte) string {
	var b strings.Builder
	b.Grow(len(s))
	first := true
	for w := range Words(s) {
		if !first {
			b.WriteByte(delim)
		}
		b.WriteString(w)
		first = false
	}
	return b.String()
}
