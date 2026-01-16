# Package Strings
- Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String.
- Ada banyak sekali function yang bisa kita gunakan.
- [https://golang.org/pkg/strings/](https://golang.org/pkg/strings/).

# Beberapa Function di Package Strings
| Function | Kegunaan |
|----------|----------|
| strings.Trim(string, cutset) | Memotong cutset di awal dan akhir string |
| strings.ToLower(string) | Membuat semua karakter string menjadi lower case |
| strings.ToUpper(string) | Membuat semua karakter string menjadi upper case |
| strings.Split(string, separator) | Memotong string berdasarkan separator |
| strings.Contain(string, search) | Mengecek apakag string mengandung string lain |
| strings.ReplaceAll(string, from, to) | Mengubah semua string dari from ke to |

# Kode Program Package Strings
```go
package main

import "fmt"

func main() {
  fmt.Println(strings.Contains("Sandy Dwi Handoko Trapsilo", "Sandy"))
  fmt.Println(strings.Split("Sandy Dwi Handoko Trapsilo", " "))
  fmt.Println(strings.ToLower("Sandy Dwi Handoko Trapsilo"))
  fmt.Println(strings.ToUpper("Sandy Dwi Handoko Trapsilo"))
  fmt.Println(strings.Trim("   Sandy Dwi Handoko Trapsilo    ", " "))
  fmt.Println(strings.ReplaceAll("Sandy Sandy Sandy", "Sandy", "Dwi"))
}
```