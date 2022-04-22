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

// IsValid returns true if StringCaseStyle is valid.
func (o *StringCaseStyle) IsValid() bool {
	return o.isValid // o.IsCamel() || o.IsKebab() || o.IsPascal() || o.IsSnake()
}

// IsCamel returns true if object contains camelCase value.
func (o *StringCaseStyle) IsCamel() bool {
	return o.style == Camel
}

// IsKebab returns true if object contains kebab-case value.
func (o *StringCaseStyle) IsKebab() bool {
	return o.style == Kebab
}

// IsPascal returns true if object contains PascalCase value.
func (o *StringCaseStyle) IsPascal() bool {
	return o.style == Pascal
}

// IsSnake returns true if object contains snake-case value.
func (o *StringCaseStyle) IsSnake() bool {
	return o.style == Snake
}

// Eat converts a string to the specified style and stores
// it as an object value.
func (o *StringCaseStyle) Eat(s string) string {
	o.value = o.do(s)
	return o.value
}

// Set sets new value.
func (o *StringCaseStyle) Set(s string) *StringCaseStyle {
	o.value = o.do(s)
	return o
}

// Value returns value of the object.
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
