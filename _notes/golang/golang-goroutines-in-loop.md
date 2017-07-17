# Golang Goroutines in Loop

## Problem

```go
func main() {
    for _, v := range values {
        go func() {
            fmt.Println(v)
        }()
    }
}
```

> each iteration of the loop uses the same instance of the variable v, so each closure shares that single variable. When the closure runs, it prints the value of v at the time fmt.Println is executed, but v may have been modified since the goroutine was launched. - "What happens with closures running as goroutines?", Golang Frequently Asked Questions (FAQ)

## Example

* Incorrect

    ```go
    func main() {
        ints := []int{1,2,3}
        for _, i := range ints {
            go func() {
                fmt.Println(i)
            }()
        }
        time.Sleep(3 * time.Second)
        // Output:
        // 3
        // 3
        // 3
    }
    ```

* Fixed: pass the variable as an argument to the closure.

    ```go
    func main() {
        ints := []int{1,2,3}
        for _, i := range ints {
            go func(i int) {
                fmt.Println(i)
            }(i)
        }
        time.Sleep(3 * time.Second)
        // Output:
        // 1
        // 2
        // 3
    }
    ```

* Fixed: copy value to a new variable, because of variables declared within the body of a loop are not shared between iterations.

    ```go
    func main() {
        ints := []int{1,2,3}
        for _, i := range ints {
            i := i
            go func() {
                fmt.Println(i)
            }()
        }
        time.Sleep(3 * time.Second)
        // Output:
        // 1
        // 2
        // 3
    }
    ```

## Example 2: Call Struct Method

* Incorrect

    ```go
    type user struct {
        name string
    }

    func(u *user) Print() {
        fmt.Println(u.name)
    }

    func main() {
        users := []user{{"dino"}, {"liam"}, {"mei"}}
        for _, user := range users {
                go user.Print()
        }
        time.Sleep(3 * time.Second)
        // Output:
        // mei
        // mei
        // mei
    }
    ```

* Fixed: pass by reference

    ```go

    type user struct {
        name string
    }

    func(u *user) Print() {
        fmt.Println(u.name)
    }

    func main() {
        users := []*user{{"dino"}, {"liam"}, {"mei"}}
        for _, user := range users {
                go user.Print()
        }
        time.Sleep(3 * time.Second)
        // Output:
        // dino
        // liam
        // mei
    }
    ```

## Reference

* [Iteration Variables and Closures in "for" Statements@50 Shades of Go: Traps, Gotchas, and Common Mistakes for New Golang Devs](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#closure_for_it_vars)
* [CommonMistakes@golang github wiki](https://github.com/golang/go/wiki/CommonMistakes)
* [What happens with closures running as goroutines?@golang doc faq](https://golang.org/doc/faq#closures_and_goroutines)