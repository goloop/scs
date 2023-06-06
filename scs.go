package scs

import (
	"fmt"
	"strings"
)

const (
	// Camel is constant that characterizes string case style as camelCase.
	Camel CaseStyle = 1 << iota

	// Kebab is constant that characterizes string case style as kebab-case.
	Kebab

	// Pascal is constant that characterizes string case style as PascalCase.
	Pascal

	// Snake is constant that characterizes string case style as snake_case.
	Snake
)

// CaseStyle is string case style type.
type CaseStyle uint8

// StringCaseStyle is object of the string case style (SCS).
// It can be created correctly through the New function only.
type StringCaseStyle struct {
	do      func(string) string // convert raw string to the case-styling value
	style   CaseStyle           // is the flag of the string case style
	value   string              // value in case-style format
	isValid bool                // true if the object was created correctly
}

// New returns a pointer to a string case style object. The style defines
// the string case style. a string (or list of strings) to format.
//
// This function creates a new instance of the StringCaseStyle struct, which
// represents a specific string case style. It takes a CaseStyle parameter
// to determine the desired case style (Camel, Kebab, Pascal, Snake), and
// one or more string values to be formatted.
//
// The function initializes the `do` field of the StringCaseStyle struct with
// the appropriate formatting function based on the specified case style.
// It then joins the input strings with a space separator and applies the
// formatting function to the resulting string. The formatted string is stored
// in the `value` field of the StringCaseStyle struct.
//
// If an incorrect case style is provided, the function returns an error along
// with a nil pointer to the StringCaseStyle.
//
// Example usage:
//
//	style, err := scs.New(scs.Camel, "hello", "world")
func New(style CaseStyle, value ...string) (*StringCaseStyle, error) {
	var do func(string) string

	switch style {
	case Camel:
		do = StrToCamel
	case Kebab:
		do = StrToKebab
	case Pascal:
		do = StrToPascal
	case Snake:
		do = StrToSnake
	default:
		return &StringCaseStyle{do: func(s string) string { return s }},
			fmt.Errorf("incorrect case style")
	}

	return &StringCaseStyle{
		do:      do,
		style:   style,
		value:   do(strings.Join(value, " ")),
		isValid: true,
	}, nil
}

// IsValid returns true if the StringCaseStyle object is valid.
//
// This method is used to check the validity of a StringCaseStyle object.
// It returns true if the object is valid, indicating that the string case
// style and value have been set successfully.
//
// Example usage:
//
//	style, _ := New(Camel, "helloWorld")
//	isValid := style.IsValid()
//	// isValid: true
//
//	invalidStyle := &StringCaseStyle{isValid: false}
//	isValid := invalidStyle.IsValid()
//	// isValid: false
func (o *StringCaseStyle) IsValid() bool {
	return o.isValid
}

// IsCamel returns true if the StringCaseStyle object represents
// a camelCase value.
//
// This method checks if the StringCaseStyle object has a camelCase style.
// It returns true if the object represents a camelCase value, indicating
// that the original value or formatted value is in camelCase.
//
// Example usage:
//
//	style, _ := New(Camel, "helloWorld")
//	isCamel := style.IsCamel()
//	// isCamel: true
//
//	style, _ = New(Pascal, "HelloWorld")
//	isCamel = style.IsCamel()
//	// isCamel: false
func (o *StringCaseStyle) IsCamel() bool {
	return o.style == Camel
}

// IsKebab returns true if the StringCaseStyle object represents
// a kebab-case value.
//
// This method checks if the StringCaseStyle object has a kebab-case style.
// It returns true if the object represents a kebab-case value, indicating
// that the original value or formatted value is in kebab-case.
//
// Example usage:
//
//	style, _ := New(Kebab, "hello-world")
//	isKebab := style.IsKebab()
//	// isKebab: true
//
//	style, _ = New(Pascal, "HelloWorld")
//	isKebab = style.IsKebab()
//	// isKebab: false
func (o *StringCaseStyle) IsKebab() bool {
	return o.style == Kebab
}

// IsPascal returns true if the StringCaseStyle object represents
// a PascalCase value.
//
// This method checks if the StringCaseStyle object has a PascalCase style.
// It returns true if the object represents a PascalCase value, indicating
// that the original value or formatted value is in PascalCase.
//
// Example usage:
//
//	style, _ := New(Pascal, "HelloWorld")
//	isPascal := style.IsPascal()
//	// isPascal: true
//
//	style, _ = New(Snake, "hello_world")
//	isPascal = style.IsPascal()
//	// isPascal: false
func (o *StringCaseStyle) IsPascal() bool {
	return o.style == Pascal
}

// IsSnake returns true if the StringCaseStyle object represents
// a snake_case value.
//
// This method checks if the StringCaseStyle object has a snake_case style.
// It returns true if the object represents a snake_case value, indicating
// that the original value or formatted value is in snake_case.
//
// Example usage:
//
//	style, _ := New(Snake, "hello_world")
//	isSnake := style.IsSnake()
//	// isSnake: true
//
//	style, _ = New(Pascal, "HelloWorld")
//	isSnake = style.IsSnake()
//	// isSnake: false
func (o *StringCaseStyle) IsSnake() bool {
	return o.style == Snake
}

