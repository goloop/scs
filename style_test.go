package scs

import "testing"

func TestStyleString(t *testing.T) {
	tests := []struct {
		style Style
		want  string
	}{
		{Camel, "camel"},
		{Pascal, "pascal"},
		{Snake, "snake"},
		{Kebab, "kebab"},
		{ScreamingSnake, "screaming_snake"},
		{Dot, "dot"},
		{Title, "title"},
		{Sentence, "sentence"},
		{Unknown, "unknown"},
		{Style(200), "unknown"},
	}
	for _, tt := range tests {
		if got := tt.style.String(); got != tt.want {
			t.Errorf("Style(%d).String() = %q, want %q", tt.style, got, tt.want)
		}
	}
}

func TestStyleValid(t *testing.T) {
	valid := []Style{Camel, Pascal, Snake, Kebab, ScreamingSnake, Dot, Title, Sentence}
	for _, s := range valid {
		if !s.Valid() {
			t.Errorf("Style %v should be Valid", s)
		}
	}
	for _, s := range []Style{Unknown, Style(99)} {
		if s.Valid() {
			t.Errorf("Style %v should not be Valid", s)
		}
	}
}

// TestStyleStringRoundTrip: String and ParseStyle are inverses for every valid
// style, so a style survives a config file or CLI flag intact.
func TestStyleStringRoundTrip(t *testing.T) {
	for s := Camel; s <= maxStyle; s++ {
		got, ok := ParseStyle(s.String())
		if !ok || got != s {
			t.Errorf("ParseStyle(%q) = (%v, %v), want (%v, true)", s.String(), got, ok, s)
		}
	}
}

func TestParseStyle(t *testing.T) {
	tests := []struct {
		in    string
		style Style
		ok    bool
	}{
		{"camel", Camel, true},
		{"screaming_snake", ScreamingSnake, true},
		{"screaming", ScreamingSnake, true},
		{"constant", ScreamingSnake, true},
		{"title", Title, true},
		{"sentence", Sentence, true},
		{"CAMEL", Unknown, false}, // exact, case-sensitive
		{"", Unknown, false},
		{"weird", Unknown, false},
	}
	for _, tt := range tests {
		style, ok := ParseStyle(tt.in)
		if style != tt.style || ok != tt.ok {
			t.Errorf("ParseStyle(%q) = (%v, %v), want (%v, %v)",
				tt.in, style, ok, tt.style, tt.ok)
		}
	}
}
