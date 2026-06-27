package scs

import "testing"

var benchInputs = []string{
	"helloWorld", "HTTPServerID", "user_id_value", "SomeLongPascalCaseName",
	"already_snake_case_string", "web2print_oauth2_utf8",
}

func benchConvert(b *testing.B, fn func(string) string) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(benchInputs[i%len(benchInputs)])
	}
}

func BenchmarkToCamel(b *testing.B)          { benchConvert(b, ToCamel) }
func BenchmarkToPascal(b *testing.B)         { benchConvert(b, ToPascal) }
func BenchmarkToSnake(b *testing.B)          { benchConvert(b, ToSnake) }
func BenchmarkToKebab(b *testing.B)          { benchConvert(b, ToKebab) }
func BenchmarkToScreamingSnake(b *testing.B) { benchConvert(b, ToScreamingSnake) }

func BenchmarkSplit(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split(benchInputs[i%len(benchInputs)])
	}
}

func BenchmarkDetect(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Detect(benchInputs[i%len(benchInputs)])
	}
}

func BenchmarkCaserAcronyms(b *testing.B) {
	c := New(WithAcronyms("ID", "URL", "HTTP", "API"))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.ToPascal(benchInputs[i%len(benchInputs)])
	}
}
