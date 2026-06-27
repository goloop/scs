package scs

import (
	"reflect"
	"slices"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		// Empty and separator-only inputs yield no words.
		{"empty", "", nil},
		{"spaces", "   ", nil},
		{"punct only", "-_-.//", nil},

		// Single words.
		{"one lower", "hello", []string{"hello"}},
		{"one upper", "HELLO", []string{"hello"}},
		{"one title", "Hello", []string{"hello"}},

		// Separator-based splitting, mixed delimiters collapse.
		{"snake", "hello_world", []string{"hello", "world"}},
		{"kebab", "hello-world", []string{"hello", "world"}},
		{"dot", "hello.world", []string{"hello", "world"}},
		{"spaces sep", "hello world", []string{"hello", "world"}},
		{"mixed seps", "hello__world--foo..bar", []string{"hello", "world", "foo", "bar"}},
		{"leading/trailing seps", "__hello__", []string{"hello"}},

		// Case-based splitting.
		{"camel", "helloWorld", []string{"hello", "world"}},
		{"pascal", "HelloWorld", []string{"hello", "world"}},
		{"three camel", "oneTwoThree", []string{"one", "two", "three"}},

		// Acronym boundaries: the trailing capital that starts a lowercase
		// word breaks away from the acronym.
		{"acronym+word", "HTTPServer", []string{"http", "server"}},
		{"word+acronym", "userID", []string{"user", "id"}},
		{"acronym+word+acronym", "HTTPServerID", []string{"http", "server", "id"}},
		{"two acronyms spaced", "HTTP API", []string{"http", "api"}},
		{"io reader", "IOReader", []string{"io", "reader"}},
		{"json data", "JSONData", []string{"json", "data"}},

		// Digits glue to letters; only explicit separators split them out.
		{"digit glued", "web2print", []string{"web2print"}},
		{"identifier nums", "utf8 sha256 oauth2", []string{"utf8", "sha256", "oauth2"}},
		{"digit then upper", "version2Final", []string{"version2", "final"}},
		{"digit word by sep", "web 2 print", []string{"web", "2", "print"}},
		{"trailing digits", "base64", []string{"base64"}},
		{"upper digits", "ID2", []string{"id2"}},

		// Unicode: caseless scripts glue; cased scripts split like ASCII.
		{"cyrillic camel", "ПриватБанк", []string{"приват", "банк"}},
		{"cyrillic snake", "привіт_світ", []string{"привіт", "світ"}},
		{"umlaut", "straßeTest", []string{"straße", "test"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Split(tt.in)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split(%q) = %#v, want %#v", tt.in, got, tt.want)
			}
		})
	}
}

// TestWordsMatchesSplit guarantees the iterator and the slice helper never
// diverge: they are two faces of the same tokenizer.
func TestWordsMatchesSplit(t *testing.T) {
	inputs := []string{
		"", "  ", "helloWorld", "HTTPServerID", "user_id-value.test",
		"web2print", "ПриватБанк", "OAuth2Provider", "a", "A-B-C",
	}
	for _, in := range inputs {
		want := Split(in)
		got := slices.Collect(Words(in))
		if len(want) == 0 && len(got) == 0 {
			continue
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Words(%q) = %#v, Split = %#v", in, got, want)
		}
	}
}

// TestWordsEarlyStop verifies the iterator honors an early break instead of
// always walking the whole string.
func TestWordsEarlyStop(t *testing.T) {
	var seen []string
	for w := range Words("one two three four") {
		seen = append(seen, w)
		if len(seen) == 2 {
			break
		}
	}
	want := []string{"one", "two"}
	if !reflect.DeepEqual(seen, want) {
		t.Errorf("early-stopped Words = %#v, want %#v", seen, want)
	}
}

// TestSplitProducesLowercaseWords is an invariant: every token returned by the
// tokenizer is already normalized to lowercase, so renderers never have to
// re-lower.
func TestSplitProducesLowercaseWords(t *testing.T) {
	inputs := []string{"HTTPServerID", "ПриватБанк", "STRAßE", "MixedCASE123"}
	for _, in := range inputs {
		for _, w := range Split(in) {
			if w != normalize(w) {
				t.Errorf("Split(%q) produced non-normalized word %q", in, w)
			}
		}
	}
}
