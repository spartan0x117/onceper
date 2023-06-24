# OncePer
A simple utility to run a command once per unique comparable key. Note that it uses a mutex and a map so it will not be as
performant as a sync.Once. 

This is intended to be used when you have a dynamic list of keys for which you want to run a function
only once, each.

## Usage

### `Do`
Do runs the passed function once per unique key.

```go
package main

import (
    "fmt"

    "github.com/spartan0x117/onceper"
)

func main() {
    o := onceper.New[string]()

    // Run a function once per unique key
    o.Do("key", func() {
        fmt.Println("Hello World!")
    })
    // Any subsequent calls to Do with the same key will not run the function, even
    // if the function is different
    o.Do("key", func() {
        fmt.Println("Goodbye World!")
    })

    // A different key will run the function again
    o.Do("key2", func() {
        fmt.Println("foo bar baz")
    })
}
```
Output:
```
Hello World!
foo bar baz
```

### `DoWith`
DoWith runs the passed function once per unique key, passing the key to the function.


```go
package main

import (
    "fmt"

    "github.com/spartan0x117/onceper"
)

func main() {
    o := onceper.New[string]()
    f := func(s string) {
        fmt.Printf("Hello %s!\n", s)
    }

    // Run a function once per unique key, passing the key to the function
    o.DoWith("world", f)
    // Any subsequent calls with the same key will not run
    o.DoWith("world", f)

    // A different key will run the function again
    o.DoWith("go", f)
}
```
Output:
```
Hello world!
Hello go!
```