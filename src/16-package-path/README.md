# Package path
- Package path digunakan untuk memanipulasi data path seperti path URL atau path di File System.
- Secara default package path menggunakan slash sebagai karakter path-nya, oleh karena itu cocok untuk data URL.
- [https://pkg.go.dev/path/](https://pkg.go.dev/path/).
- Namun jika ingin menggunakan untuk memanipulasi path di File System, karena windows menggunakan backslash, maka khusus untuk File System, perlu menggunakan package path/filepath.
- [https://pkg.go.dev/path/filepath/](https://pkg.go.dev/path/filepath/).

# Kode Program Package path
```go
func main() {
  fmt.Println(path.Dir("hello/world.go"))
  fmt.Println(path.Base("hello/world.go"))
  fmt.Println(path.Ext("hello/world.go"))
  fmt.Println(path.Join("hello", "world", "main.go"))
}
```

# Kode Program path/filepath
```go
func main() {
  fmt.Println(filepath.Dir("hello/world.go"))
  fmt.Println(filepath.Base("hello/world.go"))
  fmt.Println(filepath.Ext("hello/world.go"))
  fmt.Println(filepath.IsAbs("hello/world.go"))
  fmt.Println(filepath.IsLocal("hello/world.go"))
  fmt.Println(filepath.Join("hello", "world", "main.go"))
}
```