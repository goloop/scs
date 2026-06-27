package scs_test

import (
	"fmt"

	"github.com/goloop/scs/v2"
)

func ExampleToCamel() {
	fmt.Println(scs.ToCamel("hello_world"))
	fmt.Println(scs.ToCamel("HelloWorld"))
	fmt.Println(scs.ToCamel("HTTP server"))
	// Output:
	// helloWorld
	// helloWorld
	// httpServer
}

func ExampleToPascal() {
	fmt.Println(scs.ToPascal("hello-world"))
	fmt.Println(scs.ToPascal("userID"))
	// Output:
	// HelloWorld
	// UserId
}

func ExampleToSnake() {
	fmt.Println(scs.ToSnake("helloWorld"))
	fmt.Println(scs.ToSnake("HTTPServer"))
	// Output:
	// hello_world
	// http_server
}

func ExampleToKebab() {
	fmt.Println(scs.ToKebab("helloWorld"))
	// Output: hello-world
}

func ExampleToScreamingSnake() {
	fmt.Println(scs.ToScreamingSnake("userID"))
	// Output: USER_ID
}

func ExampleToDot() {
	fmt.Println(scs.ToDot("HelloWorld"))
	// Output: hello.world
}

func ExampleToTitle() {
	fmt.Println(scs.ToTitle("hello_world"))
	// Output: Hello World
}

func ExampleConvert() {
	for _, style := range []scs.Style{scs.Snake, scs.Kebab, scs.Camel} {
		fmt.Printf("%-6s %s\n", style, scs.Convert(style, "HTTPServerID"))
	}
	// Output:
	// snake  http_server_id
	// kebab  http-server-id
	// camel  httpServerId
}

func ExampleDetect() {
	for _, s := range []string{"user_id", "userId", "API", "api"} {
		style, ok := scs.Detect(s)
		fmt.Printf("%q -> %v (%v)\n", s, style, ok)
	}
	// Output:
	// "user_id" -> snake (true)
	// "userId" -> camel (true)
	// "API" -> screaming_snake (true)
	// "api" -> unknown (false)
}

func ExampleSplit() {
	fmt.Println(scs.Split("HTTPServerID"))
	fmt.Println(scs.Split("web2print"))
	// Output:
	// [http server id]
	// [web2print]
}

func ExampleWords() {
	for w := range scs.Words("parseJSONResponse") {
		fmt.Println(w)
	}
	// Output:
	// parse
	// json
	// response
}

func ExampleNew() {
	// Opt in to Go-style initialisms.
	c := scs.New(scs.WithAcronyms("ID", "URL", "HTTP"))
	fmt.Println(c.ToPascal("user_id"))
	fmt.Println(c.ToCamel("http_url_builder"))
	// Output:
	// UserID
	// httpURLBuilder
}
