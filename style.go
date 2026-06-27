package scs

// Style enumerates the naming conventions the package can render to.
//
// The zero value is Unknown and is never produced by a converter; it is
// returned by Detect when the input does not map unambiguously to a single
// style. Use Valid to test whether a Style names a real convention.
type Style uint8

const (
	// Unknown is the zero value: not a real style. Detect returns it for
	// inputs that match no style or match several at once.
	Unknown Style = iota

	// Camel is camelCase: words joined, first word lowercase, every
	// subsequent word capitalized (e.g. "helloWorld").
	Camel

	// Pascal is PascalCase: words joined, every word capitalized including
	// the first (e.g. "HelloWorld").
	Pascal

	// Snake is snake_case: lowercase words joined by underscores
	// (e.g. "hello_world").
	Snake

	// Kebab is kebab-case: lowercase words joined by hyphens
	// (e.g. "hello-world").
	Kebab

	// ScreamingSnake is SCREAMING_SNAKE_CASE: uppercase words joined by
	// underscores (e.g. "HELLO_WORLD"). Common for constants and env vars.
	ScreamingSnake

	// Dot is dot.case: lowercase words joined by dots (e.g. "hello.world").
	Dot

	// Title is Title Case: capitalized words joined by single spaces
	// (e.g. "Hello World").
	Title

	// Sentence is Sentence case: only the first word is capitalized, the rest
	// stay lowercase, joined by single spaces (e.g. "Hello world").
	Sentence
)

// maxStyle is the highest valid Style value; keep it in sync with the
// constant block above so Valid and String stay correct as styles are added.
const maxStyle = Sentence

// String returns the lowercase identifier of the style, suitable for flags,
// configuration files and diagnostics. Unknown and out-of-range values
// return "unknown".
func (s Style) String() string {
	switch s {
	case Camel:
		return "camel"
	case Pascal:
		return "pascal"
	case Snake:
		return "snake"
	case Kebab:
		return "kebab"
	case ScreamingSnake:
		return "screaming_snake"
	case Dot:
		return "dot"
	case Title:
		return "title"
	case Sentence:
		return "sentence"
	default:
		return "unknown"
	}
}

// Valid reports whether s names a real naming convention (i.e. is not
// Unknown and is within the known range).
func (s Style) Valid() bool {
	return s >= Camel && s <= maxStyle
}

// ParseStyle maps a style identifier (as produced by Style.String) to its
// Style. The match is exact and case-sensitive. The second result is false
// for unrecognized names, in which case the Style is Unknown.
//
// A few common aliases are accepted: "constant" and "screaming" for
// ScreamingSnake.
func ParseStyle(name string) (Style, bool) {
	switch name {
	case "camel":
		return Camel, true
	case "pascal":
		return Pascal, true
	case "snake":
		return Snake, true
	case "kebab":
		return Kebab, true
	case "screaming_snake", "screaming", "constant":
		return ScreamingSnake, true
	case "dot":
		return Dot, true
	case "title":
		return Title, true
	case "sentence":
		return Sentence, true
	default:
		return Unknown, false
	}
}
