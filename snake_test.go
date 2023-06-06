package scs

import "testing"

// TestStrIsSnake tests StrIsSnake function.
func TestStrIsSnake(t *testing.T) {
	tests := []struct {
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

		// Examples with abbreviations
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

// TestStrIsSnakeExample tests StrIsSnake function.
func TestStrIsSnakeExample(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "Snake case string",
			str:  "hello_world",
			want: true,
		},
		{
			name: "Non-snake case string (CamelCase)",
			str:  "HelloWorld",
			want: false,
		},
		{
			name: "Non-snake case string (kebab-case)",
			str:  "hello-world",
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := StrIsSnake(tc.str)

			if got != tc.want {
				t.Errorf("StrIsSnake(%q) = %v, want %v", tc.str, got, tc.want)
			}
		})
	}
}

// TestStrToSnake tests StrToSnake function.
func TestStrToSnake(t *testing.T) {
	tests := []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "one_two_three"},
		{"Ice 9", "ice_9"},

		// Examples with abbreviations
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

// TestStrToSnakeExample tests StrToSnake function.
func TestStrToSnakeExample(t *testing.T) {
	testCases := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "Converts a string with spaces to snake_case",
			str:  "Hello World",
			want: "hello_world",
		},
		{
			name: "Lowercase string with spaces to snake_case",
			str:  "hello world",
			want: "hello_world",
		},
		{
			name: "Converts a kebab-case string to snake_case",
			str:  "hello-world",
			want: "hello_world",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := StrToSnake(tc.str)

			if got != tc.want {
				t.Errorf("StrToSnake(%q) = %q, want %q", tc.str, got, tc.want)
			}
		})
	}
}

// TestToSnake tests ToSnake function.
func TestToSnake(t *testing.T) {
	tests := []struct {
		value  string
		result string
	}{
		// Simple examples
		{"One", "one"},
		{" One two Three ", "one_two_three"},
		{"Ice 9", "ice_9"},

		// Examples with abbreviations
		{"isWWWConnection", "is_www_connection"},
		{"HTTPToHTTPS", "http_to_https"},
		{"is-http-or-https", "is_http_or_https"},
		{"is_http_or_https", "is_http_or_https"},
	}

	for i, s := range tests {
		if r := ToSnake(s.value); s.result != r {
			t.Errorf("test for %d is failed, "+
				"expected %s but %s", i, s.result, r)
		}
	}
}

// TestToSnakeExample tests ToSnake function.
func TestToSnakeExample(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "CamelCase to snake_case",
			arg:  "helloWorld",
			want: "hello_world",
		},
		{
			name: "PascalCase to snake_case",
			arg:  "HelloWorld",
			want: "hello_world",
		},
		{
			name: "Kebab-case to snake_case",
			arg:  "hello-world",
			want: "hello_world",
		},
		{
			name: "snake_case remains the same",
			arg:  "hello_world",
			want: "hello_world",
		},
		{
			name: "Space separated words to snake_case",
			arg:  "Hello World",
			want: "hello_world",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToSnake(tt.arg); got != tt.want {
				t.Errorf("ToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSnakeToCamel tests SnakeToCamel function.
func TestSnakeToCamel(t *testing.T) {
	tests := []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"one_two_three", "oneTwoThree"},
		{"ice_9", "ice9"},

		// Examples with abbreviations
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

// TestSnakeToCamelExample tests SnakeToCamel function.
func TestSnakeToCamelExample(t *testing.T) {
	tests := []struct {
		name    string
		snake   string
		want    string
		wantErr bool
	}{
		{
			name:    "Snake case string",
			snake:   "hello_world",
			want:    "helloWorld",
			wantErr: false,
		},
		{
			name:    "Camel case string",
			snake:   "HelloWorld",
			want:    "",
			wantErr: true,
		},
		{
			name:    "Kebab case string",
			snake:   "hello-world",
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SnakeToCamel(tt.snake)
			if (err != nil) != tt.wantErr {
				t.Errorf("SnakeToCamel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SnakeToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSnakeToCamelError tests SnakeToCamel function with wrong value.
func TestSnakeToCamelError(t *testing.T) {
	notSnake := "oneTwoThree"

	_, err := SnakeToCamel(notSnake)
	if err == nil {
		t.Error("not snake to camel")
	}
}

// TestSnakeToKebab tests SnakeToKebab function.
func TestSnakeToKebab(t *testing.T) {
	tests := []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "one"},
		{"one_two_three", "one-two-three"},
		{"ice_9", "ice-9"},

		// Examples with abbreviations
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

// TestSnakeToKebabExample tests SnakeToKebab function.
func TestSnakeToKebabExample(t *testing.T) {
	cases := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			name:  "Valid Snake Case",
			input: "hello_world",
			want:  "hello-world",
		},
		{
			name:    "Invalid Snake Case - CamelCase",
			input:   "HelloWorld",
			wantErr: true,
		},
		{
			name:    "Invalid Snake Case - camelCase",
			input:   "helloWorld",
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SnakeToKebab(tc.input)
			if tc.wantErr && err == nil {
				t.Errorf("SnakeToKebab(%q) expected error, got nil",
					tc.input)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("SnakeToKebab(%q) unexpected error: %v",
					tc.input, err)
			}
			if got != tc.want {
				t.Errorf("SnakeToKebab(%q) = %q, want %q",
					tc.input, got, tc.want)
			}
		})
	}
}

// TestSnakeToKebabError tests SnakeToKebab function with wrong value.
func TestSnakeToKebabError(t *testing.T) {
	notSnake := "oneTwoThree"

	_, err := SnakeToKebab(notSnake)
	if err == nil {
		t.Error("not snake to kebab")
	}
}

// TestSnakeToPascal tests SnakeToPascal function.
func TestSnakeToPascal(t *testing.T) {
	tests := []struct {
		value  string
		result string
	}{
		// Simple examples
		{"one", "One"},
		{"one_two_three", "OneTwoThree"},
		{"ice_9", "Ice9"},

		// Examples with abbreviations
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

// TestSnakeToPascalExample tests SnakeToPascal function.
func TestSnakeToPascalExample(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		want      string
		wantError bool
	}{
		{
			name:      "Valid snake_case string",
			input:     "hello_world",
			want:      "HelloWorld",
			wantError: false,
		},
		{
			name:      "Invalid non-snake_case string with uppercase",
			input:     "HelloWorld",
			want:      "",
			wantError: true,
		},
		{
			name:      "Invalid non-snake_case string with hyphen",
			input:     "hello-world",
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SnakeToPascal(tt.input)

			if (err != nil) != tt.wantError {
				t.Errorf("SnakeToPascal() error = %v, wantError %v",
					err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("SnakeToPascal() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestSnakeToPascalError tests SnakeToPascal function with wrong value.
func TestSnakeToPascalError(t *testing.T) {
	notSnake := "oneTwoThree"

	_, err := SnakeToPascal(notSnake)
	if err == nil {
		t.Error("not snake to pascal")
	}
}
