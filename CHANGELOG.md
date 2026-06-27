# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - Unreleased

Version 2 is a full redesign around a single universal tokenizer. The module
path is now `github.com/goloop/scs/v2` and requires Go 1.24.

### Added
- Universal tokenizer exposed as `Split(s) []string` and `Words(s) iter.Seq[string]`.
- New case styles: `SCREAMING_SNAKE_CASE` (`ToScreamingSnake`), `dot.case`
  (`ToDot`) and `Title Case` (`ToTitle`).
- `Convert(to Style, s string) string` for runtime-selected styles, plus
  `Style.String`, `Style.Valid` and `ParseStyle`.
- `Detect(s string) (Style, bool)` with an unambiguous contract, and the generic
  `Is(style, s)` predicate.
- Reusable, concurrency-safe `Caser` with `New` and the `WithAcronyms` option
  for opt-in Go-style initialisms.

### Changed
- All converters are now **total**: `ToCamel`, `ToPascal`, `ToSnake`, `ToKebab`
  take a string and return a string; they never return an error and correctly
  handle input that is already in any style.
- Default acronym handling is Title casing (`Http`, `Api`, `Id`), which always
  round-trips. All-caps initialisms are opt-in via `WithAcronyms`.
- Digits attach to neighboring letters and no longer split a word on their own;
  only an explicit separator splits a number out.

### Removed
- The lossy `StrTo*` family and the twelve error-returning pairwise converters
  (`CamelToKebab`, `SnakeToPascal`, …) — superseded by the total `To*` functions
  and `Convert`.
- The stateful `StringCaseStyle` object (`New`/`Eat`/`Set`/`Value`/`CopyTo*`),
  replaced by pure functions and the immutable `Caser`.

### Fixed
- Case-based word boundaries are preserved: `ToSnake("userID")` is `user_id`,
  not `userid`.
- Adjacent words no longer merge into an unsplittable block by default.
- Number handling is consistent across every entry point.
- Detector overlap is resolved by `Detect`'s single-match contract.
