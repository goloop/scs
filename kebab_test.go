package scs

import "testing"

// TestToKebab tests CamelToKebab function.
func TestToKebab(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "one-two-three"},
		{"Ice 9", "ice-9"},

		// Examples with acronyms
		{"is www Connection", "is-www-connection"},
		{"http to https", "http-to-https"},
		{"is http or https", "is-http-or-https"},
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

		// Examples with acronyms
		{"is-www-connection", "isWWWConnection"},
		{"http-to-https", "httpToHTTPS"},
		{"is-http-or-https", "isHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r := KebabToCamel(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
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

		// Examples with acronyms
		{"is-www-connection", "is_www_connection"},
		{"http-to-https", "http_to_https"},
		{"is-http-or-https", "is_http_or_https"},
	}

	for i, s := range tests {
		if r := KebabToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
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

		// Examples with acronyms
		{"is-www-connection", "IsWWWConnection"},
		{"http-to-https", "HTTPToHTTPS"},
		{"is-http-or-https", "IsHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r := KebabToPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}
