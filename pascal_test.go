package scs

import "testing"

// TestStrIsPascal tests StrIsPascal function.
func TestStrIsPascal(t *testing.T) {
	tests := []struct {
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

// TestStrIsPascalExample tests StrIsPascal function.
func TestStrIsPascalExample(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected bool
	}{
		{"HelloWorld", true},
		{"helloWorld", false},
		{"hello_world", false},
		{"Hello_World", false},
	}

	for _, test := range tests {
		result := StrIsPascal(test.input)
		if result != test.expected {
			t.Errorf("StrIsPascal(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

// TestStrToPascal tests StrToPascal function.
func TestStrToPascal(t *testing.T) {
	tests := []struct {
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

// TestStrToPascalExample tests StrToPascal function.
func TestStrToPascalExample(t *testing.T) {
	result := StrToPascal("hello_world")
	expected := "HelloWorld"
	if result != expected {
		t.Errorf("StrToPascal(\"hello_world\") returned %s, expected %s",
			result, expected)
	}

	result = StrToPascal("hello world")
	expected = "HelloWorld"
	if result != expected {
		t.Errorf("StrToPascal(\"hello world\") returned %s, expected %s",
			result, expected)
	}

	result = StrToPascal("helloWorld")
	expected = "Helloworld"
	if result != expected {
		t.Errorf("StrToPascal(\"helloWorld\") returned %s, expected %s",
			result, expected)
	}
}

// TestToPascal tests ToPascal function.
func TestToPascal(t *testing.T) {
	tests := []struct {
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

// TestToPascalExample tests ToPascal function.
func TestToPascalExample(t *testing.T) {
	t.Run("Converts snake_case to PascalCase", func(t *testing.T) {
		result := ToPascal("hello_world")
		expected := "HelloWorld"
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})

	t.Run("Converts kebab-case to PascalCase", func(t *testing.T) {
		result := ToPascal("hello-world")
		expected := "HelloWorld"
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})

	t.Run("Converts camelCase to PascalCase", func(t *testing.T) {
		result := ToPascal("helloWorld")
		expected := "HelloWorld"
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})

	t.Run("Handles strings already in PascalCase", func(t *testing.T) {
		result := ToPascal("HelloWorld")
		expected := "HelloWorld"
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})

	t.Run("Handles strings with alphanumeric characters", func(t *testing.T) {
		result := ToPascal("helloWorld123")
		expected := "HelloWorld123"
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})
}

// TestPascalToKebab tests PascalToKebab function.
func TestPascalToKebab(t *testing.T) {
	tests := []struct {
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

// TestPascalToKebabExample tests PascalToKebab function.
func TestPascalToKebabExample(t *testing.T) {
	t.Run("Converts PascalCase to kebab-case", func(t *testing.T) {
		result, err := PascalToKebab("HelloWorld")
		expected := "hello-world"
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})

	t.Run("Handles PascalCase with numbers", func(t *testing.T) {
		result, err := PascalToKebab("HelloWorld123")
		expected := "hello-world-123"
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})

	t.Run("Handles single-word PascalCase", func(t *testing.T) {
		result, err := PascalToKebab("Hello")
		expected := "hello"
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})

	t.Run("Handles non-PascalCase input", func(t *testing.T) {
		result, err := PascalToKebab("helloWorld")
		expected := ""
		if err == nil {
			t.Error("Expected an error, but got nil")
		}
		if result != expected {
			t.Errorf("Expected %q, but got %q", expected, result)
		}
	})
}

// TestPascalToKebabError tests PascalToKebab function with wrong value.
func TestPascalToKebabError(t *testing.T) {
	notPascal := "one_two_three"

	_, err := PascalToKebab(notPascal)
	if err == nil {
		t.Error("not pascal to kebab")
	}
}

// TestPascalToCamel tests PascalToCamel function.
func TestPascalToCamel(t *testing.T) {
	tests := []struct {
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
	notPascal := "one_two_three"

	_, err := PascalToCamel(notPascal)
	if err == nil {
		t.Error("not pascal to camel")
	}
}

// TestPascalToSnake tests PascalToSnake function.
func TestPascalToSnake(t *testing.T) {
	tests := []struct {
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
	notPascal := "one_two_three"

	_, err := PascalToSnake(notPascal)
	if err == nil {
		t.Error("not pascal to snake")
	}
}
