# Package io
- IO atau singkatan dari Input Output, merupakan fitur di Go-Lang yang digunakan sebagai standard untuk proses Input Output.
- Di Go-Lang, semua mekanisme input output pasti mengikuti standard package io.
- [httos://pkg.go.dev/io/](httos://pkg.go.dev/io/).

# Reader
- Untuk membaca input, Go-Lang menggunakan kontrak interface bernama Reader yang terdapat di package io.
```go
// Implementations must not retain p.
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

# Writer
- Untuk menulis ke output, Go-Lang menggunakan kontrak interface bernama Writer yang terdapat di package io.
```go
// Implementations must not retain p.
type Writer interface {
  Write(p []byte) (n int, err error)
}
```

# Implementasi IO
- Penggunaan dari IO sendiri di Go-Lang terdapat di banyak package, sebelumnya contoh-nya kita menggunakan CSV Reader dan CSV Writer.
- Karena package IO sebenarnya hanya kontrak untuk IO, untuk implementasinya kita harus lakukan sendiri.
- Tapi untungnya, Go-Lang juga menyediakan package untuk mengimplementasikan IO secara mudah, yaitu menggunakan package bufio.