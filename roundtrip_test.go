package scs

import (
	"slices"
	"strings"
	"testing"
)

// corpus is a broad set of inputs spanning every style, mixed delimiters,
// acronyms, digits, Unicode and degenerate cases. The property tests below run
// every invariant against all of them.
var corpus = []string{
	"", "   ", "_", "-", ".", "a", "A", "API", "id",
	"helloWorld", "HelloWorld", "hello_world", "hello-world", "hello.world",
	"HELLO_WORLD", "Hello World", "hello world",
	"HTTPServer", "HTTPServerID", "userID", "OAuthClient", "IOReader",
	"HTTP API", "parse XMLHTTPRequest now",
	"web2print", "version2Final", "utf8", "sha256", "oauth2", "web 2 print",
	"  leading", "trailing  ", "__dunder__", "mixed_Case-style.here",
	"ПриватБанк", "привіт_світ", "straßeTest", "日本語Text",
	"aB", "Ab", "ABc", "aBc", "ID2", "v2", "X",
}

var allStyles = []Style{Camel, Pascal, Snake, Kebab, ScreamingSnake, Dot, Title}

// stableStyles are the styles guaranteed idempotent and word-preserving for
// *every* input: they lowercase and join with a separator, so their output
// re-parses to exactly the same words. The cased joined styles (camel, pascal,
// title) and the all-caps style cannot offer this universally because adjacent
// single capitals are inherently ambiguous with acronyms — see
// TestCasedStylesAmbiguityIsDocumented.
var stableStyles = []Style{Snake, Kebab, Dot}

// wellBehaved holds typical multi-letter identifiers, free of the inherently
// ambiguous shapes (single-letter words, lone caseless runes inside all-caps).
// Every style is idempotent and round-trips on these.
var wellBehaved = []string{
	"userName", "HTTPServer", "parse_xml_document", "OAuthClient",
	"version2Final", "привітСвіт", "some-kebab-value", "ALREADY_CONST",
	"Title Cased Words", "httpServerId", "load.config.file",
}

// TestStableStylesAlwaysIdempotent: snake, kebab and dot are fixed points of
// themselves for any input whatsoever.
func TestStableStylesAlwaysIdempotent(t *testing.T) {
	for _, s := range corpus {
		for _, style := range stableStyles {
			once := Convert(style, s)
			if twice := Convert(style, once); once != twice {
				t.Errorf("Convert(%v) not idempotent on %q: once=%q twice=%q",
					style, s, once, twice)
			}
		}
	}
}

// TestAllStylesIdempotentOnWellBehaved: on ordinary identifiers every style,
// including the cased and all-caps ones, is idempotent.
func TestAllStylesIdempotentOnWellBehaved(t *testing.T) {
	for _, s := range wellBehaved {
		for _, style := range allStyles {
			once := Convert(style, s)
			if twice := Convert(style, once); once != twice {
				t.Errorf("Convert(%v) not idempotent on %q: once=%q twice=%q",
					style, s, once, twice)
			}
		}
	}
}

// TestIsMatchesConverter: a non-empty converted value is, by definition,
// canonical for its style — Is must confirm it. Checked on the styles and
// inputs where canonicity is guaranteed.
func TestIsMatchesConverter(t *testing.T) {
	for _, s := range corpus {
		for _, style := range stableStyles {
			if out := Convert(style, s); out != "" && !Is(style, out) {
				t.Errorf("Is(%v, %q) is false for output of Convert(%v, %q)",
					style, out, style, s)
			}
		}
	}
	for _, s := range wellBehaved {
		for _, style := range allStyles {
			if out := Convert(style, s); out != "" && !Is(style, out) {
				t.Errorf("Is(%v, %q) is false for output of Convert(%v, %q)",
					style, out, style, s)
			}
		}
	}
}

