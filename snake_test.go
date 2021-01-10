package scs

import "testing"

// TestStrIsSnake tests StrIsSnake function.
func TestStrIsSnake(t *testing.T) {
	var tests = []struct {
		value  string
		result bool
	}{
		// Simple examples
		{"One", false},
		{"one", true},
		{"oneTwoThree", false},
		{"one_two_three", true},
		{"OneTwoThree", false},
		{"ice9", true},
		{"Ice9", false},

		// Examples with acronyms
		{"is_www_connection", true},
		{"IsWWWConnection", false},
		{"HTTPToHTTPS", false},
		{"isHTTPOrHTTPS", false},
		{"HTTPToHTTPS", false},
		{"http_to_https", true},
		{"http-to-https", false},
	}

	for i, s := range tests {
		if r := StrIsSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %t but %t", i, s.result, r)
		}
	}
}

// TestStrToSnake tests StrToSnake function.
func TestStrToSnake(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "one_two_three"},
		{"Ice 9", "ice_9"},

		// Examples with acronyms
		{"is www Connection", "is_www_connection"},
		{"http to https", "http_to_https"},
		{"is http or https", "is_http_or_https"},
	}

	for i, s := range tests {
		if r := StrToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestToSnake tests ToSnake function.
func TestToSnake(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "one_two_three"},
		{"Ice 9", "ice_9"},

		// Examples with acronyms
		{"isWWWConnection", "is_www_connection"},
		{"HTTPToHTTPS", "http_to_https"},
		{"is-http-or-https", "is_http_or_https"},
	}

	for i, s := range tests {
		if r := ToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestSnakeToCamel tests SnakeToCamel function.
func TestSnakeToCamel(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"one_two_three", "oneTwoThree"},
		{"ice_9", "ice9"},

		// Examples with acronyms
		{"is_www_connection", "isWWWConnection"},
		{"http_to_https", "httpToHTTPS"},
		{"is_http_or_https", "isHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r, _ := SnakeToCamel(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestSnakeToKebab tests SnakeToKebab function.
func TestSnakeToKebab(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"one_two_three", "one-two-three"},
		{"ice_9", "ice-9"},

		// Examples with acronyms
		{"is_www_connection", "is-www-connection"},
		{"http_to_https", "http-to-https"},
		{"is_http_or_https", "is-http-or-https"},
	}

	for i, s := range tests {
		if r, _ := SnakeToKebab(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestSnakeToPascal tests SnakeToPascal function.
func TestSnakeToPascal(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "One"},
		{"one_two_three", "OneTwoThree"},
		{"ice_9", "Ice9"},

		// Examples with acronyms
		{"is_www_connection", "IsWWWConnection"},
		{"http_to_https", "HTTPToHTTPS"},
		{"is_http_or_https", "IsHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r, _ := SnakeToPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}
