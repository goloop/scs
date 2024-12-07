package scs

import (
	"testing"
)

var (
	testStr           = "Hello World HTTP API"
	benchmarkResult   string
	benchmarkBool     bool
	benchmarkSCSObj   *StringCaseStyle
	benchmarkSCSError error
)

func BenchmarkStrToCamel(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = StrToCamel(testStr)
	}
	benchmarkResult = r
}

func BenchmarkStrToKebab(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = StrToKebab(testStr)
	}
	benchmarkResult = r
}

func BenchmarkStrToPascal(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = StrToPascal(testStr)
	}
	benchmarkResult = r
}

func BenchmarkStrToSnake(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = StrToSnake(testStr)
	}
	benchmarkResult = r
}

func BenchmarkCamelToKebab(b *testing.B) {
	input := "helloWorldHTTPAPI"
	var r string
	var err error
	for i := 0; i < b.N; i++ {
		r, err = CamelToKebab(input)
	}
	benchmarkResult = r
	benchmarkSCSError = err
}

func BenchmarkKebabToPascal(b *testing.B) {
	input := "hello-world-http-api"
	var r string
	var err error
	for i := 0; i < b.N; i++ {
		r, err = KebabToPascal(input)
	}
	benchmarkResult = r
	benchmarkSCSError = err
}

func BenchmarkPascalToSnake(b *testing.B) {
	input := "HelloWorldHTTPAPI"
	var r string
	var err error
	for i := 0; i < b.N; i++ {
		r, err = PascalToSnake(input)
	}
	benchmarkResult = r
	benchmarkSCSError = err
}

func BenchmarkSnakeToCamel(b *testing.B) {
	input := "hello_world_http_api"
	var r string
	var err error
	for i := 0; i < b.N; i++ {
		r, err = SnakeToCamel(input)
	}
	benchmarkResult = r
	benchmarkSCSError = err
}

func BenchmarkStrIsCamel(b *testing.B) {
	input := "helloWorldHTTPAPI"
	var r bool
	for i := 0; i < b.N; i++ {
		r = StrIsCamel(input)
	}
	benchmarkBool = r
}

func BenchmarkStrIsKebab(b *testing.B) {
	input := "hello-world-http-api"
	var r bool
	for i := 0; i < b.N; i++ {
		r = StrIsKebab(input)
	}
	benchmarkBool = r
}

func BenchmarkStrIsPascal(b *testing.B) {
	input := "HelloWorldHTTPAPI"
	var r bool
	for i := 0; i < b.N; i++ {
		r = StrIsPascal(input)
	}
	benchmarkBool = r
}

func BenchmarkStrIsSnake(b *testing.B) {
	input := "hello_world_http_api"
	var r bool
	for i := 0; i < b.N; i++ {
		r = StrIsSnake(input)
	}
	benchmarkBool = r
}

func BenchmarkNewStringCaseStyle(b *testing.B) {
	var obj *StringCaseStyle
	var err error
	for i := 0; i < b.N; i++ {
		obj, err = New(Camel, testStr)
	}
	benchmarkSCSObj = obj
	benchmarkSCSError = err
}

func BenchmarkStringCaseStyleConversions(b *testing.B) {
	obj, _ := New(Camel, testStr)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		obj.ToKebab()
		obj.ToPascal()
		obj.ToSnake()
		obj.ToCamel()
	}
	benchmarkSCSObj = obj
}
