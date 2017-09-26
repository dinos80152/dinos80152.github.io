# Golang All about Testing

## Command

```bash
go test -run=Prime -v -timeout 100ms -short -parallel 2 cnet/ctcp
```

```bash
go test -cover -covermode set -coverpkg bufio,net -coverprofile cover.out cnet/ctcp
```

```bash
go tool cover -func=cover.out
go tool cover -html=cover.out
```

## TestMain

1. TestMain could be used to setup something before a testing run and teardown something after testing finish.
  * setup:
    * export environment variable
    * presetting some testing data
      * database
      * session
      * cache
      * etc...
  * teardown:
    * clear data or environment variable produces by testing
2. When the testing file have TestMain

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

```bash

```

## t.Log() and t.Logf()

1. t.Log() and t.Logf() only print when testing is FAIL in normal mode
1. If you wanna print log when testing PASS, **add -v flag to print t.Log() and t.Logf()**

## End to End HTTP Testing

net/http/httptest

### main.go

```go
func nameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "dinolai")
}

func upperNameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", strings.ToUpper(r.URL.Path[len("/uppername/"):]))
}

func router() *http.ServeMux {
	h := http.NewServeMux()
	h.HandleFunc("/name", nameHandler)
	h.HandleFunc("/uppername/", upperNameHandler)
	return h
}

func main() {
	http.ListenAndServe(":8080", router())
}

```

### Test from Handler

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

func TestUpperNameHandler(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"string", "dinolai", "DINOLAI"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/uppername/"+tt.input, nil)
			w := httptest.NewRecorder()
			upperNameHandler(w, req)

			got := w.Body.String()
			if got != tt.want {
				t.Errorf("Got %s, Want %s", got, tt.want)
			}
		})
	}
}
```

### Test from Router

```go
func TestFromRouter(t *testing.T) {
	tests := []struct {
		name string
		uri  string
		want string
	}{
		{"name", "name", "dinolai"},
		{"uppername", "uppername/dinolai", "DINOLAI"},
	}

	ts := httptest.NewServer(router())
	defer ts.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := ts.URL + "/" + tt.uri
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

## Mock

[official mock package](https://github.com/golang/mock)

## Reference

* [Main - Package testing](https://golang.org/pkg/testing/#hdr-Main)
* [Testing Techniques](https://talks.golang.org/2014/testing.slide)