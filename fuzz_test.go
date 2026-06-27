package scs

import (
	"slices"
	"testing"
	"unicode/utf8"
)

// FuzzConvertInvariants exercises every converter on arbitrary input and
// asserts the package's core contracts hold for all of them:
//
//   - converters never panic and always emit valid UTF-8 (every style);
//   - snake_case is a faithful carrier: routing any input through it does not
//     change any rendering (every style);
//   - snake/kebab/dot are idempotent, canonical and word-preserving (stable
//     styles), for any input at all.
//
// The cased and all-caps styles are not asserted idempotent here because
// adjacent single capitals are inherently ambiguous with acronyms; that limit
// is pinned by TestCasedStylesAmbiguityIsDocumented.
func FuzzConvertInvariants(f *testing.F) {
	seeds := []string{
		"", " ", "helloWorld", "HTTPServerID", "user_id", "web2print",
		"ПриватБанк", "straßeTest", "A-b_c.D", "____", "OAuth2Client",
		"\x00\xff", "MixedCASEhere123",
	}
	for _, s := range seeds {
		f.Add(s)
	}

	stable := map[Style]bool{Snake: true, Kebab: true, Dot: true}

	f.Fuzz(func(t *testing.T, s string) {
		words := Split(s)
		snake := ToSnake(s)

		for _, style := range allStyles {
			out := Convert(style, s)

			// Output must always be valid UTF-8.
			if !utf8.ValidString(out) {
				t.Fatalf("Convert(%v, %q) produced invalid UTF-8: %q", style, s, out)
			}

			// snake_case is a canonical carrier for every style.
			if viaSnake := Convert(style, snake); viaSnake != out {
				t.Fatalf("Convert(%v, %q)=%q != Convert(%v, snake)=%q",
					style, s, out, style, viaSnake)
			}

			if !stable[style] {
				continue
			}

			// Stable styles: idempotent, canonical, word-preserving.
			if again := Convert(style, out); again != out {
				t.Fatalf("Convert(%v) not idempotent on %q: %q vs %q",
					style, s, out, again)
			}
			if out != "" && !Is(style, out) {
				t.Fatalf("Is(%v, %q) false for Convert(%v, %q)", style, out, style, s)
			}
			if got := Split(out); !slices.Equal(got, words) {
				t.Fatalf("Split(Convert(%v, %q)) = %#v, want %#v",
					style, s, got, words)
			}
		}
	})
}

// FuzzDetectRoundTrip checks that whenever Detect commits to a style, the input
// really is that style's fixed point — Detect never lies.
func FuzzDetectRoundTrip(f *testing.F) {
	for _, s := range []string{"user_id", "userId", "API", "api", "Hello_World", ""} {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, s string) {
		style, ok := Detect(s)
		if !ok {
			if style != Unknown {
				t.Fatalf("Detect(%q) returned ok=false but style=%v", s, style)
			}
			return
		}
		if Convert(style, s) != s {
			t.Fatalf("Detect(%q)=%v but Convert(%v, %q) != %q", s, style, style, s, s)
		}
	})
}
