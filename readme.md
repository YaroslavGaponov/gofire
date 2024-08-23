Fire in terminal
======

# Run

```sh
go run samples/fire.go
```

# Use

```golang
package main

import (
	"github.com/YaroslavGaponov/gofire"
	"time"
)

func main() {
	fire := gofire.NewFire(50, 30)
	for {
		fire.Draw()
		time.Sleep(100 * time.Millisecond)
	}
}
```
