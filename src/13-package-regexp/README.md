# Package regexp
- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression.
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
- [https://github.com/google/re2/wiki/syntax/](https://github.com/google/re2/wiki/syntax/).
- [https://golang.org/pkg/regexp/](https://golang.org/pkg/regexp/).

# Beberapa Function di Package regexp
| Function | Kegunaan |
|----------|----------|
| regexp.MustCompile(string) | Membuat regexp |
| regexp.MatchString(string) bool | Mengecek apakag regexp match dengan string |
| regexp.FindAllString(string, max) | Mencari string yang match dengan maximum jumlah hasil |

# Kode Program Package regexp
```go
func main() {
  var regex = regexp.MustCompile(`e([a-z])o`)

  fmt.Println(regex.MatchString("sandy"))
  fmt.Println(regex.MatchString("dwi"))
  fmt.Println(regex.MatchString("sAndy"))

  fmt.Println(regex.FindAllString("eko edo egi ego e10 eto", 10))
}
```