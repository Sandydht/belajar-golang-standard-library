# Package time
- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang.
- [https://golang.org/pkg/time/](https://golang.org/pkg/time/).

# Beberapa Function di Package time
| function | Kegunaan |
|----------|----------|
| time.Now() | Untuk mendapatkan waktu saat ini |
| time.Date(...) | Untuk membuat waktu |
| time.Parse(layout, string) | Untuk memparsing waktu dari string |

# Kode Program Package time
```go
func main() {
  now := time.Now()
  fmt.Println(now.Local())

  utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
  fmt.Println(utc)

  parse, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
  fmt.Println(parse)
}
```

# Duration
- Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi.
- Package time memiliki type Duration, yang sebenarnya adalah alias untuk int64.
- Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi duration.

# Kode Program Duration
```go
var duration1 = time.Second * 100
var duration2 = time.Millisecond * 10
var duration3 = duration1 - duration2

fmt.Printf("duration1: %d\n", duration3)
```