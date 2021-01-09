package scs

import "testing"

// TestToCamel tests ToCamel function.
func TestToCamel(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "oneTwoThree"},
		{"Ice 9", "ice9"},

		// Examples with acronyms
		{"is www Connection", "isWWWConnection"},
		{"http to https", "httpToHTTPS"},
		{"is http or https", "isHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r := ToCamel(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestCamelToKebab tests CamelToKebab function.
func TestCamelToKebab(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"oneTwoThree", "one-two-three"},
		{"ice9", "ice9"},

		// Examples with acronyms
		{"isWWWConnection", "is-www-connection"},
		{"httpToHTTPS", "http-to-https"},
		{"isHTTPOrHTTPS", "is-http-or-https"},
	}

	for i, s := range tests {
		if r := CamelToKebab(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestCamelToPascal tests CamelToPascal function.
func TestCamelToPascal(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "One"},
		{"oneTwoThree", "OneTwoThree"},
		{"ice9", "Ice9"},

		// Examples with acronyms
		{"isWWWConnection", "IsWWWConnection"},
		{"httpToHTTPS", "HTTPToHTTPS"},
		{"isHTTPOrHTTPS", "IsHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r := CamelToPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestCamelToSnake tests CamelToSnake function.
func TestCamelToSnake(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"oneTwoThree", "one_two_three"},
		{"ice9", "ice9"},

		// Examples with acronyms
		{"isWWWConnection", "is_www_connection"},
		{"httpToHTTPS", "http_to_https"},
		{"isHTTPOrHTTPS", "is_http_or_https"},
	}

	for i, s := range tests {
		if r := CamelToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}
