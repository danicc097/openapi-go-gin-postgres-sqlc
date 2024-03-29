#### SMARTY DISCLAIMER: Subject to the terms of the associated license agreement, this software is freely available for your use. This software is FREE, AS IN PUPPIES, and is a gift. Enjoy your new responsibility. This means that while we may consider enhancement requests, we may or may not choose to entertain requests at our sole and absolute discretion.

[![Build Status](https://travis-ci.org/smartystreets/gunit.svg?branch=master)](https://travis-ci.org/smartystreets/gunit)
[![Code Coverage](https://codecov.io/gh/smartystreets/gunit/branch/master/graph/badge.svg)](https://codecov.io/gh/smartystreets/gunit)
[![Go Report Card](https://goreportcard.com/badge/github.com/smartystreets/gunit)](https://goreportcard.com/report/github.com/smartystreets/gunit)
[![GoDoc](https://godoc.org/github.com/smartystreets/gunit?status.svg)](http://godoc.org/github.com/smartystreets/gunit)

# gunit

## Installation

```
$ go get github.com/smartystreets/gunit
```

-------------------------

We now present `gunit`, yet another testing tool for Go.

> Not again... ([GoConvey](http://goconvey.co) was crazy enough...but sort of cool, ok I'll pay attention...)

No wait, this tool has some very interesting properties. It's a mix of good things provided by the built-in testing package, the [assertions](https://github.com/smartystreets/assertions) you know and love from the [GoConvey](http://goconvey.co) project, the [xUnit](https://en.wikipedia.org/wiki/XUnit) testing style (the first real unit testing framework), and it's all glued together with `go test`.

> Blah, blah, yeah, yeah. Ok, so what's wrong with just using the standard "testing" package? What's better about this `gunit` thing?

The convention established by the "testing" package and the `go test` tool only allows for local function scope:

```
func TestSomething(t *testing.T) {
	// blah blah blah
}
```

This limited scope makes extracting functions or structs inconvenient as state will have to be passed to such extractions or state returned from them. It can get messy to keep a test nice and short. Here's the basic idea of what the test author using `gunit` would implement in a `*_test.go` file:

```go

package examples

import (
    "time"
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestExampleFixture(t *testing.T) {
	gunit.Run(new(ExampleFixture), t)
}

type ExampleFixture struct {
	*gunit.Fixture // Required: Embedding this type is what makes the magic happen.

	// Declare useful state here (probably the stuff being tested, any fakes, etc...).
}

func (this *ExampleFixture) SetupStuff() {
	// This optional method will be executed before each "Test"
	// method (because it starts with "Setup").
}
func (this *ExampleFixture) TeardownStuff() {
	// This optional method will be executed after each "Test"
	// method (because it starts with "Teardown"), even if the test method panics.
}


// This is an actual test case:
func (this *ExampleFixture) TestWithAssertions() {
	// Here's how to use the functions from the `should`
	// package at github.com/smartystreets/assertions/should
	// to perform assertions:
	this.So(42, should.Equal, 42)
	this.So("Hello, World!", should.ContainSubstring, "World")
}

func (this *ExampleFixture) SkipTestWithNothing() {
	// Because this method's name starts with 'Skip', it will be skipped.
}

func (this *ExampleFixture) LongTestSlowOperation() {
	// Because this method's name starts with 'Long', it will be skipped if `go test` is run with the `short` flag.
	time.Sleep(time.Hour)
	this.So(true, should.BeTrue)
}
```

-------------------------

> So, I see just one traditional test function and it's only one line long. What's the deal with that?

Astute observations. `gunit` allows the test author to use a _struct_ as the scope for a group of related test cases, in the style of [xUnit](https://en.wikipedia.org/wiki/XUnit) fixtures. This makes extraction of setup/teardown behavior (as well as invoking the system under test) much simpler because all state for the test can be declared as fields on a struct which embeds the `Fixture` type from the `gunit` package. All you have to do is create a Test function and pass a new instance of your fixture struct to gunit's Run function along with the *testing.T and it will run all defined Test methods along with the Setup and Teardown method.

Enjoy.

### Parallelism
By default all fixtures are run in parallel as they should be independent, but if you for some reason have fixtures which need to be run sequentially, you can change the `Run()` method to `RunSequential()`, e.g. in the above example

```go
func TestExampleFixture(t *testing.T) {
	gunit.RunSequential(new(ExampleFixture), t)
}
```

[Advanced Examples](https://github.com/smartystreets/gunit/tree/master/advanced_examples)

----------------------------------------------------------------------------

For users of JetBrains IDEs, here's LiveTemplate you can use for generating the scaffolding for a new fixture:

- Abbreviation: `fixture`
- Description: `Generate gunit Fixture boilerplate`
- Template Text:

```
func Test$NAME$(t *testing.T) {
    gunit.Run(new($NAME$), t)
}

type $NAME$ struct {
    *gunit.Fixture
}

func (this *$NAME$) Setup() {
}

func (this *$NAME$) Test$END$() {
}

```

Be sure to specify that this LiveTemplate is applicable in Go files.
