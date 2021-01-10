/*
# scs

The scs (String Case Style) module implements methods for converting string
case to various case styles: camelCase, kebab-case, PascalCase and snake_case.

## Installation

To install this module use `go get` as:

    $ go get -u github.com/goloop/scs

## Quick Start

To use this module import it as: `github.com/goloop/scs`

### Conversion functions

Example:

    package main

    import "github.com/goloop/scs"

    func main() {
        var s string

        // Simple text
        s = "hello world"
        scs.StrToCamel(s)  // helloWorld <nil>
        scs.StrToPascal(s) // HelloWorld <nil>
        scs.StrToSnake(s)  // hello_world <nil>
        scs.StrToKebab(s)  // hello-world <nil>

        // Text with acronyms
        s = "http to https"
        scs.StrToCamel(s)  // httpToHTTPS <nil>
        scs.StrToPascal(s) // HTTPToHTTPS <nil>
        scs.StrToSnake(s)  // http_to_https <nil>
        scs.StrToKebab(s)  // http-to-https <nil>

        // Converting
        s = "http to https"
        camel, _ := scs.StrToCamel(s)         // httpToHTTPS <nil>
        pascal, _ := scs.CamelToPascal(camel) // HTTPToHTTPS <nil>
        kebab, _ := scs.PascalToKebab(pascal) // http-to-https <nil>
        snake, _ := scs.KebabToSnake(kebab)   // http_to_https <nil>

        scs.SnakeToPascal(snake) // HTTPToHTTPS <nil>
        scs.CamelToKebab(camel)  // http-to-https <nil>

        // Errors
        scs.CamelToSnake(kebab)  // value http-to-https isn't camelCase style
        scs.PascalToCamel(camel) // value httpToHTTPS isn't PascalCase style
    }

### Style objects

Example:

    package main

    import "github.com/goloop/scs"

    func main() {
        var s string

        // Simple text
        s = "hello world"
        snake, _ := scs.New(scs.Snake) // scs.New(scs.Snake, s)

        snake.Eat(s)    // hello_world <nil>
        snake.IsCamel() // false
        snake.IsSnake() // true
        snake.Value()   // hello_world

        camel, _ := snake.ToCamel()
        camel.IsSnake() // false
        camel.IsCamel() // true
        camel.Value()   // helloWorld

        // Text with acronyms
        s = "http to https"
        pascal, _ := scs.New(scs.Pascal, s)
        pascal.Value() // HTTPToHTTPS

        kebab, _ := pascal.ToKebab()
        kebab.Value() // http-to-https
    }

*/
package scs
