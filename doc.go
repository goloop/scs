// Package scs (String Case Style) implements methods for converting string
// cases between different naming conventions: camelCase, kebab-case,
// PascalCase, and snake_case.
//
// # String Case Styles
//
// The package supports four main case styles:
//
//   - camelCase: Words are joined together, first word starts with lowercase,
//     subsequent words start with uppercase (e.g., "helloWorld")
//   - kebab-case: Words are lowercase and separated by hyphens (e.g., "hello-world")
//   - PascalCase: Words are joined together, each word starts with uppercase
//     (e.g., "HelloWorld")
//   - snake_case: Words are lowercase and separated by underscores
//     (e.g., "hello_world")
//
// # Usage
//
// The package provides two main ways to work with string case styles:
//
//  1. Direct conversion functions:
//     str := scs.StrToCamel("hello-world")    // returns "helloWorld"
//     str := scs.StrToKebab("HelloWorld")     // returns "hello-world"
//     str := scs.StrToPascal("hello_world")   // returns "HelloWorld"
//     str := scs.StrToSnake("helloWorld")     // returns "hello_world"
//
//  2. Object-oriented approach using StringCaseStyle:
//     style, _ := scs.New(scs.Camel, "hello-world")
//     style.Value()    // returns "helloWorld"
//     style.ToKebab()  // converts to kebab-case
//     style.Value()    // returns "hello-world"
//
// # Special Cases
//
// The package handles special cases like abbreviations and numbers:
//   - Abbreviations are preserved: "HTTP API" -> "HTTPApi" (PascalCase)
//   - Numbers are treated as word boundaries: "web2print" -> "web-2-print" (kebab-case)
//
// # Thread Safety
//
// All functions in this package are thread-safe and can be used concurrently.
// The StringCaseStyle object methods are not thread-safe and should be
// protected if used concurrently.
//
// # Performance
//
// The package uses optimized regular expressions and string operations
// to provide efficient case conversions. Benchmark tests are available
// in the scs_test.go file.
//
// # Error Handling
//
// Conversion functions that require specific input formats (like CamelToKebab)
// return errors if the input string doesn't match the expected format.
// The StringCaseStyle object provides methods to check if conversions
// were successful.
package scs
