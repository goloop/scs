[//]: # (!!!Don't modify the README.md, use `make readme` to generate it!!!)


[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/scs)](https://goreportcard.com/report/github.com/goloop/scs) [![License](https://img.shields.io/badge/license-BSD-blue)](https://github.com/goloop/scs/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/scs)

*Version: 0.0.1*


# scs

The scs (String Case Style) module implements methods for converting string case
to various case styles: camelCase, kebab-case, PascalCase and snake_case.

## Installation

To install this module use `go get` as:

    $ go get -u github.com/goloop/scs

## Quick Start

To use this module import it as:

    package main

    import (
        "github.com/goloop/scs"
    )

    func main() {
        camel := scs.New(scs.Camel)
        camel.Eat("Hello World") // helloWorld
        camel.Value() // helloWorld

        kebab := camel.ToKebab()
        kebab.Value() // hello-world

        pascal := kebab.ToPascal()
        pascal.Value() // HelloWorld

        snake := pascal.ToSnake()
        snake.Value() // hello_world

        // Acronyms
        camel.Eat("Http to https")  // httpToHTTPS
        kebab.Eat("Http to https")  // http-to-https
        pascal.Eat("Http to https") // HTTPToHTTPS
        snake.Eat("Http to https")  // http_to_https
    }


## Usage

#### func  CamelToKebab

    func CamelToKebab(camel string) string

CamelToKebab converts a camelCase-style string to kebab-case. The conversion
will be invalid if the input string is not camelCase style.

#### func  CamelToPascal

    func CamelToPascal(camel string) string

CamelToPascal converts a camelCase-style string to PascalCase. The conversion
will be invalid if the input string is not camelCase style.

#### func  CamelToSnake

    func CamelToSnake(camel string) string

CamelToSnake converts a camelCase-style string to snake_case. The conversion
will be invalid if the input string is not camelCase style.

#### func  KebabToCamel

    func KebabToCamel(kebab string) string

KebabToCamel converts a kebab-case-style string to camelCase. The conversion
will be invalid if the input string is not kebab-case style.

#### func  KebabToPascal

    func KebabToPascal(kebab string) string

KebabToPascal converts a kebab-case-style string to PascalCase. The conversion
will be invalid if the input string is not kebab-case style.

#### func  KebabToSnake

    func KebabToSnake(kebab string) string

KebabToSnake converts a kebab-case-style string to snake_case. The conversion
will be invalid if the input string is not kebab-case style.

#### func  PascalToCamel

    func PascalToCamel(pascal string) string

PascalToCamel converts a PascalCase-style string to camelCase. The conversion
will be invalid if the input string is not PascalCase style.

#### func  PascalToKebab

    func PascalToKebab(pascal string) string

PascalToKebab converts a PascalCase-style string to kebab-case. The conversion
will be invalid if the input string is not PascalCase style.

#### func  PascalToSnake

    func PascalToSnake(pascal string) string

PascalToSnake converts a PascalCase-style string to snake_case. The conversion
will be invalid if the input string is not PascalCase style.

#### func  SnakeToCamel

    func SnakeToCamel(snake string) string

SnakeToCamel converts a snake_case-style string to camelCase. The conversion
will be invalid if the input string is not snake_case style.

#### func  SnakeToKebab

    func SnakeToKebab(snake string) string

SnakeToKebab converts a snake_case-style string to kebab-case. The conversion
will be invalid if the input string is not snake_case style.

#### func  SnakeToPascal

    func SnakeToPascal(snake string) string

SnakeToPascal converts a snake_case-style string to PascalCase. The conversion
will be invalid if the input string is not snake_case style.

#### func  ToCamel

    func ToCamel(s string) string

ToCamel converts a string to camelCase.

#### func  ToKebab

    func ToKebab(s string) string

ToKebab converts a string to kebab-case.

#### func  ToPascal

    func ToPascal(s string) string

ToPascal converts a string to PascalCase.

#### func  ToSnake

    func ToSnake(s string) string

ToSnake converts a string to snake_case.

#### type CaseStyle

    type CaseStyle uint8


CaseStyle is string case style type.

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


#### type Object

    type Object struct {
    }


Object is object of the string case style (SCS).

#### func  New

    func New(style CaseStyle, ss ...string) (*Object, error)

New returns a pointer to a string case style object. The style defines the
string case style. a string (or list of strings) to format.

#### func (*Object) Eat

    func (o *Object) Eat(s string) string

Eat converts a string to the specified style and stores it as an object value.

#### func (*Object) IsCamel

    func (o *Object) IsCamel() bool

IsCamel returns true if object contains camelCase value.

#### func (*Object) IsKebab

    func (o *Object) IsKebab() bool

IsKebab returns true if object contains kebab-case value.

#### func (*Object) IsPascal

    func (o *Object) IsPascal() bool

IsPascal returns true if object contains PascalCase value.

#### func (*Object) IsSnake

    func (o *Object) IsSnake() bool

IsSnake returns true if object contains snake-case value.

#### func (*Object) ToCamel

    func (o *Object) ToCamel() *Object

ToCamel converts an object to Camel Type Object and returns a pointer to it.

#### func (*Object) ToKebab

    func (o *Object) ToKebab() *Object

ToKebab converts an object to Kebab Type Object and returns a pointer to it.

#### func (*Object) ToPascal

    func (o *Object) ToPascal() *Object

ToPascal converts an object to Pascal Type Object and returns a pointer to it.

#### func (*Object) ToSnake

    func (o *Object) ToSnake() *Object

ToSnake converts an object to Snake Type Object and returns a pointer to it.

#### func (*Object) Value

    func (o *Object) Value() string

Value returns value of the object.
