# Package bufio
- Package bufio atau singkatan dari buffered io.
- Package ini digunakan untuk membuat data IO seperti Reader dan Writer.
- [https://pkg.go.dev/bufio/](https://pkg.go.dev/bufio/).

# Kode Program Reader
```go
func main() {
  input := strings.NewReader("this is long string\nwith new line\n")
  reader := bufio.NewReader(input)

  for {
    line, _, err := reader.ReadLine()
    if err == io.EOF {
      break
    }
    fmt.Println(string(line))
  }
}
```

# Kode Program Writer
```go
func main() {
  writer := bufio.NewWriter(os.Stdout)
  writer.WriteString("hello world\n")
  writer.WriteString("selamat belajar")
  writer.Flush()
}
```