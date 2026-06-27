package scs

import (
	"sync"
	"testing"
)

func TestCaserWithAcronyms(t *testing.T) {
	c := New(WithAcronyms("ID", "URL", "HTTP", "API"))

	tests := []struct {
		fn   func(string) string
		name string
		in   string
		want string
	}{
		// camelCase keeps the first word lowercase even when it is an
		// acronym; later acronyms are upper-cased.
		{c.ToCamel, "ToCamel", "user_id", "userID"},
		{c.ToCamel, "ToCamel", "http_server", "httpServer"},
		{c.ToCamel, "ToCamel", "id_token", "idToken"},

		// PascalCase upper-cases every acronym, including the first.
		{c.ToPascal, "ToPascal", "user_id", "UserID"},
		{c.ToPascal, "ToPascal", "url_value", "URLValue"},

		// Title keeps spaces between upper-cased acronyms.
		{c.ToTitle, "ToTitle", "user id", "User ID"},

		// Sentence capitalizes only the first word, but still upper-cases
		// initialisms wherever they appear.
		{c.ToSentence, "ToSentence", "get http url now", "Get HTTP URL now"},
		{c.ToSentence, "ToSentence", "id of user", "ID of user"},

		// Separated styles ignore the acronym set entirely.
		{c.ToSnake, "ToSnake", "HTTPServer", "http_server"},
		{c.ToKebab, "ToKebab", "userID", "user-id"},
		{c.ToScreamingSnake, "ToScreamingSnake", "userID", "USER_ID"},
	}
	for _, tt := range tests {
		t.Run(tt.name+"/"+tt.in, func(t *testing.T) {
			if got := tt.fn(tt.in); got != tt.want {
				t.Errorf("%s(%q) = %q, want %q", tt.name, tt.in, got, tt.want)
			}
		})
	}
}

// TestDefaultCaserHasNoAcronyms: the package functions must not magically know
// any initialisms — that is exactly the ambiguity v2 avoids by default.
func TestDefaultCaserHasNoAcronyms(t *testing.T) {
	if got := ToPascal("user_id"); got != "UserId" {
		t.Errorf("default ToPascal(user_id) = %q, want %q", got, "UserId")
	}
	if got := New().ToPascal("user_id"); got != "UserId" {
		t.Errorf("New().ToPascal(user_id) = %q, want %q", got, "UserId")
	}
}

// TestWithAcronymsNormalizesInput: acronyms are matched case-insensitively and
// trimmed, and empty entries are ignored.
func TestWithAcronymsNormalizesInput(t *testing.T) {
	c := New(WithAcronyms("  id  ", "Url", "", "   "))
	if got := c.ToPascal("user_id_url"); got != "UserIDURL" {
		t.Errorf("ToPascal(user_id_url) = %q, want %q", got, "UserIDURL")
	}
}

// TestWithAcronymsEmpty: passing no words, or only blank words, leaves the
// Caser in its default (no-initialism) state instead of allocating an empty
// set or panicking.
func TestWithAcronymsEmpty(t *testing.T) {
	for _, c := range []*Caser{
		New(WithAcronyms()),
		New(WithAcronyms("", "   ")),
		New(),
	} {
		if c.acronyms != nil {
			t.Errorf("expected no acronym set, got %v", c.acronyms)
		}
		if got := c.ToPascal("user_id"); got != "UserId" {
			t.Errorf("ToPascal(user_id) = %q, want %q", got, "UserId")
		}
	}
}

// TestCaserConvertMethod: the Caser.Convert method honors the acronym set and
// is total for an invalid style, mirroring the package-level Convert.
func TestCaserConvertMethod(t *testing.T) {
	c := New(WithAcronyms("ID"))
	if got := c.Convert(Pascal, "user_id"); got != "UserID" {
		t.Errorf("c.Convert(Pascal, user_id) = %q, want %q", got, "UserID")
	}
	if got := c.Convert(Unknown, "user_id"); got != "user_id" {
		t.Errorf("c.Convert(Unknown, ...) = %q, want unchanged", got)
	}
}

// TestCaserConcurrentUse: a Caser is immutable, so concurrent reads must be
// race-free. Run with -race to make this meaningful.
func TestCaserConcurrentUse(t *testing.T) {
	c := New(WithAcronyms("ID", "HTTP"))
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				if c.ToPascal("http_id") != "HTTPID" {
					t.Error("unexpected concurrent result")
					return
				}
			}
		}()
	}
	wg.Wait()
}

// TestAdjacentAcronymsAreDocumentedLossy pins the trade-off WithAcronyms makes:
// adjacent initialisms merge into an unsplittable run. This is why Title casing
// is the round-trip-safe default.
func TestAdjacentAcronymsAreDocumentedLossy(t *testing.T) {
	c := New(WithAcronyms("HTTP", "API", "ID"))
	got := c.ToPascal("http_api_id")
	if got != "HTTPAPIID" {
		t.Fatalf("ToPascal(http_api_id) = %q, want %q", got, "HTTPAPIID")
	}
	// The merged run can no longer be separated back into the original words.
	if back := ToSnake(got); back == "http_api_id" {
		t.Errorf("expected lossy round-trip for adjacent acronyms, got reversible %q", back)
	}
}
