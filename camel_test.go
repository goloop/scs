package scs

import "testing"

// TestStrIsCamel tests StrIsCamel function.
func TestStrIsCamel(t *testing.T) {
	var tests = []struct {
		value  string
		result bool
	}{
		// Simple examples
		{"One", false},
		{"oneTwoThree", true},
		{"OneTwoThree", false},
		{"ice9", true},
		{"Ice9", false},

		// Examples with abbreviations
		{"isWWWConnection", true},
		{"httpToHTTPS", true},
		{"isHTTPOrHTTPS", true},
		{"IsHTTPOrHTTPS", false},
		{"HTTPToHTTPS", false},
		{"http_To_HTTPS", false},
		{"http-To-HTTPS", false},
	}

	for i, s := range tests {
		if r := StrIsCamel(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %t but %t", i, s.result, r)
		}
	}
}

// TestStrToCamel tests StrToCamel function.
func TestStrToCamel(t *testing.T) {
	var tests = []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "oneTwoThree"},
		{"Ice 9", "ice9"},

		// Examples with abbreviations
		{"is www Connection", "isWWWConnection"},
		{"http to https", "httpToHTTPS"},
		{"is http or https", "isHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r := StrToCamel(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

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

		// Examples with abbreviations
		{"IsWWWConnection", "isWWWConnection"},
		{"http-to-https", "httpToHTTPS"},
		{"is_http_or_https", "isHTTPOrHTTPS"},
		{"httpToHTTPS", "httpToHTTPS"},
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
		{"ice9", "ice-9"},

		// Examples with abbreviations
		{"isWWWConnection", "is-www-connection"},
		{"httpToHTTPS", "http-to-https"},
		{"isHTTPOrHTTPS", "is-http-or-https"},
	}

	for i, s := range tests {
		if r, _ := CamelToKebab(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestCamelToKebabError tests CamelToKebab function with wrong value.
func TestCamelToKebabError(t *testing.T) {
	var notCamel = "one-two-three"

	_, err := CamelToKebab(notCamel)
	if err == nil {
		t.Error("not camel to kebab")
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

		// Examples with abbreviations
		{"isWWWConnection", "IsWWWConnection"},
		{"httpToHTTPS", "HTTPToHTTPS"},
		{"isHTTPOrHTTPS", "IsHTTPOrHTTPS"},
	}

	for i, s := range tests {
		if r, _ := CamelToPascal(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestCamelToPascalError tests CamelToPascal function with wrong value.
func TestCamelToPascalError(t *testing.T) {
	var notCamel = "one-two-three"

	_, err := CamelToPascal(notCamel)
	if err == nil {
		t.Error("not camel to pascal")
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
		{"ice9", "ice_9"},

		// Examples with abbreviations
		{"isWWWConnection", "is_www_connection"},
		{"httpToHTTPS", "http_to_https"},
		{"isHTTPOrHTTPS", "is_http_or_https"},
	}

	for i, s := range tests {
		if r, _ := CamelToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestCamelToSnakeError tests CamelToSnake function with wrong value.
func TestCamelToSnakeError(t *testing.T) {
	var notCamel = "one-two-three"

	_, err := CamelToSnake(notCamel)
	if err == nil {
		t.Error("not camel to snake")
	}
}
