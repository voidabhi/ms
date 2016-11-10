# ms
Millisecond conversion utility for golang

### Installation

```
$ go get github.com/voidabhi/ms
```

### Usage

```go
package main

import (
    "github.com/voidabhi/ms"
    "fmt"
)

func main() {
    fmt.Println(ms.Parse("10s")) // 10000
    fmt.Println(ms.Parse("3mins")) // 180000
    fmt.Println(ms.Parse("2h")) // 7.2e+06
}
```

