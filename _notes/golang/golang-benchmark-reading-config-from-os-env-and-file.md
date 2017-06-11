# Golang Benchmark Reading Config from OS ENV and File

## Intro

Compare reading config from a json file and os enviroment variable.

## Conclusion

***Cache your config by variable whatever you use OS ENV or json file***

Get OS ENV is 30 times faster than read json file.

Caching config by variable is 1000 times faster than OS ENV

Caching config by variable is 30000 times faster than reading json file

## Example

```go
package main

import (
    "fmt"
    "os"

    "io/ioutil"

    "encoding/json"
)

func main() {
    os.Setenv("TESTENV", "abc")
    fmt.Println(ReadConfigFromENV())
    fmt.Println(ReadConfigFromENVWithCached())
    fmt.Println(ReadConfigFromFile())
    fmt.Println(ReadConfigFromFileWithCached())
}

func ReadConfigFromENV() string {
    return os.Getenv("TESTENV")
}

var osenv string

func ReadConfigFromENVWithCached() string {
    if osenv != "" {
        return osenv
    }
    osenv = os.Getenv("TESTENV")
    return osenv
}

func ReadConfigFromFile() string {
    config := readFile()
    return config["TESTENV"]
}

var filenv string

func ReadConfigFromFileWithCached() string {
    if filenv != "" {
        return filenv
    }
    config := readFile()
    filenv = config["TESTENV"]
    return filenv
}

func readFile() map[string]string {
    data, _ := ioutil.ReadFile("config.json")
    var config map[string]string
    json.Unmarshal(data, &config)
    return config
}

```

```go
package main

import (
    "os"
    "testing"
)

func TestMain(m *testing.M) {
    os.Setenv("TESTENV", "abc")
    os.Exit(m.Run())
}

func Benchmark_ReadConfigFromENV(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ReadConfigFromENV()
    }
}

func Benchmark_ReadConfigFromENVWithCached(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ReadConfigFromENVWithCached()
    }
}

func Benchmark_ReadConfigFromFile(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ReadConfigFromFile()
    }
}

func Benchmark_ReadConfigFromFileWithCached(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ReadConfigFromFileWithCached()
    }
}

```

```bash
$ go test -bench=.
Benchmark_ReadConfigFromENV-2                     300000              3686 ns/op
Benchmark_ReadConfigFromENVWithCached-2         300000000                4.82 ns/op
Benchmark_ReadConfigFromFile-2                     10000            107606 ns/op
Benchmark_ReadConfigFromFileWithCached-2        300000000                4.45 ns/op
PASS
ok      github.com/dinos80152/golang-leetcode   6.126s
```