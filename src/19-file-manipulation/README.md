# File Management
- Di package os, terdapata file management, namun sengaja ditunda pembahasannya, karena kita harus tau dulu tentang IO.
- Saat kita membuat atau membaca file menggunakan package os, struct file merupakan implementasi dari io.Reader dan io.Writer.
- Oleh karena itu, kita bisa melakukan baca dan tulis terhadap file tersebut menggunakan package io / bufio.

# Open File
- Untuk membuat / membaca file, kita bisa menggunakan os.OpenFile(name, flag, permission).
- **name** berisikan nama file, bisa absolute atau relative / local.
- **flag** merupakan penanda file, apakah untuk membaca, menulis, dan lain-lain.
- **permission**, merupakan permission yang diperlukan ketika membuat file, bisa kita simulasikan disini: [https://chmod-calculator.com/](https://chmod-calculator.com/).

# Kode Program File Flag di Package os
```go
const (
  // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
  O_RDONLY int = syscall.O_RDONLY // open the file read-only.
  O_WRONLY int = syscall.O_WRONLY // open the file write-only.
  O_RDWR int = syscall.O_RDWR // open the file read-write.
  // The reamining values may be or'ed in to control behavior.
  O_APPEND int = syscall.O_APPEND // append data to the file when writing.
  O_CREATE int = syscall.O_CREATE // create a new file if one exists.
  O_EXCL int = syscall.O_EXCL // used with O_CREATE, file must not exist.
  O_SYNC int = syscall.O_SYNC // open for synchronous I/O.
  O_TRUNC int = syscall.O_TRUNC // truncate regular writable file when opened.
)
```

# Kode Program Membuat File Baru
```go
func createNewFile(name string, message string) error {
  file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
  if err != nil {
    return err
  }

  defer file.Close()
  file.WriteString(message)
  return nil
}
```

# Kode Program Membaca File
```go
func readFile(name string) (string, error) {
  file, err := os.OpenFile(name, os.O_RDONLY, 0666)
  if err != nil {
    return "", err
  }

  defer file.Close()

  reader := bufio.NewReader(file)
  var message string
  for {
    line, _, err := reader.ReadLine()
    message += string(line)
    if err == io.EOF {
      break
    }
  }

  return message, nil
}
```

# Kode Program Membaca dan Menambah ke File
```go
func addToFile(name string, message string) error {
  file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
  if err != nil {
    return err
  }
  defer file.Close()
  file.WriteString(message)
  return nil
}
```