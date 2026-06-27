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
