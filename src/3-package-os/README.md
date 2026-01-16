# Package os
- Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os.
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi).
- [https://golang.org/pkg/os](https://golang.org/pkg/os).

# Kode Program Package os (1)
```go
package main

import (
  "fmt"
  "os"
)

func main() {
  args := os.Args
  fmt.Println(args)
}
```

# Kode Program Package os (2)
```go
package main

import (
  "fmt"
  "os"
)

func main() {
  hostname, err := os.Hostname()

  if err == nil {
    fmt.Println(hostname)
  } else {
    fmt.Println("Error", err.Error())
  }
}
```