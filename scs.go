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

// Object is object of the string case style (SCS).
type Object struct {
	// the style is flag of the string case style.
	style CaseStyle

	// the value string value in case-style format
	// that set in the style parameter.
	value string

	// the do is method to convert the raw string to the specified style.
	do func(string) string
}

// New returns a pointer to a string case style object. The style defines
// the string case style. a string (or list of strings) to format.
func New(style CaseStyle, ss ...string) (*Object, error) {
	var obj = Object{style: style}

	switch style {
	case Camel:
		obj.do = StrToCamel
	case Kebab:
		obj.do = StrToKebab
	case Pascal:
		obj.do = StrToPascal
	case Snake:
		obj.do = StrToSnake
	default:
		obj.do = func(s string) string { return s }
		return &obj, fmt.Errorf("incorrect case style")
	}

	obj.value = obj.do(strings.Join(ss, " "))
	return &obj, nil
}

// IsValid returns true if Object is valid.
func (o *Object) IsValid() bool {
	return o.IsCamel() || o.IsKebab() || o.IsPascal() || o.IsSnake()
}

// IsCamel returns true if object contains camelCase value.
func (o *Object) IsCamel() bool {
	return o.style == Camel
}

// IsKebab returns true if object contains kebab-case value.
func (o *Object) IsKebab() bool {
	return o.style == Kebab
}

// IsPascal returns true if object contains PascalCase value.
func (o *Object) IsPascal() bool {
	return o.style == Pascal
}

// IsSnake returns true if object contains snake-case value.
func (o *Object) IsSnake() bool {
	return o.style == Snake
}

// Eat converts a string to the specified style and stores
// it as an object value.
func (o *Object) Eat(s string) string {
	o.value = o.do(s)
	return o.value
}

// Value returns value of the object.
func (o *Object) Value() string {
	return o.value
}

// ToCamel converts an object to Camel Type Object
// and returns a pointer to it.
func (o *Object) ToCamel() (*Object, error) {
	var (
		err error
		obj = Object{style: Camel, do: StrToCamel}
	)

	switch o.style {
	case Camel:
		obj.value = o.value
	case Kebab:
		obj.value, err = KebabToCamel(o.value)
	case Pascal:
		obj.value, err = PascalToCamel(o.value)
	case Snake:
		obj.value, err = SnakeToCamel(o.value)
	}

	return &obj, err
}

// ToKebab converts an object to Kebab Type Object
// and returns a pointer to it.
func (o *Object) ToKebab() (*Object, error) {
	var (
		err error
		obj = Object{style: Kebab, do: StrToKebab}
	)

	switch o.style {
	case Camel:
		obj.value, err = CamelToKebab(o.value)
	case Kebab:
		obj.value = o.value
	case Pascal:
		obj.value, err = PascalToKebab(o.value)
	case Snake:
		obj.value, err = SnakeToKebab(o.value)
	}

	return &obj, err
}

// ToPascal converts an object to Pascal Type Object
// and returns a pointer to it.
func (o *Object) ToPascal() (*Object, error) {
	var (
		err error
		obj = Object{style: Pascal, do: StrToPascal}
	)

	switch o.style {
	case Camel:
		obj.value, err = CamelToPascal(o.value)
	case Kebab:
		obj.value, err = KebabToPascal(o.value)
	case Pascal:
		obj.value = o.value
	case Snake:
		obj.value, err = SnakeToPascal(o.value)
	}

	return &obj, err
}

// ToSnake converts an object to Snake Type Object
// and returns a pointer to it.
func (o *Object) ToSnake() (*Object, error) {
	var (
		err error
		obj = Object{style: Snake, do: StrToSnake}
	)

	switch o.style {
	case Camel:
		obj.value, err = CamelToSnake(o.value)
	case Kebab:
		obj.value, err = KebabToSnake(o.value)
	case Pascal:
		obj.value, err = PascalToSnake(o.value)
	case Snake:
		obj.value = o.value
	}

	return &obj, err
}
