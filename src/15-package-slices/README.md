# Package Slices
- Di Go-Lang versi terbaru, terdapat fitur bernama generic.
- Fitur generic ini memungkinkan kita bisa membuat parameter dengan tipe yang bisa berubah-ubah, tanpa harus menggunakan interface kosong / any.
- Salah satu package yang menggunakan fitur generic ini adalah package slices.
- Package slices ini digunakan untuk memanipulasi data di slice.
- [https://pkg.go.dev/slices/](https://pkg.go.dev/slices/).

# Kode Program Package slices
```go
func main() {
  names := []string{"John", "Paul", "George", "Ringo"}
  values := []int{100, 95, 80, 90}

  fmt.Println(slices.Min(values))
  fmt.Println(slices.Max(values))
  fmt.Println(slices.Contains(names, "Paul"))
  fmt.Println(slices.Index(names, "George"))
}
```