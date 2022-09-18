# $$\textcolor{yellow}{\text{GOO}}$$

## **G**olang standard library - the **OO** way

<br>

> The library you never knew you needed

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
