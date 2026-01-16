# Package fmt
- Sebelumnya kita sudah sering menggunakan package ```fmt``` dengan menggunakan function ```Println```.
- Selain ```Println```, masih banyak function yang terdapat di package ```fmt```, contohnya banyak digunakan untuk melakukan format.
- [https://pkg.go.dev/fmt](https://pkg.go.dev/fmt).

# Kode Program Package fmt
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello, world!")

  firstName := "Sandy"
  lastName := "Dwi"

  fmt.Printf("Hello %s %s!\n", firstName, lastName)
}
```