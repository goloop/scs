// Package scs (String Case Style) converts identifiers between naming
// conventions: camelCase, PascalCase, snake_case, kebab-case,
// SCREAMING_SNAKE_CASE, dot.case and Title Case.
//
// # Model
//
// Every conversion is built on one universal tokenizer. Split breaks any input
// into normalized words by honoring separators, case transitions and acronym
// boundaries; each style is then a different rendering of those words:
//
//	scs.ToSnake("HTTPServerID")  // "http_server_id"
//	scs.ToCamel("user_id")       // "userId"
//	scs.ToKebab("HelloWorld")    // "hello-world"
//
// Because a single tokenizer feeds every renderer, the converters are total:
// they never return an error and never need to know the input's original
// style. Any string maps to a well-defined result.
//
// # Initialisms
//
// By default words are Title-cased ("Id", "Url", "Http"), which always
// round-trips. To follow the Go convention of all-caps initialisms, build a
// Caser with preserved acronyms:
//
//	c := scs.New(scs.WithAcronyms("ID", "URL", "HTTP"))
//	c.ToPascal("user_id") // "UserID"
//
// # Numbers
//
// Digits attach to neighboring letters and never split a word on their own, so
// identifier fragments stay intact ("utf8", "sha256", "oauth2"). Only an
// explicit separator turns a number into its own word ("web 2 print" ->
// "web_2_print").
//
// # Thread safety
//
// All package-level functions and all *Caser methods are safe for concurrent
// use; a Caser is immutable once built.
package scs

import "unicode"

// ToCamel converts s to camelCase using the default configuration.
//
//	scs.ToCamel("hello_world") // "helloWorld"
//	scs.ToCamel("HelloWorld")  // "helloWorld"
//	scs.ToCamel("HTTP server")  // "httpServer"
func ToCamel(s string) string { return defaultCaser.ToCamel(s) }

// ToPascal converts s to PascalCase using the default configuration.
//
//	scs.ToPascal("hello_world") // "HelloWorld"
//	scs.ToPascal("helloWorld")  // "HelloWorld"
func ToPascal(s string) string { return defaultCaser.ToPascal(s) }

// ToSnake converts s to snake_case.
//
//	scs.ToSnake("helloWorld") // "hello_world"
//	scs.ToSnake("HTTPServer")  // "http_server"
func ToSnake(s string) string { return defaultCaser.ToSnake(s) }

// ToKebab converts s to kebab-case.
//
//	scs.ToKebab("helloWorld") // "hello-world"
//	scs.ToKebab("HelloWorld") // "hello-world"
func ToKebab(s string) string { return defaultCaser.ToKebab(s) }

// ToScreamingSnake converts s to SCREAMING_SNAKE_CASE, the usual style for
// constants and environment variables.
//
//	scs.ToScreamingSnake("helloWorld") // "HELLO_WORLD"
//	scs.ToScreamingSnake("userID")     // "USER_ID"
func ToScreamingSnake(s string) string { return defaultCaser.ToScreamingSnake(s) }

// ToDot converts s to dot.case.
//
//	scs.ToDot("helloWorld") // "hello.world"
func ToDot(s string) string { return defaultCaser.ToDot(s) }

// ToTitle converts s to Title Case.
//
//	scs.ToTitle("hello_world") // "Hello World"
func ToTitle(s string) string { return defaultCaser.ToTitle(s) }

// ToSentence converts s to Sentence case (only the first word capitalized).
//
//	scs.ToSentence("hello_world") // "Hello world"
func ToSentence(s string) string { return defaultCaser.ToSentence(s) }

// Convert renders s in the requested style using the default configuration.
// An Unknown or out-of-range style returns s unchanged, so Convert is total
// and never panics.
//
//	scs.Convert(scs.Kebab, "userID") // "user-id"
func Convert(to Style, s string) string { return defaultCaser.Convert(to, s) }

// IsCamel reports whether s is already in canonical camelCase, i.e. it is
// non-empty and ToCamel leaves it unchanged.
//
// Note that a single lowercase word such as "api" is canonical in several
// styles at once (camel, snake, kebab, dot); use Detect when you need a single
// unambiguous answer.
func IsCamel(s string) bool { return s != "" && ToCamel(s) == s }

// IsPascal reports whether s is already in canonical PascalCase.
func IsPascal(s string) bool { return s != "" && ToPascal(s) == s }

// IsSnake reports whether s is already in canonical snake_case.
func IsSnake(s string) bool { return s != "" && ToSnake(s) == s }

