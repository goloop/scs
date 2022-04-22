[//]: # (!!!Don't modify the README.md, use `make readme` to generate it!!!)


[![Go Report Card](https://goreportcard.com/badge/github.com/goloop/scs)](https://goreportcard.com/report/github.com/goloop/scs) [![License](https://img.shields.io/badge/license-BSD-blue)](https://github.com/goloop/scs/blob/master/LICENSE) [![License](https://img.shields.io/badge/godoc-YES-green)](https://godoc.org/github.com/goloop/scs)

*Version: v1.0.0*


# SCS

Package scs (String Case Style) module implements methods for converting string
case to various case styles: camelCase, kebab-case, PascalCase and snake_case.

## Installation

To install this module use `go get` as:

    $ go get -u github.com/goloop/scs

## Quick Start

To use this module import it as: `github.com/goloop/scs`

### Conversion functions

Example:

```go
    package main

    import "github.com/goloop/scs"

    func main() {
        var s string

        // Simple text
        s = "hello world"
        scs.StrToCamel(s)  // helloWorld
        scs.StrToPascal(s) // HelloWorld
        scs.StrToSnake(s)  // hello_world
        scs.StrToKebab(s)  // hello-world

        // Text with abbreviations
        s = "http to https"
        scs.StrToCamel(s)  // httpToHTTPS
        scs.StrToPascal(s) // HTTPToHTTPS
        scs.StrToSnake(s)  // http_to_https
        scs.StrToKebab(s)  // http-to-https

        // Converting
        s = "http to https"
        camel := scs.StrToCamel(s)            // httpToHTTPS
        pascal, _ := scs.CamelToPascal(camel) // HTTPToHTTPS <nil>
        kebab, _ := scs.PascalToKebab(pascal) // http-to-https <nil>
        snake, _ := scs.KebabToSnake(kebab)   // http_to_https <nil>

        // Use strings.ToUpper(snake) for convert to UPPER_SNAKE_CASE.

        scs.SnakeToPascal(snake) // HTTPToHTTPS <nil>
        scs.CamelToKebab(camel)  // http-to-https <nil>

        // Errors
        scs.CamelToSnake(kebab)  // value http-to-https isn't camelCase style
        scs.PascalToCamel(camel) // value httpToHTTPS isn't PascalCase style

        // Convert anything to anything correctly
        scs.ToCamel(snake)  // httpToHTTPS
        scs.ToPascal(kebab) // HTTPToHTTPS
        scs.ToSnake(s)      // http_to_https
    }
```

### Style objects

A safer way. Since each object knows what type it is and knows
which conversion rules to use. This removes the need to return
a second parameter as err when converting styles.

Example:

```go
    package main

    import "github.com/goloop/scs"

    func main() {
        var s string

        // Simple text
        s = "hello world"
        snake, _ := scs.New(scs.Snake) // scs.New(scs.Snake, s)

        snake.Eat(s)         // hello_world
        snake.Set(s).Value() // hello_world
        snake.IsCamel()      // false
        snake.IsSnake()      // true
        snake.Value()        // hello_world

        camel := snake.CopyToCamel()
        camel.IsSnake() // false
        camel.IsCamel() // true
        camel.Value()   // helloWorld

        // Text with abbreviations
        s = "http to https"
        pascal, _ := scs.New(scs.Pascal, s)
        pascal.Value() // HTTPToHTTPS

        kebab := pascal.CopyToKebab()
        kebab.Value() // http-to-https
    }
```

## Usage

#### func  CamelToKebab

    func CamelToKebab(camel string) (string, error)

CamelToKebab converts a camelCase-style string to kebab-case. The conversion
will be invalid if the input string is not camelCase style.

#### func  CamelToPascal

    func CamelToPascal(camel string) (string, error)

CamelToPascal converts a camelCase-style string to PascalCase. The conversion
will be invalid if the input string is not camelCase style.

#### func  CamelToSnake

    func CamelToSnake(camel string) (string, error)

CamelToSnake converts a camelCase-style string to snake_case. The conversion
will be invalid if the input string is not camelCase style.

#### func  KebabToCamel

    func KebabToCamel(kebab string) (string, error)

KebabToCamel converts a kebab-case-style string to camelCase. The conversion
will be invalid if the input string is not kebab-case style.

#### func  KebabToPascal

    func KebabToPascal(kebab string) (string, error)

KebabToPascal converts a kebab-case-style string to PascalCase. The conversion
will be invalid if the input string is not kebab-case style.

#### func  KebabToSnake

    func KebabToSnake(kebab string) (string, error)

KebabToSnake converts a kebab-case-style string to snake_case. The conversion
will be invalid if the input string is not kebab-case style.

#### func  PascalToCamel

    func PascalToCamel(pascal string) (string, error)

PascalToCamel converts a PascalCase-style string to camelCase. The conversion
will be invalid if the input string is not PascalCase style.

#### func  PascalToKebab

    func PascalToKebab(pascal string) (string, error)

PascalToKebab converts a PascalCase-style string to kebab-case. The conversion
will be invalid if the input string is not PascalCase style.

#### func  PascalToSnake

    func PascalToSnake(pascal string) (string, error)

PascalToSnake converts a PascalCase-style string to snake_case. The conversion
will be invalid if the input string is not PascalCase style.

#### func  SnakeToCamel

    func SnakeToCamel(snake string) (string, error)

SnakeToCamel converts a snake_case-style string to camelCase. The conversion
will be invalid if the input string is not snake_case style.

#### func  SnakeToKebab

    func SnakeToKebab(snake string) (string, error)

SnakeToKebab converts a snake_case-style string to kebab-case. The conversion
will be invalid if the input string is not snake_case style.

#### func  SnakeToPascal

    func SnakeToPascal(snake string) (string, error)

SnakeToPascal converts a snake_case-style string to PascalCase. The conversion
will be invalid if the input string is not snake_case style.

#### func  StrIsCamel

    func StrIsCamel(s string) bool

StrIsCamel returns true if string is camelCase.

#### func  StrIsKebab

    func StrIsKebab(s string) bool

StrIsKebab returns true if string is kebab-case.

#### func  StrIsPascal

    func StrIsPascal(s string) bool

StrIsPascal returns true if string is PascalCase.

#### func  StrIsSnake

    func StrIsSnake(s string) bool

StrIsSnake returns true if string is snake_case.

#### func  StrToCamel

    func StrToCamel(s string) string

StrToCamel converts a string to camelCase.

#### func  StrToKebab

    func StrToKebab(s string) string

StrToKebab converts a string to kebab-case.

#### func  StrToPascal

    func StrToPascal(s string) string

StrToPascal converts a string to PascalCase.

#### func  StrToSnake

    func StrToSnake(s string) string

StrToSnake converts a string to snake_case.

#### func  ToCamel

    func ToCamel(s string) string

ToCamel converts a string to camelCase. Unlike the StrToCamel function, if the
source string already has a certain format, it will be correctly converted to
camelCase.

#### func  ToKebab

    func ToKebab(s string) string

ToKebab converts a string to kebab-case. Unlike the StrToKebab function, if the
source string already has a certain format, it will be correctly converted to
kebab-case.

#### func  ToPascal

    func ToPascal(s string) string

ToPascal converts a string to PascalCase. Unlike the StrToPascal function, if
the source string already has a certain format, it will be correctly converted
to PascalCase.

#### func  ToSnake

    func ToSnake(s string) string

ToSnake converts a string to snake_case. Unlike the StrToSnake function, if the
source string already has a certain format, it will be correctly converted to
snake_case.

#### func  Version

    func Version() string

Version returns the version of the module.

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


#### type StringCaseStyle

    type StringCaseStyle struct {}


StringCaseStyle is object of the string case style (SCS). It can be created
correctly through the New function only.

#### func  New

    func New(style CaseStyle, value ...string) (*StringCaseStyle, error)

New returns a pointer to a string case style object. The style defines the
string case style. a string (or list of strings) to format.

#### func (*StringCaseStyle) CopyToCamel

    func (o *StringCaseStyle) CopyToCamel() (*StringCaseStyle, error)

CopyToCamel converts an object to Camel Type StringCaseStyle and returns new
pointer to it.

#### func (*StringCaseStyle) CopyToKebab

    func (o *StringCaseStyle) CopyToKebab() (*StringCaseStyle, error)

CopyToKebab converts an object to Kebab Type StringCaseStyle and returns new
pointer to it.

#### func (*StringCaseStyle) CopyToPascal

    func (o *StringCaseStyle) CopyToPascal() (*StringCaseStyle, error)

CopyToPascal converts an object to Pascal Type StringCaseStyle and returns new
pointer to it.

#### func (*StringCaseStyle) CopyToSnake

    func (o *StringCaseStyle) CopyToSnake() (*StringCaseStyle, error)

CopyToSnake converts an object to Snake Type StringCaseStyle and returns new
pointer to it.

#### func (*StringCaseStyle) Eat

    func (o *StringCaseStyle) Eat(s string) string

Eat converts a string to the specified style and stores it as an object value.

#### func (*StringCaseStyle) IsCamel

    func (o *StringCaseStyle) IsCamel() bool

IsCamel returns true if object contains camelCase value.

#### func (*StringCaseStyle) IsKebab

    func (o *StringCaseStyle) IsKebab() bool

IsKebab returns true if object contains kebab-case value.

#### func (*StringCaseStyle) IsPascal

    func (o *StringCaseStyle) IsPascal() bool

IsPascal returns true if object contains PascalCase value.

#### func (*StringCaseStyle) IsSnake

    func (o *StringCaseStyle) IsSnake() bool

IsSnake returns true if object contains snake-case value.

#### func (*StringCaseStyle) IsValid

    func (o *StringCaseStyle) IsValid() bool

IsValid returns true if StringCaseStyle is valid.

#### func (*StringCaseStyle) Set

    func (o *StringCaseStyle) Set(s string) *StringCaseStyle

Set sets new value.

#### func (*StringCaseStyle) ToCamel

    func (o *StringCaseStyle) ToCamel() error

ToCamel converts an object to Camel Type StringCaseStyle.

#### func (*StringCaseStyle) ToKebab

    func (o *StringCaseStyle) ToKebab() error

ToKebab converts an object to Kebab Type StringCaseStyle.

#### func (*StringCaseStyle) ToPascal

    func (o *StringCaseStyle) ToPascal() error

ToPascal converts an object to Pascal Type StringCaseStyle.

#### func (*StringCaseStyle) ToSnake

    func (o *StringCaseStyle) ToSnake() error

ToSnake converts an object to Snake Type StringCaseStyle.

#### func (*StringCaseStyle) Value

    func (o *StringCaseStyle) Value() string

Value returns value of the object.