// TestStableStylesPreserveWords: snake, kebab and dot carry exactly the
// tokenizer's words, so re-splitting their output reproduces the original word
// list. This is the structural guarantee behind round-tripping.
func TestStableStylesPreserveWords(t *testing.T) {
	for _, s := range corpus {
		want := Split(s)
		for _, style := range stableStyles {
			got := Split(Convert(style, s))
			if !slices.Equal(got, want) {
				t.Errorf("Split(Convert(%v, %q)) = %#v, want words %#v",
					style, s, got, want)
			}
		}
	}
}

// TestCasedStylesAmbiguityIsDocumented pins the inherent limit: rendering
// single-letter words in a cased joined style produces adjacent capitals that
// are indistinguishable from an acronym, so the cased styles are not idempotent
// there. This is true of any dictionary-free converter; the test exists so the
// behavior is observed and intentional.
func TestCasedStylesAmbiguityIsDocumented(t *testing.T) {
	if once, twice := ToPascal("aB"), ToPascal(ToPascal("aB")); once == twice {
		t.Errorf("expected cased ambiguity on %q, but Pascal was idempotent (%q)", "aB", once)
	}
}

// TestSnakeIsACanonicalCarrier: because snake_case preserves every word
// boundary, converting through it must not change any rendering. This proves
// the converters depend only on the word sequence, not on the input's spelling.
func TestSnakeIsACanonicalCarrier(t *testing.T) {
	for _, s := range corpus {
		snake := ToSnake(s)
		for _, style := range allStyles {
			direct := Convert(style, s)
			viaSnake := Convert(style, snake)
			if direct != viaSnake {
				t.Errorf("Convert(%v, %q)=%q but Convert(%v, ToSnake)=%q",
					style, s, direct, style, viaSnake)
			}
		}
	}
}

// TestLetterWordsBijection: for identifiers made of letter-only words, the
// separated and joined styles are fully reversible — convert away and back and
// the canonical snake form is recovered. (Digit-bearing words are excluded
// here because the glued-digit policy deliberately drops standalone-number
// boundaries; that trade-off is pinned in TestDigitWordsAreLossyUnderCamel.)
func TestLetterWordsBijection(t *testing.T) {
	wordSets := [][]string{
		{"user", "name"},
		{"http", "server", "id"},
		{"one", "two", "three", "four"},
		{"привіт", "світ"},
	}
	for _, words := range wordSets {
		snake := strings.Join(words, "_")
		for _, style := range allStyles {
			there := Convert(style, snake)
			back := ToSnake(there)
			if back != snake {
				t.Errorf("round-trip via %v failed: %q -> %q -> %q",
					style, snake, there, back)
			}
		}
	}
}

// TestDigitWordsAreLossyUnderCamel pins the deliberate consequence of the
// glued-digit policy: a standalone number separated by spaces survives snake
// but merges under camel, which has no separator to preserve it.
func TestDigitWordsAreLossyUnderCamel(t *testing.T) {
	const in = "web 2 print"
	if got := ToSnake(in); got != "web_2_print" {
		t.Fatalf("ToSnake(%q) = %q, want web_2_print", in, got)
	}
	camel := ToCamel(in) // "web2Print"
	if got := ToSnake(camel); got == "web_2_print" {
		t.Errorf("expected glued-digit loss through camel, but %q round-tripped", camel)
	}
}

// TestNoPanicOnPathologicalInput: a battery of awkward strings must convert
// without panicking. Fuzzing extends this; the explicit list documents intent.
func TestNoPanicOnPathologicalInput(t *testing.T) {
	pathological := []string{
		"", " ", "\t\n", "____", "----", "....", "\x00", "\xff\xfe",
		strings.Repeat("A", 1000), strings.Repeat("aB", 500),
		strings.Repeat("_", 1000), "👍🏽emoji👍test", "​zero​width",
		"ﬀ ligature", "Ⅷ roman", "½ fraction",
	}
	for _, s := range pathological {
		for _, style := range allStyles {
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("Convert(%v, %q) panicked: %v", style, s, r)
					}
				}()
				Convert(style, s)
			}()
		}
	}
}
