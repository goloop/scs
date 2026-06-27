package scs

import (
	"iter"
	"unicode"
	"unicode/utf8"
)

// class categorizes a rune for boundary detection.
type class uint8

const (
	clOther class = iota // separator: anything that is not a letter or digit
	clLower              // lowercase (or caseless, e.g. Han) letter
	clUpper              // uppercase letter
	clDigit              // decimal digit
)

// classOf maps a rune to its boundary class. A rune counts as uppercase only
// when it has a distinct lowercase form; uppercase letterlikes that do not
// lowercase (e.g. U+210C ℌ) and caseless scripts (Han, Hiragana, Arabic) are
// treated as lowercase, so they glue into words and survive normalization
// without spawning phantom boundaries on re-parse.
func classOf(r rune) class {
	if r < utf8.RuneSelf { // ASCII fast path
		switch {
		case r >= 'a' && r <= 'z':
			return clLower
		case r >= 'A' && r <= 'Z':
			return clUpper
		case r >= '0' && r <= '9':
			return clDigit
		default:
			return clOther
		}
	}

	switch {
	case unicode.IsDigit(r):
		return clDigit
	case unicode.IsUpper(r) && unicode.ToLower(r) != r:
		return clUpper
	case unicode.IsLetter(r):
		return clLower
	default:
		return clOther
	}
}

// splitSeq walks s once and yields normalized (lowercased) word tokens.
// It is the single source of truth for word boundaries in the package; every
// converter is just a different rendering of these words.
//
// Boundaries are placed:
//   - at any run of separators (non-letter, non-digit), which is dropped;
//   - between a lowercase/caseless letter and an uppercase letter
//     ("helloWorld" -> hello | World);
//   - at the end of an uppercase acronym that is followed by a lowercase
//     letter ("HTTPServer" -> HTTP | Server);
//   - between a digit and an uppercase letter that begins a Title-cased word
//     ("v2Final" -> v2 | Final).
//
// Digits never break away from adjacent letters on their own, so identifier
// fragments such as "utf8", "sha256" and "oauth2" stay intact; only an
// explicit separator splits a number into its own word ("web 2 print").
//
// A digit followed by an all-uppercase run does not split ("version2ID" stays
// one word), because an all-caps spelling is ambiguous: it is exactly what
// SCREAMING_SNAKE_CASE produces, so honoring such a break would make that style
// non-idempotent. This keeps the lowercase separated styles (snake, kebab, dot)
// and screaming-snake stable under re-parsing.
//
// Yielding stops early if yield returns false.
func splitSeq(s string, yield func(string) bool) {
	// start is the byte offset of the current word, or -1 when we are
	// between words (inside a separator run). prev is the class of the
	// previous rune; prevUpper marks whether that rune was uppercase, which
	// the acronym rule needs as one-rune lookbehind.
	start := -1
	var prev class

	for i, r := range s {
		cur := classOf(r)

		if cur == clOther {
			if start >= 0 {
				if !yield(normalize(s[start:i])) {
					return
				}
				start = -1
			}
			prev = clOther
			continue
		}

		if start < 0 {
			// First rune of a new word.
			start = i
			prev = cur
			continue
		}

		if boundaryBetween(prev, cur, s, i) {
			if !yield(normalize(s[start:i])) {
				return
			}
			start = i
		}
		prev = cur
	}

	if start >= 0 {
		yield(normalize(s[start:]))
	}
}

// boundaryBetween reports whether a word boundary falls before the rune at
// byte offset i (whose class is cur), given that the previous rune had class
// prev. s and i are used only for the acronym look-ahead.
func boundaryBetween(prev, cur class, s string, i int) bool {
	switch {
	case prev == clLower && cur == clUpper:
		// lower -> Upper: "helloWorld".
		return true
	case prev == clDigit && cur == clUpper:
		// digit -> Upper: only when the uppercase starts a Title-cased word
		// ("v2Final" -> v2 | Final). A digit before an all-caps run does not
		// split, so SCREAMING_SNAKE output re-parses to the same words.
		return startsLowerWord(s, i)
	case prev == clUpper && cur == clUpper:
		// Acronym run: break before the last uppercase letter when it is
		// followed by a lowercase letter, so "HTTPServer" -> HTTP | Server.
		return startsLowerWord(s, i)
	default:
		return false
	}
}

// startsLowerWord reports whether the uppercase rune at byte offset i is
// immediately followed by a lowercase/caseless letter, i.e. it begins a new
// capitalized word rather than continuing an acronym.
func startsLowerWord(s string, i int) bool {
	_, size := utf8.DecodeRuneInString(s[i:])
	next := i + size
	if next >= len(s) {
		return false
	}
	r, _ := utf8.DecodeRuneInString(s[next:])
	return classOf(r) == clLower
}

// normalize lowercases a word token. It fast-paths pure ASCII (the common
// case for identifiers) to avoid allocating when nothing changes.
func normalize(w string) string {
	ascii := true
	hasUpper := false
	for i := 0; i < len(w); i++ {
		c := w[i]
		if c >= utf8.RuneSelf {
			ascii = false
			break
		}
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		}
	}

	if ascii {
		if !hasUpper {
			return w
		}
		b := make([]byte, len(w))
		for i := 0; i < len(w); i++ {
			c := w[i]
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
			}
			b[i] = c
		}
		return string(b)
	}

	return toLowerUnicode(w)
}

// toLowerUnicode lowercases a string that contains non-ASCII runes.
func toLowerUnicode(w string) string {
	out := make([]rune, 0, len(w))
	for _, r := range w {
		out = append(out, unicode.ToLower(r))
	}
	return string(out)
}

// Split returns the normalized (lowercased) words of s, as detected by the
// package's universal tokenizer. It is the building block every converter
// shares: ToSnake(s) is the words joined by "_", ToCamel(s) is the words
// rendered camelCase, and so on.
//
// The result is empty for an empty string or a string of separators only.
//
//	scs.Split("HTTPServerID")  // ["http", "server", "id"]
//	scs.Split("user_id")       // ["user", "id"]
//	scs.Split("web2print")     // ["web2print"]
func Split(s string) []string {
	// Pre-size for the typical case: one word per separator run plus one.
	words := make([]string, 0, 4)
	splitSeq(s, func(w string) bool {
		words = append(words, w)
		return true
	})
	return words
}

// Words returns an iterator over the normalized words of s. It yields exactly
// the same tokens as Split without materializing a slice, which lets callers
// build custom renderings cheaply.
//
//	for w := range scs.Words("HTTPServerID") {
//		fmt.Println(w) // http, server, id
//	}
func Words(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		splitSeq(s, yield)
	}
}
