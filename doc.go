/*
# scs

The scs (String Case Style) module implements methods for converting string
case to various case styles: camelCase, kebab-case, PascalCase and snake_case.

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
*/
package scs
