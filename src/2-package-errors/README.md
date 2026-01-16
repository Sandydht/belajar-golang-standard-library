# Package Errors
- Sebelumnya kita sudah membahas tentang interface error yang merupakan representasi dari error di Go-Lang, dan membuat error menggunakan function ```errors.New()```.
- Sebenarnya masih banyak yang bisa kita lakukan menggunakan package errors, contohnya ketika kita ingin membuat beberapa value error yang berbeda.
- [https://pkg.go.dev/errors](https://pkg.go.dev/errors).

# Kode Program Membuat Error
```go
var (
  ValidationError = errors.New("validation error")
  NotFoundError = errors.New("not found error")
)
```

# Kode Program Menggunakan Error
```go
func GetById(id string) error {
  if id == "" {
    return ValidationError
  }

  if id != "sandy" {
    return NotFoundError
  }

  return nil
}
```

# Mengecek Jenis Error
- Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis error-nya.
- Kita bisa menggunakan ```errors.Is()``` untuk mengecek jenis type error-nya.

# Kode Program Mengecek Jenis Error
```go
func main() {
  err := GetById("")

  if err != nil {
    if errors.Is(err, ValidationError) {
      fmt.Println("validation error")
    } else if errors.Is(err, NotFoundError) {
      fmt.Println("not found error")
    } else {
      fmt.Println("unknown error")
    }
  }
}
```