// IsKebab reports whether s is already in canonical kebab-case.
func IsKebab(s string) bool { return s != "" && ToKebab(s) == s }

// IsScreamingSnake reports whether s is already in canonical
// SCREAMING_SNAKE_CASE.
func IsScreamingSnake(s string) bool { return s != "" && ToScreamingSnake(s) == s }

// IsDot reports whether s is already in canonical dot.case.
func IsDot(s string) bool { return s != "" && ToDot(s) == s }

// IsTitle reports whether s is already in canonical Title Case.
func IsTitle(s string) bool { return s != "" && ToTitle(s) == s }

// IsSentence reports whether s is already in canonical Sentence case.
func IsSentence(s string) bool { return s != "" && ToSentence(s) == s }

// Is reports whether s is already canonical for the given style. An invalid
// style always reports false.
func Is(style Style, s string) bool {
	if s == "" || !style.Valid() {
		return false
	}
	return Convert(style, s) == s
}

// Separator bits recorded by scanSeps. A canonical value uses exactly one of
// these as its separator, so the set present in the input prunes the styles
// worth probing.
const (
	sepUnderscore uint8 = 1 << iota
	sepHyphen
	sepDot
	sepSpace
)

// Candidate style lists for Detect, keyed by the separator found in the input.
// These are package vars so the hot path allocates nothing.
var (
	// noSepUpperStyles: a separatorless token that contains an uppercase letter
	// can only be canonical for a cased or all-caps style; the lowercase joined
	// styles (snake/kebab/dot) would lowercase it and never match.
	noSepUpperStyles = []Style{Camel, Pascal, ScreamingSnake, Title, Sentence}
	underscoreStyles = []Style{Snake, ScreamingSnake}
	hyphenStyles     = []Style{Kebab}
	dotStyles        = []Style{Dot}
	spaceStyles      = []Style{Title, Sentence}
)

// scanSeps walks s once and reports which canonical separators ('_', '-', '.',
// ' ') appear, whether any uppercase letter is present, and whether any rune is
// neither a letter, a digit, nor one of those separators (hasOther) — such a
// rune cannot occur in any canonical value, so the input is non-canonical.
//
// "Uppercase" here means an uppercase letter with a distinct lowercase form,
// matching the tokenizer's classOf, so the result stays consistent with how the
// converters actually behave.
func scanSeps(s string) (mask uint8, hasUpper, hasOther bool) {
	for _, r := range s {
		switch r {
		case '_':
			mask |= sepUnderscore
		case '-':
			mask |= sepHyphen
		case '.':
			mask |= sepDot
		case ' ':
			mask |= sepSpace
		default:
			if unicode.IsUpper(r) && unicode.ToLower(r) != r {
				hasUpper = true
			} else if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				return mask, hasUpper, true
			}
		}
	}
	return mask, hasUpper, false
}

// Detect returns the single style for which s is already canonical. If s is
// empty, matches no style, or matches more than one style (for example the
// bare word "api", which is valid camel, snake, kebab and dot at once), it
// returns Unknown and false.
//
//	scs.Detect("user_id")  // Snake, true
//	scs.Detect("userId")   // Camel, true
//	scs.Detect("USER_ID")  // ScreamingSnake, true
//	scs.Detect("api")      // Unknown, false (ambiguous)
//	scs.Detect("Hello_World") // Unknown, false (no canonical style)
//
// The input is scanned once for its separator so that only the styles which
// could possibly match are probed; a string mixing two separators, or carrying
// any other punctuation, is rejected without conversion.
func Detect(s string) (Style, bool) {
	if s == "" {
		return Unknown, false
	}

	mask, hasUpper, hasOther := scanSeps(s)
	if hasOther {
		return Unknown, false
	}

	var candidates []Style
	switch mask {
	case 0:
		// A separatorless, all-lowercase token equals its own snake, kebab,
		// dot and camel renderings at once, so it is always ambiguous; only a
		// token carrying an uppercase letter can pin a single cased style.
		if !hasUpper {
			return Unknown, false
		}
		candidates = noSepUpperStyles
	case sepUnderscore:
		candidates = underscoreStyles
	case sepHyphen:
		candidates = hyphenStyles
	case sepDot:
		candidates = dotStyles
	case sepSpace:
		candidates = spaceStyles
	default:
		// More than one separator type: no canonical value mixes them.
		return Unknown, false
	}

	found := Unknown
	for _, style := range candidates {
		if Convert(style, s) == s {
			if found != Unknown {
				return Unknown, false // ambiguous: more than one match
			}
			found = style
		}
	}

	if found == Unknown {
		return Unknown, false
	}
	return found, true
}
