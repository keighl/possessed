# Possessed

[![Build Status](https://travis-ci.org/keighl/possessed.png?branch=master)](https://travis-ci.org/keighl/possessed) [![Coverage Status](https://coveralls.io/repos/keighl/possessed/badge.svg)](https://coveralls.io/r/keighl/possessed)

Golang functions for converting an English string to/from its possessive form.

```go
package main

import (
    ps "github.com/keighl/possessed"
    "fmt"
)

func main() {
    fmt.Println(ps.Possess("Dave"))
    // => Dave's

    fmt.Println(ps.Possess("Gladys"))
    // => Gladys'

    fmt.Println(ps.Possess("it"))
    // => its

    fmt.Println(ps.Unpossess("Dave's"))
    // => Dave

    fmt.Println(ps.Unpossess("Gladys'"))
    // => Gladys

    fmt.Println(ps.Unpossess("its"))
    // => it

    // Change the apostrophe character
    ps.APOSTROPHE_CHAR = `’`
    fmt.Println(ps.Possess("Dave"))
    // => Dave’s
}
```
