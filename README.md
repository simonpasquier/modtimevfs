[![GoDoc](https://godoc.org/github.com/simonpasquier/modtimevfs?status.svg)](https://godoc.org/github.com/simonpasquier/modtimevfs)
[![CircleCI](https://circleci.com/gh/simonpasquier/modtimevfs.svg?style=svg)](https://circleci.com/gh/simonpasquier/modtimevfs)

`modtimevfs` provides an implementation of `http.FileSystem` that wraps another `FileSystem` and returns the same (fixed) modification time for all files in the collection. 

It is useful in combination with [`vfsgen`](https://github.com/shurcooL/vfsgen) to generate Go code that is identical when generated from different environments. For more details, see (https://github.com/shurcooL/vfsgen/issues/26).

## Example

```go
fs = modtimevfs.New(Assets, time.Unix(1, 0))
err := vfsgen.Generate(fs)
if err != nil {
    panic(err)
}
```

## License

Apache License 2.0, see [LICENSE](https://github.com/simonpasquier/modtimevfs/blob/master/LICENSE).
