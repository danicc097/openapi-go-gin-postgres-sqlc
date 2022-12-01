# inflector [![Go Package][gopkg]][gopkg-link]

Package `inflector` pluralizes and singularizes English nouns.

[gopkg]: https://pkg.go.dev/badge/github.com/kenshaw/inflector.svg (Go Package)
[gopkg-link]: https://pkg.go.dev/github.com/kenshaw/inflector

## Basic Usage

There are only two exported functions: `Pluralize` and `Singularize`.

```go
inflector.Singularize("People") // returns "Person"
inflector.Pluralize("octopus") // returns "octopuses"
```
