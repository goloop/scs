package scs

import "testing"

// convCase is one input and its expected rendering in every style.
type convCase struct {
	in    string
	camel string
	pasc  string
	snake string
	kebab string
	scr   string
	dot   string
	title string
	sent  string
}

var convCases = []convCase{
	{
		in: "helloWorld", camel: "helloWorld", pasc: "HelloWorld",
		snake: "hello_world", kebab: "hello-world", scr: "HELLO_WORLD",
		dot: "hello.world", title: "Hello World", sent: "Hello world",
	},
	{
		in: "HelloWorld", camel: "helloWorld", pasc: "HelloWorld",
		snake: "hello_world", kebab: "hello-world", scr: "HELLO_WORLD",
		dot: "hello.world", title: "Hello World", sent: "Hello world",
	},
	{
		in: "hello_world", camel: "helloWorld", pasc: "HelloWorld",
		snake: "hello_world", kebab: "hello-world", scr: "HELLO_WORLD",
		dot: "hello.world", title: "Hello World", sent: "Hello world",
	},
	{
		in: "hello-world", camel: "helloWorld", pasc: "HelloWorld",
		snake: "hello_world", kebab: "hello-world", scr: "HELLO_WORLD",
		dot: "hello.world", title: "Hello World", sent: "Hello world",
	},
	{
		in: "Hello World", camel: "helloWorld", pasc: "HelloWorld",
		snake: "hello_world", kebab: "hello-world", scr: "HELLO_WORLD",
		dot: "hello.world", title: "Hello World", sent: "Hello world",
	},
	// Acronym handling without an initialism set: everything Title-cases.
	{
		in: "HTTP API", camel: "httpApi", pasc: "HttpApi",
		snake: "http_api", kebab: "http-api", scr: "HTTP_API",
		dot: "http.api", title: "Http Api", sent: "Http api",
	},
	{
		in: "HTTPServerID", camel: "httpServerId", pasc: "HttpServerId",
		snake: "http_server_id", kebab: "http-server-id", scr: "HTTP_SERVER_ID",
		dot: "http.server.id", title: "Http Server Id", sent: "Http server id",
	},
	{
		in: "userID", camel: "userId", pasc: "UserId",
		snake: "user_id", kebab: "user-id", scr: "USER_ID",
		dot: "user.id", title: "User Id", sent: "User id",
	},
	// Digits glued.
	{
		in: "web2print", camel: "web2print", pasc: "Web2print",
		snake: "web2print", kebab: "web2print", scr: "WEB2PRINT",
		dot: "web2print", title: "Web2print", sent: "Web2print",
	},
	{
		in: "version2Final", camel: "version2Final", pasc: "Version2Final",
		snake: "version2_final", kebab: "version2-final", scr: "VERSION2_FINAL",
		dot: "version2.final", title: "Version2 Final", sent: "Version2 final",
	},
	// Explicit separators split numbers into their own word.
	{
		in: "web 2 print", camel: "web2Print", pasc: "Web2Print",
		snake: "web_2_print", kebab: "web-2-print", scr: "WEB_2_PRINT",
		dot: "web.2.print", title: "Web 2 Print", sent: "Web 2 print",
	},
	// Empty input is total: empty out everywhere.
	{
		in: "", camel: "", pasc: "", snake: "", kebab: "", scr: "", dot: "", title: "", sent: "",
	},
	{
		in: "   ", camel: "", pasc: "", snake: "", kebab: "", scr: "", dot: "", title: "", sent: "",
	},
	// Unicode words.
	{
		in: "привіт світ", camel: "привітСвіт", pasc: "ПривітСвіт",
		snake: "привіт_світ", kebab: "привіт-світ", scr: "ПРИВІТ_СВІТ",
		dot: "привіт.світ", title: "Привіт Світ", sent: "Привіт світ",
	},
}

func TestConverters(t *testing.T) {
	for _, c := range convCases {
		t.Run(c.in, func(t *testing.T) {
			checkEq(t, "ToCamel", c.in, ToCamel(c.in), c.camel)
			checkEq(t, "ToPascal", c.in, ToPascal(c.in), c.pasc)
			checkEq(t, "ToSnake", c.in, ToSnake(c.in), c.snake)
			checkEq(t, "ToKebab", c.in, ToKebab(c.in), c.kebab)
			checkEq(t, "ToScreamingSnake", c.in, ToScreamingSnake(c.in), c.scr)
			checkEq(t, "ToDot", c.in, ToDot(c.in), c.dot)
			checkEq(t, "ToTitle", c.in, ToTitle(c.in), c.title)
			checkEq(t, "ToSentence", c.in, ToSentence(c.in), c.sent)
		})
	}
}

func checkEq(t *testing.T, fn, in, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s(%q) = %q, want %q", fn, in, got, want)
	}
}

// TestConvertMatchesNamedFunctions guarantees Convert and the named helpers
// agree for every valid style — there is exactly one rendering per style.
func TestConvertMatchesNamedFunctions(t *testing.T) {
	pairs := []struct {
		style Style
		fn    func(string) string
	}{
		{Camel, ToCamel},
		{Pascal, ToPascal},
		{Snake, ToSnake},
		{Kebab, ToKebab},
		{ScreamingSnake, ToScreamingSnake},
		{Dot, ToDot},
		{Title, ToTitle},
		{Sentence, ToSentence},
	}
	for _, c := range convCases {
		for _, p := range pairs {
			if got, want := Convert(p.style, c.in), p.fn(c.in); got != want {
				t.Errorf("Convert(%v, %q) = %q, named fn = %q", p.style, c.in, got, want)
			}
		}
	}
}

// TestConvertUnknownStyleIsTotal: an invalid style must not panic and must
// pass the input through unchanged.
func TestConvertUnknownStyleIsTotal(t *testing.T) {
	for _, s := range []string{"", "userID", "hello_world"} {
		if got := Convert(Unknown, s); got != s {
			t.Errorf("Convert(Unknown, %q) = %q, want unchanged", s, got)
		}
		if got := Convert(Style(250), s); got != s {
			t.Errorf("Convert(invalid, %q) = %q, want unchanged", s, got)
		}
	}
}

// TestIPv6IsADocumentedEdge pins the known limitation: a dictionary-free
// tokenizer cannot tell "IP"+"v6" apart from the "acronym + lowercase word"
// pattern it splits everywhere else. This test exists so the behavior is a
// conscious, observed contract rather than a silent surprise.
func TestIPv6IsADocumentedEdge(t *testing.T) {
	if got := ToSnake("IPv6Address"); got != "i_pv6_address" {
		t.Errorf("ToSnake(IPv6Address) = %q; documented edge expects %q", got, "i_pv6_address")
	}
}