// Eat converts a string to the specified style and stores it
// as the object value.
//
// This method converts the input string to the style defined by the
// StringCaseStyle object and updates the object's value.
// The converted value is returned as the result.
//
// Example usage:
//
//	style, _ := New(Camel, "hello_world")
//	result := style.Eat("example_input")
//	// result: "exampleInput"
//	// style.Value(): "exampleInput"
//
//	style, _ = New(Kebab, "hello-world")
//	result = style.Eat("example_input")
//	// result: "example-input"
//	// style.Value(): "example-input"
func (o *StringCaseStyle) Eat(s string) string {
	o.value = o.do(s)
	return o.value
}

// Set sets a new value for the StringCaseStyle object.
//
// This method converts the input string to the style defined by the
// StringCaseStyle object and sets the converted value as the new value
// of the object. The updated object is returned for method chaining.
//
// Example usage:
//
//	style, _ := New(Camel, "hello_world")
//	result := style.Set("example_input")
//	// result.Value(): "exampleInput"
//
//	style, _ = New(Kebab, "hello-world")
//	result = style.Set("example_input")
//	// result.value: "example-input"
func (o *StringCaseStyle) Set(s string) *StringCaseStyle {
	o.value = o.do(s)
	return o
}

// Value returns the current value of the StringCaseStyle object.
//
// This method returns the current value of the StringCaseStyle object.
//
// Example usage:
//
//	style, _ := New(Camel, "Hello World")
//	result := style.Value()
//	// result: "Hello World"
//
//	style, _ = New(Kebab, "Show me")
//	result = style.Value()
//	// result: "Show me"
func (o *StringCaseStyle) Value() string {
	return o.value
}

// CopyToCamel converts an object to Camel Type StringCaseStyle
// and returns new pointer to it.
func (o *StringCaseStyle) CopyToCamel() (*StringCaseStyle, error) {
	var (
		value string
		err   error
	)

	switch o.style {
	case Camel:
		value = o.value
	case Kebab:
		value, err = KebabToCamel(o.value)
	case Pascal:
		value, err = PascalToCamel(o.value)
	case Snake:
		value, err = SnakeToCamel(o.value)
	}

	return &StringCaseStyle{
		do:      StrToCamel,
		style:   Camel,
		value:   value,
		isValid: err == nil,
	}, err
}

// ToCamel converts an object to Camel Type StringCaseStyle.
func (o *StringCaseStyle) ToCamel() error {
	obj, err := o.CopyToCamel()
	o.style = obj.style
	o.value = obj.value
	o.do = obj.do
	o.isValid = obj.isValid

	return err
}

// CopyToKebab converts an object to Kebab Type StringCaseStyle
// and returns new pointer to it.
func (o *StringCaseStyle) CopyToKebab() (*StringCaseStyle, error) {
	var (
		value string
		err   error
	)

	switch o.style {
	case Camel:
		value, err = CamelToKebab(o.value)
	case Kebab:
		value = o.value
	case Pascal:
		value, err = PascalToKebab(o.value)
	case Snake:
		value, err = SnakeToKebab(o.value)
	}

	return &StringCaseStyle{
		do:      StrToKebab,
		style:   Kebab,
		value:   value,
		isValid: err == nil,
	}, err
}

// ToKebab converts an object to Kebab Type StringCaseStyle.
func (o *StringCaseStyle) ToKebab() error {
	obj, err := o.CopyToKebab()
	o.style = obj.style
	o.value = obj.value
	o.do = obj.do
	o.isValid = obj.isValid

	return err
}

// CopyToPascal converts an object to Pascal Type StringCaseStyle
// and returns new pointer to it.
func (o *StringCaseStyle) CopyToPascal() (*StringCaseStyle, error) {
	var (
		value string
		err   error
	)

	switch o.style {
	case Camel:
		value, err = CamelToPascal(o.value)
	case Kebab:
		value, err = KebabToPascal(o.value)
	case Pascal:
		value = o.value
	case Snake:
		value, err = SnakeToPascal(o.value)
	}

	return &StringCaseStyle{
		do:      StrToPascal,
		style:   Pascal,
		value:   value,
		isValid: err == nil,
	}, err
}

// ToPascal converts an object to Pascal Type StringCaseStyle.
func (o *StringCaseStyle) ToPascal() error {
	obj, err := o.CopyToPascal()
	o.style = obj.style
	o.value = obj.value
	o.do = obj.do
	o.isValid = obj.isValid

	return err
}

// CopyToSnake converts an object to Snake Type StringCaseStyle
// and returns new pointer to it.
func (o *StringCaseStyle) CopyToSnake() (*StringCaseStyle, error) {
	var (
		value string
		err   error
	)

	switch o.style {
	case Camel:
		value, err = CamelToSnake(o.value)
	case Kebab:
		value, err = KebabToSnake(o.value)
	case Pascal:
		value, err = PascalToSnake(o.value)
	case Snake:
		value = o.value
	}

	return &StringCaseStyle{
		do:      StrToSnake,
		style:   Snake,
		value:   value,
		isValid: err == nil,
	}, err
}

// ToSnake converts an object to Snake Type StringCaseStyle.
func (o *StringCaseStyle) ToSnake() error {
	obj, err := o.CopyToSnake()
	o.style = obj.style
	o.value = obj.value
	o.do = obj.do
	o.isValid = obj.isValid

	return err
}
