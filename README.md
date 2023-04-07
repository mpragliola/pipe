# pipe

This Go library was written as an exercise to familiarize with the language.

## What is `pipe`?

With `pipe` you can implement asynchronous pipelines in Golang without
having to worry too much about channels and concurrency.

```go
emit := pipe.Pipe(square, pipe.Of(1,2,3,4))

for r := range emit {
    fmt.Println(r)
}

// 1
// 4
// 9
// 16
```

## API and examples

### Of

Value emitter.

```go
emit := pipe.Of(1,2,3,4)

for r := range emit {
    fmt.Println(r)
}

// 1
// 2
// 3
// 4
```

You can emit array or slices with:

```go
mySlice := []int{1,2,3}
emit := pipe.Of(mySlice...)
```

### OfFunc

Uses a generator function to emit values. It can leverage Golang's
`context.Context` to provide cancellation, timeout, etc.

```go
ctx, cancel := context.WithCancel(context.Background())
i, j := 0, 1
p := pipe.OfFunc(
    ctx,
    func() int {
        i, j = j, j+i

        if i > 10 {
            cancel()
        }

        return i
    },
)

for r := range p {
    fmt.Println(r)
}

// 1
// 1
// 2
// 3
// 5
// 8
// 13
```

### Pipe

```go
emit := pipe.Of(1,2,3,4)

square := func(i int) int {
    return i*i
}

for r := range pipe.Pipe(square, emit) {
    fmt.Println(r)
}

// 1
// 4
// 9
// 16
```

An example of composable pipes:

```go
emit := pipe.Of(1,2,3)

square := func(i int) int { return i*i }
add10 := func(i int) int { return i+10 }
stringify := func(i int) string { return strconv.Itoa(i) }

for r := range pipe.Pipe(stringify(pipe.Pipe(add10, pipe.Pipe(square, emit))) {
    fmt.Println(r)
}

// "11"
// "14"
// "19"
```

### Filter

```go
input := []int{1, 2, 3, 4, 5, 6}
oddFilter := func(i int) bool { return i%2 == 0 }

for r := range pipe.Filter(oddFilter, pipe.Of(input...)) {
    fmt.Println(r)
}

\\ 2
\\ 4
\\ 6
```