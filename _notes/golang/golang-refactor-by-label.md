# Golang Refactor by Label

## Introduction

> Labels are declared by labeled statements and are used in the **"break"**, **"continue"**, and **"goto"** statements. - "Label scopes", The Go Programming Language Specification

## Example

* Traditional, use varible as a flag

    ```go
    func main() {
        var sum int
        stop := false // set a flag
        for i := 0; i < 10; i++ {
            for j := 0; j < 10; j++ {
                sum = i * j
                if sum > 80 {
                    stop = true // set flag to true for break outer loop
                    break
                }
            }
            if stop {
                break
            }
        }
        fmt.Println(sum)
    }
    ```

* Refactor by Label

    ```go
    func main() {
        var sum int
    FirstLoop: // define a label
        for i := 0; i < 10; i++ {
            for j := 0; j < 10; j++ {
                sum = i * j
                if sum > 80 {
                    break FirstLoop // break label scope
                }
            }
        }
        fmt.Println(sum)
    }
    ```

## Reference

* [Label Scopes@The Go Programming Language Specification](https://golang.org/ref/spec#Label_scopes)
* [Break statements@The Go Programming Language Specification](https://golang.org/ref/spec#Break_statements)