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
		obj.do = ToCamel
	case Kebab:
		obj.do = ToKebab
	case Pascal:
		obj.do = ToPascal
	case Snake:
		obj.do = ToSnake
	default:
		return &obj, fmt.Errorf("incorrect case style")
	}

	obj.value = obj.do(strings.Join(ss, " "))
	return &obj, nil
}

// IsCamel returns true if object contains camelCase value.
func (o *Object) IsCamel() bool { return o.style == Camel }

// IsKebab returns true if object contains kebab-case value.
func (o *Object) IsKebab() bool { return o.style == Kebab }

// IsPascal returns true if object contains PascalCase value.
func (o *Object) IsPascal() bool { return o.style == Pascal }

// IsSnake returns true if object contains snake-case value.
func (o *Object) IsSnake() bool { return o.style == Snake }

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
func (o *Object) ToCamel() *Object {
	var obj = Object{style: Camel, do: ToCamel}

	switch o.style {
	case Camel:
		obj.value = o.value
	case Kebab:
		obj.value = KebabToCamel(o.value)
	case Pascal:
		obj.value = PascalToCamel(o.value)
	case Snake:
		obj.value = SnakeToCamel(o.value)
	}

	return &obj
}

// ToKebab converts an object to Kebab Type Object
// and returns a pointer to it.
func (o *Object) ToKebab() *Object {
	var obj = Object{style: Kebab, do: ToKebab}

	switch o.style {
	case Camel:
		obj.value = CamelToKebab(o.value)
	case Kebab:
		obj.value = o.value
	case Pascal:
		obj.value = PascalToKebab(o.value)
	case Snake:
		obj.value = SnakeToKebab(o.value)
	}

	return &obj
}

// ToPascal converts an object to Pascal Type Object
// and returns a pointer to it.
func (o *Object) ToPascal() *Object {
	var obj = Object{style: Pascal, do: ToPascal}

	switch o.style {
	case Camel:
		obj.value = CamelToPascal(o.value)
	case Kebab:
		obj.value = KebabToPascal(o.value)
	case Pascal:
		obj.value = o.value
	case Snake:
		obj.value = SnakeToPascal(o.value)
	}

	return &obj
}

// ToSnake converts an object to Snake Type Object
// and returns a pointer to it.
func (o *Object) ToSnake() *Object {
	var obj = Object{style: Snake, do: ToSnake}

	switch o.style {
	case Camel:
		obj.value = CamelToSnake(o.value)
	case Kebab:
		obj.value = KebabToSnake(o.value)
	case Pascal:
		obj.value = PascalToSnake(o.value)
	case Snake:
		obj.value = o.value
	}

	return &obj
}
