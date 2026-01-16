# Package reflect
- Dalam bahasa pemrograman, biasanya ada fitur reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan.
- Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect.
- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan.
- [https://golang.org/pkg/reflect/](https://golang.org/pkg/reflect/).

# Kode Program Package reflect
```go
type Sample struct {
  Name string
}

func main() {
  sample := Sample{"Sandy"}
  sampleType := reflect.TypeOf(sample)
  structField := sampleType.Field(0)

  fmt.Println(structField.Name)
}
```

# Kode Program StructTag
```go
type Sample struct {
  Name string `required:"true" max:"10"`
}

func main() {
  sample := Sample{"Sandy"}
  sampleType := reflect.TypeOf(sample)
  structField := sampleType.Field(0)
  required := structField.Tag.Get("required")

  fmt.Println(required)
}
```

# Kode Program Validation Library
```go
func IsValid(data interface{}) bool {
  t := reflect.TypeOf(data)
  for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    if field.Tag.Get("required") == "true" {
      return reflect.ValueOf(data).Field(i).Interface() != ""
    }
  }

  return true
}
```