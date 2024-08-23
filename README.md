# Par

Par provides parfor and parmap functions, limiting the number of goroutines to be the number of logical cores.

## Usage

```go
objs := []int64{1, 2, 3, 4, 5}
sum := 0
par.For(objs, func(obj int64) {
    atomic.AddInt64(&sum, obj)
})
```

```go
objs := []int64{1, 2, 3, 4, 5}
res := par.Map(objs, func(obj int64) int64 {
    return obj
}, func(obj int64) int64 {
    return obj * obj
})
```
