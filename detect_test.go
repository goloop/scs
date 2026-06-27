package scs

import "testing"

func TestDetect(t *testing.T) {
	tests := []struct {
		in    string
		style Style
		ok    bool
	}{
		{"user_id", Snake, true},
		{"userId", Camel, true},
		{"UserId", Pascal, true},
		{"user-id", Kebab, true},
		{"user.id", Dot, true},
		{"USER_ID", ScreamingSnake, true},
		{"User Id", Title, true},

		// Ambiguous: a bare lowercase word is canonical in several styles.
		{"api", Unknown, false},
		{"hello", Unknown, false},

		// Not canonical in any single style.
		{"Hello_World", Unknown, false},
		{"hello__world", Unknown, false}, // double separator is not canonical
		{"", Unknown, false},
		{"   ", Unknown, false},

		// Multi-word disambiguates even when one word would be ambiguous.
		{"http_api", Snake, true},
		{"httpApi", Camel, true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			style, ok := Detect(tt.in)
			if style != tt.style || ok != tt.ok {
				t.Errorf("Detect(%q) = (%v, %v), want (%v, %v)",
					tt.in, style, ok, tt.style, tt.ok)
			}
		})
	}
}

// referenceDetect is the obvious, un-pruned definition of Detect: probe every
// style and commit only on a unique match. The optimized Detect must agree with
// it on every input — this guards the separator-based candidate pruning.
func referenceDetect(s string) (Style, bool) {
	if s == "" {
		return Unknown, false
	}
	found := Unknown
	for _, style := range allStyles {
		if Convert(style, s) == s {
			if found != Unknown {
				return Unknown, false
			}
			found = style
		}
	}
	if found == Unknown {
		return Unknown, false
	}
	return found, true
}

// TestDetectMatchesReference proves the pruned Detect is equivalent to the
// brute-force reference across the full corpus and a batch of separator-mixing
// and punctuation edge cases.
func TestDetectMatchesReference(t *testing.T) {
	inputs := append([]string{}, corpus...)
	inputs = append(inputs,
		"a_b-c", "a.b c", "a/b", "user@id", "mixed_Case", "USER_ID", "user-id",
		"user.id", "User Id", "User id", "a_b_c", "A_B_C", "café_au_lait",
	)
	for _, s := range inputs {
		gotStyle, gotOK := Detect(s)
		wantStyle, wantOK := referenceDetect(s)
		if gotStyle != wantStyle || gotOK != wantOK {
			t.Errorf("Detect(%q) = (%v, %v); reference = (%v, %v)",
				s, gotStyle, gotOK, wantStyle, wantOK)
		}
	}
}

// TestDetectSentence: Sentence case is now detectable and distinct from Title.
func TestDetectSentence(t *testing.T) {
	if style, ok := Detect("Hello world"); style != Sentence || !ok {
		t.Errorf("Detect(%q) = (%v, %v), want (Sentence, true)", "Hello world", style, ok)
	}
	if style, ok := Detect("Hello World"); style != Title || !ok {
		t.Errorf("Detect(%q) = (%v, %v), want (Title, true)", "Hello World", style, ok)
	}
}

// TestDetectAgreesWithIs: whenever Detect commits to a style, that style's Is
// predicate must agree. Detect is the unambiguous tightening of Is.
func TestDetectAgreesWithIs(t *testing.T) {
	inputs := []string{
		"user_id", "userId", "UserId", "user-id", "user.id", "USER_ID",
		"User Id", "api", "Hello_World", "httpServerId",
	}
	for _, in := range inputs {
		if style, ok := Detect(in); ok && !Is(style, in) {
			t.Errorf("Detect(%q) committed to %v but Is(%v, %q) is false",
				in, style, style, in)
		}
	}
}

func TestIsPredicates(t *testing.T) {
	tests := []struct {
		fn   func(string) bool
		name string
		in   string
		want bool
	}{
		{IsCamel, "IsCamel", "helloWorld", true},
		{IsCamel, "IsCamel", "HelloWorld", false},
		{IsCamel, "IsCamel", "", false},
		{IsPascal, "IsPascal", "HelloWorld", true},
		{IsPascal, "IsPascal", "helloWorld", false},
		{IsSnake, "IsSnake", "hello_world", true},
		{IsSnake, "IsSnake", "Hello_World", false},
		{IsKebab, "IsKebab", "hello-world", true},
		{IsKebab, "IsKebab", "Hello-World", false},
		{IsScreamingSnake, "IsScreamingSnake", "HELLO_WORLD", true},
		{IsScreamingSnake, "IsScreamingSnake", "hello_world", false},
		{IsDot, "IsDot", "hello.world", true},
		{IsTitle, "IsTitle", "Hello World", true},
		{IsTitle, "IsTitle", "hello world", false},
		{IsSentence, "IsSentence", "Hello world", true},
		{IsSentence, "IsSentence", "Hello World", false},

		// A single lowercase word is canonical in every caseless style.
		{IsCamel, "IsCamel", "api", true},
		{IsSnake, "IsSnake", "api", true},
		{IsKebab, "IsKebab", "api", true},
	}
	for _, tt := range tests {
		t.Run(tt.name+"/"+tt.in, func(t *testing.T) {
			if got := tt.fn(tt.in); got != tt.want {
				t.Errorf("%s(%q) = %v, want %v", tt.name, tt.in, got, tt.want)
			}
		})
	}
}

// TestIsInvalidStyle: the generic Is must reject an invalid style.
func TestIsInvalidStyle(t *testing.T) {
	if Is(Unknown, "hello") {
		t.Error("Is(Unknown, ...) must be false")
	}
	if Is(Style(99), "hello") {
		t.Error("Is(invalid, ...) must be false")
	}
}
