package scs

import "testing"

// TestStrIsKebab tests StrIsKebab function.
func TestStrIsKebab(t *testing.T) {
	var tests = []struct {
		value  string
		result bool
	}{
		// Simple examples
		{"One", false},
		{"one", true},
		{"oneTwoThree", false},
		{"one-two-three", true},
		{"OneTwoThree", false},
		{"ice9", true},
		{"Ice9", false},

		// Examples with abbreviations
		{"is-www-connection", true},
		{"IsWWWConnection", false},
		{"HTTPToHTTPS", false},
		{"isHTTPOrHTTPS", false},
		{"HTTPToHTTPS", false},
		{"http_to_https", false},
		{"http-to-https", true},
	}

	for i, s := range tests {
		if r := StrIsKebab(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %t but %t", i, s.result, r)
		}
	}
}

// TestStrToKebab tests StrToKebab function.
func TestStrToKebab(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "one-two-three"},
		{"Ice 9", "ice-9"},

		// Examples with abbreviations
		{"is www Connection", "is-www-connection"},
		{"http to https", "http-to-https"},
		{"is http or https", "is-http-or-https"},
	}

	for i, s := range tests {
		if r := StrToKebab(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestToKebab tests ToKebab function.
func TestToKebab(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "one-two-three"},
		{"Ice 9", "ice-9"},

		// Examples with abbreviations
		{"is_www_connection", "is-www-connection"},
		{"httpToHTTPS", "http-to-https"},
		{"IsHTTPOrHTTPS", "is-http-or-https"},
		{"is-http-or-https", "is-http-or-https"},
	}

	for i, s := range tests {
		if r := ToKebab(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestKebabToCamel tests KebabToCamel function.
func TestKebabToCamel(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"one-two-three", "oneTwoThree"},
		{"ice-9", "ice9"},

		// Examples with abbreviations
		{"is-www-connection", "isWWWConnection"},
		{"http-to-https", "httpToHTTPS"},
		{"is-http-or-https", "isHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r, _ := KebabToCamel(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestKebabToCamelError tests KebabToCamel function with wrong value.
func TestKebabToCamelError(t *testing.T) {
	var notKebab = "oneTwoThree"

	_, err := KebabToCamel(notKebab)
	if err == nil {
		t.Error("not kebab to camel")
	}
}

// TestKebabToSnake tests KebabToSnake function.
func TestKebabToSnake(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"one-two-three", "one_two_three"},
		{"ice-9", "ice_9"},

		// Examples with abbreviations
		{"is-www-connection", "is_www_connection"},
		{"http-to-https", "http_to_https"},
		{"is-http-or-https", "is_http_or_https"},
	}

	for i, s := range tests {
		if r, _ := KebabToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestKebabToSnakeError tests KebabToSnake function with wrong value.
func TestKebabToSnakeError(t *testing.T) {
	var notKebab = "oneTwoThree"

	_, err := KebabToSnake(notKebab)
	if err == nil {
		t.Error("not kebab to snake")
	}
}

// TestKebabToPascal tests KebabToPascal function.
func TestKebabToPascal(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "One"},
		{"one-two-three", "OneTwoThree"},
		{"ice-9", "Ice9"},

		// Examples with abbreviations
		{"is-www-connection", "IsWWWConnection"},
		{"http-to-https", "HTTPToHTTPS"},
		{"is-http-or-https", "IsHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r, _ := KebabToPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestKebabToPascalError tests KebabToPascal function with wrong value.
func TestKebabToPascalError(t *testing.T) {
	var notKebab = "oneTwoThree"

	_, err := KebabToPascal(notKebab)
	if err == nil {
		t.Error("not kebab to pascal")
	}
}
