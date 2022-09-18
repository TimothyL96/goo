# $$\textcolor{yellow}{\text{GOO}}$$

## **G**olang standard library - the **OO** way

<br>

> The library you never knew you needed :bulb:

<br>

[![Maintainability](https://api.codeclimate.com/v1/badges/bbc9b7ccb927732ec239/maintainability)](https://codeclimate.com/github/TimothyL96/goo/maintainability)
[![CodeFactor](https://www.codefactor.io/repository/github/timothyl96/goo/badge)](https://www.codefactor.io/repository/github/timothyl96/goo)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/TimothyL96/goo/tree/master.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/TimothyL96/goo/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/timothyl96/goo)](https://goreportcard.com/report/github.com/timothyl96/goo)
[![codecov](https://codecov.io/gh/TimothyL96/goo/branch/master/graph/badge.svg?token=o61pQVm2m9)](https://codecov.io/gh/TimothyL96/goo)
[![Travis Build Status](https://app.travis-ci.com/TimothyL96/goo.svg?branch=master)](https://app.travis-ci.com/TimothyL96/goo)
[![DeepSource](https://deepsource.io/gh/TimothyL96/goo.svg/?label=active+issues&show_trend=true&token=lc6AhgyQ_EjizXaVrr2ehW_K)](https://deepsource.io/gh/TimothyL96/goo/?ref=repository-badge)

<br>

Adding methods to 'builtin' types so we can simplify calling functions in OO way.

<br>

```Go
// Example:
str := goo.FromString("myString")
str.length() // compared to len(str)

// Another example:
var i goo.Int
i = 30
i += 20
i.Itoa() // == 50, compared to strconv.Itoa(i)

// Another example:
var s goo.String
s = "myuppercasestring"
s.ToUpper() // == "MYUPPERCASESTRING", compared to strings.ToUpper(s)
```

More:

How to use Goo:
[GOO Example](https://github.com/TimothyL96/goo-example)

<br>

---  

## <b>Running the source code</b>

Run generate for code generation: `go generate ./...`

Run test: `go test -v ./...`

Run benchmark: `go test -bench=Bench -run Bench -benchmem -benchtime 10s -v ./...`
