# $$\textcolor{yellow}{\text{GOO}}$$

## **G**olang standard library Extension + the **OO** way

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

There are 2 ways to use this library.

1. Use the methods added to '_builtin_' types so we can simplify calling functions in OO way.

    This includes builtin + extended functions like `unique()`.

 2. :exclamation: Or, simply apply the extended functions in Go style.

<br>

Example - Extended functions in Go Style:

```Go
s := []int{1, 1, 1, 2, 3, 3}
_ = goo.Unique(s) // returns {1, 2, 3}

toGet := []string{"a", "a", "b", "b", "b", "b", "c"}
_ = goo.Unique(toGet) // returns {a, b, c}
```

<br>

Example - OOP Style:

```Go
// ##### Example:
// ### Length:
str1 := goo.FromString("myString")
str1.length() // same as len(str1)

// ### Arithmetic and Itoa()
i1 := goo.FromInt(30)
i1 += 20
i1.Itoa() // same as strconv.Itoa(i1)

// ### ToUpper():
var str2 goo.String = "myuppercasestring"
str2.ToUpper() // same as strings.ToUpper(str2)

// ##### Slice
// ### Append()
s1 := goo.FromSlice([]int{1, 2, 3})
s1 = s1.Append(4, 5, 6) // same as s1 = append(s1, 4, 5, 6)

// ##### Extension library example:
// ### Unique()
s2 := goo.FromSlice([]goo.Int{1, 1, 2, 3, 3})
s2 = s2.Unique() // == [1, 2, 3]



// ##### How to declare the variable:
// ### Base type
// Using Goo function
t1 := goo.FromString("myString")

// Using default = operator, but explicitly specify the type
var t2 goo.Int = 30

// Using explicit type conversion
t3 := goo.Int(2) 

// ### Declaring Slice
// Using NewSlice (recommended)
mySlice := goo.NewSlice(1, 2, 3)

// Using Goo function with builtin int
t4 := goo.FromSlice([]int{1, 2, 3})
// Same as above but with goo.Int
t5 := goo.FromSlice([]goo.Int{1, 2, 3})

// Using default = operation, but explicitly specify the type for the generic
var t6 goo.Slice[int] = []int{1, 2, 3}
// Same as above but with goo.Int
var t7 goo.Slice[goo.Int] = []goo.Int{1, 2, 3}

// ### Declaring Map
t8 := goo.FromMap(make(map[int]int))
var t9 goo.Map[int, struct{}] = make(map[int]struct{})

```

<br>

If typing `goo.` is too much, use dot import for this library.

<br>

More:

How to use Goo:
[GOO Example](https://github.com/TimothyL96/goo-example)

---  

## <b>Benchmark</b> :clock330:

The result below shows that the performances are similar when compared to builtin functions.

![Benchmarking result](https://github.com/TimothyL96/goo/blob/master/images/benchmark.jpg?raw=true)

<br>

---  

## <b>Running the source code</b> :runner:

There are 3 useful commands: 

Run `go generate` for code generation :point_right: `go generate ./...`
    
- This will generate most of the goo types e.g. string.go, int.go
- Check the gen/ folder for more information

Run `go test` :point_right: `go test -v ./...`

Run `go benchmark` :point_right: `go test -bench=Bench -run Bench -benchmem -benchtime 10s -v ./...`

<br>

---  

## <b>Contributing</b>
Feel free to raise issue or PR
