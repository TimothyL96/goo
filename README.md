# $$\textcolor{yellow}{\text{GOO}}$$

## **G**olang standard library - the **OO** way

<br>

> The library you never knew you needed :bulb:

<br>

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/26b98e19151f471fbd3729dbbe56c2d8)](https://www.codacy.com/gh/TimothyL96/goo/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TimothyL96/goo&amp;utm_campaign=Badge_Grade)
[![Maintainability](https://api.codeclimate.com/v1/badges/bbc9b7ccb927732ec239/maintainability)](https://codeclimate.com/github/TimothyL96/goo/maintainability)
[![CodeFactor](https://www.codefactor.io/repository/github/timothyl96/goo/badge)](https://www.codefactor.io/repository/github/timothyl96/goo)
[![Go Report Card](https://goreportcard.com/badge/github.com/timothyl96/goo)](https://goreportcard.com/report/github.com/timothyl96/goo)
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/26b98e19151f471fbd3729dbbe56c2d8)](https://www.codacy.com/gh/TimothyL96/goo/dashboard?utm_source=github.com&utm_medium=referral&utm_content=TimothyL96/goo&utm_campaign=Badge_Coverage)
[![codecov](https://codecov.io/gh/TimothyL96/goo/branch/master/graph/badge.svg?token=o61pQVm2m9)](https://codecov.io/gh/TimothyL96/goo)

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/TimothyL96/goo/tree/master.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/TimothyL96/goo/tree/master)
[![Travis Build Status](https://app.travis-ci.com/TimothyL96/goo.svg?branch=master)](https://app.travis-ci.com/TimothyL96/goo)
[![DeepSource](https://deepsource.io/gh/TimothyL96/goo.svg/?label=active+issues&show_trend=true&token=lc6AhgyQ_EjizXaVrr2ehW_K)](https://deepsource.io/gh/TimothyL96/goo/?ref=repository-badge)

<br>

Adding methods to 'builtin' types so we can simplify calling functions in OO way.

<br>

```Go
// Example:
str := goo.FromString("myString")
str.length() // same as len(str)

// Another example:
var i goo.Int
i = 30
i += 20
i.Itoa() // == "50", same as strconv.Itoa(i)

// Another example:
var s goo.String
s = "myuppercasestring"
s.ToUpper() // == "MYUPPERCASESTRING", same as strings.ToUpper(s)
```

More:

How to use Goo:
[GOO Example](https://github.com/TimothyL96/goo-example)

<br>

---  

## <b>Benchmark</b> :clock330:

The result below shows that the performances are similar when compared to builtin functions.

![Benchmarking result](https://github.com/TimothyL96/goo/blob/master/images/benchmark.jpg?raw=true)

<br>

---  

## <b>Running the source code</b> :runner:

1. Run generate for code generation: `go generate ./...`

2. Run test: `go test -v ./...`

3. Run benchmark: `go test -bench=Bench -run Bench -benchmem -benchtime 10s -v ./...`

<br>

---  

## <b>Contributing</b>
Feel free to raise issue or PR