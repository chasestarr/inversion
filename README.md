## [Inversion](https://en.wikipedia.org/wiki/Inversion_(discrete_mathematics))

[![GoDoc](https://godoc.org/github.com/chasestarr/inversion?status.svg)](https://godoc.org/github.com/chasestarr/inversion)

```go
import (
  "fmt"

  "github.com/chasestarr/inversion"
)

func main() {
  count := inversion.Count([]int{6, 5, 4, 3, 2, 1})
  fmt.Println(count) // 15
  // # of inversions in reversed array are: (n(n-1))/2
}
```