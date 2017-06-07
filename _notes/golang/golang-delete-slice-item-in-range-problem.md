# Golang Delete Slice Item in Range Problem

## Problem

When range start, it is not iterate value one by one, it is iterate by index.

If you modified a slice while it is in range, it will induce some problem.

## Example

Remove char a from chars.

* Incorrect Way

    ```go
    chars := []string{"a", "a", "b"}

    for i, v := range chars {
        fmt.Printf("%+v, %d, %s\n", chars, i, v)
        if v == "a" {
            chars = append(chars[:i], chars[i+1:]...)
        }
    }
    fmt.Printf("%+v", chars)

    ```

  * Expected

    ```go
    [a a b], 0, a
    [a b], 0, a
    [b], 0, b
    Result: [b]
    ```

  * Actual

    ```go
    // Autual
    [a a b], 0, a
    [a b], 1, b
    [a b], 2, b
    Result: [a b]
    ```
  [Playground](https://play.golang.org/p/BqMpRGlF2H)

* Correct Way

    ```go
    chars := []string{"a", "a", "b"}

    for i := 0; i < len(chars); i++ {
        if chars[i] == "a" {
            chars = append(chars[:i], chars[i+1:]...)
            i-- // form the remove item index to start iterate next item
        }
    }

    fmt.Printf("%+v", chars)
    ```

    [Playground](https://play.golang.org/p/z-ODXAF8xU)