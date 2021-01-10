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
        scs.StrToCamel(s)  // helloWorld
        scs.StrToPascal(s) // HelloWorld
        scs.StrToSnake(s)  // hello_world
        scs.StrToKebab(s)  // hello-world

        // Text with acronyms
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

### Style objects

A safer way. Since each object knows what type it is and knows
which conversion rules to use. This removes the need to return
a second parameter as err when converting styles.

Example:

    package main

    import "github.com/goloop/scs"

    func main() {
        var s string

        // Simple text
        s = "hello world"
        snake, _ := scs.New(scs.Snake) // scs.New(scs.Snake, s)

        snake.Eat(s)    // hello_world
        snake.IsCamel() // false
        snake.IsSnake() // true
        snake.Value()   // hello_world

        camel := snake.ToCamel()
        camel.IsSnake() // false
        camel.IsCamel() // true
        camel.Value()   // helloWorld

        // Text with acronyms
        s = "http to https"
        pascal, _ := scs.New(scs.Pascal, s)
        pascal.Value() // HTTPToHTTPS

        kebab := pascal.ToKebab()
        kebab.Value() // http-to-https
    }

*/
package scs
