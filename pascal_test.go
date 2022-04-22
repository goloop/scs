package scs

import "testing"

// TestStrIsPascal tests StrIsPascal function.
func TestStrIsPascal(t *testing.T) {
	var tests = []struct {
		value  string
		result bool
	}{
		// Simple examples
		{"One", true},
		{"oneTwoThree", false},
		{"OneTwoThree", true},
		{"ice9", false},
		{"Ice9", true},

		// Examples with abbreviations
		{"isWWWConnection", false},
		{"IsWWWConnection", true},
		{"HTTPToHTTPS", true},
		{"isHTTPOrHTTPS", false},
		{"HTTPToHTTPS", true},
		{"Http_To_HTTPS", false},
		{"Http-To-HTTPS", false},
	}

	for i, s := range tests {
		if r := StrIsPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %t but %t", i, s.result, r)
		}
	}
}

// TestStrToPascal tests StrToPascal function.
func TestStrToPascal(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "One"},
		{" One two Three ", "OneTwoThree"},
		{"Ice 9", "Ice9"},

		// Examples with abbreviations
		{"is www Connection", "IsWWWConnection"},
		{"http to https", "HTTPToHTTPS"},
		{"is http or https", "IsHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r := StrToPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestToPascal tests ToPascal function.
func TestToPascal(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "One"},
		{" One two Three ", "OneTwoThree"},
		{"Ice 9", "Ice9"},

		// Examples with abbreviations
		{"isWWWConnection", "IsWWWConnection"},
		{"http-to-https", "HTTPToHTTPS"},
		{"is_http_or_https", "IsHTTPOrHTTPS"},
		{"IsHTTPOrHTTPS", "IsHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r := ToPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestPascalToKebab tests PascalToKebab function.
func TestPascalToKebab(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{"OneTwoThree", "one-two-three"},
		{"Ice9", "ice-9"},

		// Examples with abbreviations
		{"IsWWWConnection", "is-www-connection"},
		{"HTTPToHTTPS", "http-to-https"},
		{"IsHTTPOrHTTPS", "is-http-or-https"},
	}

	for i, s := range tests {
		if r, _ := PascalToKebab(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestPascalToKebabError tests PascalToKebab function with wrong value.
func TestPascalToKebabError(t *testing.T) {
	var notPascal = "one_two_three"

	_, err := PascalToKebab(notPascal)
	if err == nil {
		t.Error("not pascal to kebab")
	}
}

// TestPascalToCamel tests PascalToCamel function.
func TestPascalToCamel(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{"OneTwoThree", "oneTwoThree"},
		{"Ice9", "ice9"},

		// Examples with abbreviations
		{"IsWWWConnection", "isWWWConnection"},
		{"HTTPToHTTPS", "httpToHTTPS"},
		{"IsHTTPOrHTTPS", "isHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r, _ := PascalToCamel(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestPascalToCamelError tests PascalToCamel function with wrong value.
func TestPascalToCamelError(t *testing.T) {
	var notPascal = "one_two_three"

	_, err := PascalToCamel(notPascal)
	if err == nil {
		t.Error("not pascal to camel")
	}
}

// TestPascalToSnake tests PascalToSnake function.
func TestPascalToSnake(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{"OneTwoThree", "one_two_three"},
		{"Ice9", "ice_9"},

		// Examples with abbreviations
		{"IsWWWConnection", "is_www_connection"},
		{"HTTPToHTTPS", "http_to_https"},
		{"IsHTTPOrHTTPS", "is_http_or_https"},
	}

	for i, s := range tests {
		if r, _ := PascalToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestPascalToSnakeError tests PascalToSnake function with wrong value.
func TestPascalToSnakeError(t *testing.T) {
	var notPascal = "one_two_three"

	_, err := PascalToSnake(notPascal)
	if err == nil {
		t.Error("not pascal to snake")
	}
}
