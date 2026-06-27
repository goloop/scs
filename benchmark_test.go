package scs

import "testing"

var benchInputs = []string{
	"helloWorld", "HTTPServerID", "user_id_value", "SomeLongPascalCaseName",
	"already_snake_case_string", "web2print_oauth2_utf8",
}

// benchUnicodeInputs exercise the non-ASCII path of the tokenizer and the
// renderers (Cyrillic, German, Greek), which take the slower rune-aware code.
var benchUnicodeInputs = []string{
	"привітСвіт", "файл_конфігурації_сервера", "straßeTestWert",
	"ΕλληνικάKeimena", "日本語Text", "ОбробникHTTPЗапитів",
}

func benchOver(b *testing.B, inputs []string, fn func(string) string) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fn(inputs[i%len(inputs)])
	}
}

func benchConvert(b *testing.B, fn func(string) string) { benchOver(b, benchInputs, fn) }

func BenchmarkToCamel(b *testing.B)          { benchConvert(b, ToCamel) }
func BenchmarkToPascal(b *testing.B)         { benchConvert(b, ToPascal) }
func BenchmarkToSnake(b *testing.B)          { benchConvert(b, ToSnake) }
func BenchmarkToKebab(b *testing.B)          { benchConvert(b, ToKebab) }
func BenchmarkToScreamingSnake(b *testing.B) { benchConvert(b, ToScreamingSnake) }
func BenchmarkToSentence(b *testing.B)       { benchConvert(b, ToSentence) }

func BenchmarkToCamelUnicode(b *testing.B) { benchOver(b, benchUnicodeInputs, ToCamel) }
func BenchmarkToSnakeUnicode(b *testing.B) { benchOver(b, benchUnicodeInputs, ToSnake) }

func BenchmarkSplitUnicode(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Split(benchUnicodeInputs[i%len(benchUnicodeInputs)])
	}
}

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
