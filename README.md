# pipe

## Of

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

## Pipe

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

```go
emit := pipe.Of(1,2,3)

square := func(i int) int { return i*i }
add10 := func(i int) int { return i+10 }
stringify := func(i int) string { return strconv.Itoa(i) }

for r := range pipe.Pipe(stringify(pipe.Pipe(add10, pipe.Pipe(square, emit))) {
    fmt.Println(r)
}

```
