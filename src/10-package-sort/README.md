# Package sort
- Package sort adalah package yang berisikan utilitas untuk proses pengurutan.
- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface ```sort.Interface```.
- [https://golang.org/pkg/sort/](https://golang.org/pkg/sort/).

# Kode Program sort.Interface
```go
// elements of the collection be enumerated by an integer index.
type Interface interface {
  // Len is the number of elements in the collection.
  Len() int
  // Less reports whether the element with
  // index i should sort before the element with index j.
  Less(i, j, int) bool
  // Swap swaps the elements with indexes i and j.
  Swap(i, j int)
}
```

# Kode Program Package sort (1)
```go
type User struct {
  Name  string
  Age   int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
  return len(userSlice)
}
```

# Kode Program Package sort (2)
```go
func (userSlice UserSlice) Len() int {
  return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
  return userSlice[i].Age < userSlice[j].Age
} 

func (userSlice UserSlice) Swap(i, j int) {
  userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}
```

# Kode Program Package sort (3)
```go
func main() {
  users := []User{
    {"Sandy", 30},
    {"Dwi", 25},
    {"Budi", 28},
  }

  sort.Sort(UserSlice(users))
  fmt.Println(users)
}
```