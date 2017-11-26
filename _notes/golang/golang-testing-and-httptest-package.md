# Golang testing and httptest Package

* [Test Full Command Example](#test-full-command-example)
* [TestMain](#testmain)
  * [Example](#example)
* [t.Log() and t.Logf()](#tlog-and-tlogf)
* [httptest: End to End HTTP Testing](#httptest-end-to-end-http-testing)
  * [Example](#example-1)
    * [main.go](#maingo)
    * [main_test.go](#main_testgo)
* [Ignore File from Test Coverage](#ignore-file-from-test-coverage)
* [Reference](#reference)

## Test Full Command Example

* Test

```bash
go test -run=Prime -v -timeout 100ms -short -parallel 2 cnet/ctcp
```

* Coverage

```bash
go test -cover -covermode set -coverpkg bufio,net -coverprofile cover.out cnet/ctcp
```

* Open a web browser displaying annotated source code by coverage profile

```bash
go tool cover -html=cover.out
```

## TestMain

```go
func TestMain(m *testing.M)
```

* When TestMain function in package, every test will execute TestMain before test running.
* TestMain could be used to setup something before a testing run and teardown something after testing finish.
  * setup:
    * export environment variable
    * presetting some testing data
      * database
      * session
      * cache
      * etc...
  * teardown:
    * clear data or environment variable produces by testing

### Example

```go
func TestMain(m *testing.M) {
	setup()
	v := m.Run() // run testing
    teardown()
	os.Exit(v)
}

func setup() {
	fmt.Println("before testing")
}

func teardown() {
	fmt.Println("after testing")
}
```

## t.Log() and t.Logf()

```go
func TestFunc(t *testing.T) {
	t.Log("Print Something")
	t.Logf("Print %s", str)
}
```

* t.Log() and t.Logf() **only print when testing is FAIL** in normal mode
* If you wanna print log when testing PASS, **use `go test -v` to print t.Log() and t.Logf()**

## httptest: End to End HTTP Testing

```go
package net/http/httptest
```

> Package httptest provides utilities for HTTP testing.

### Example

#### main.go

```go
func nameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "dinolai")
}

func router() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc("/name", nameHandler)
	return h
}

func main() {
	http.ListenAndServe(":8080", router())
}

```

#### main_test.go

* Testing from handler

	```go
	func TestNameHandler(t *testing.T) {
		tests := []struct {
			name string
			want string
		}{
			{"string", "dinolai"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				req, _ := http.NewRequest("GET", "", nil)
				w := httptest.NewRecorder()
				nameHandler(w, req)

				got := w.Body.String()
				if got != tt.want {
					t.Errorf("Got %s, Want %s", got, tt.want)
				}
			})
		}
	}

	```

* Test from Router

	```go
	func TestFromRouter(t *testing.T) {
		tests := []struct {
			name string
			uri  string
			want string
		}{
			{"name", "/name", "dinolai"},
		}

		ts := httptest.NewServer(router())
		defer ts.Close()

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				url := ts.URL + tt.uri
				resp, _ := http.Get(url)
				respBody, _ := ioutil.ReadAll(resp.Body)
				resp.Body.Close()

				got := string(respBody)
				if got != tt.want {
					t.Errorf("Got %s, Want %s", got, tt.want)
				}
			})
		}
	}
	```

## Ignore File from Test Coverage

Use [Build Tag](https://golang.org/pkg/go/build/#hdr-Build_Constraints)

* add tags on the top of file, **must have the `line break` between tags and package.**

	```go
	//+build !test

	package main

	func main() {
		// codes
	}
	```

* add `-tags` flag when test with coverage

	```bash
	$ go test -cover -tags test
	ok      github.com/dinos80152/golang-lab/main    0.142s  coverage: 100% of statements
	```

## Reference

* [Main - Package testing](https://golang.org/pkg/testing/#hdr-Main)
* [Package httptest](https://golang.org/pkg/net/http/httptest/)
* [Testing Techniques](https://talks.golang.org/2014/testing.slide)
* [Ignore code blocks in Golang test coverage calculation](https://stackoverflow.com/questions/25511076/ignore-code-blocks-in-golang-test-coverage-calculation)