# names-generator
Generate random Nigerian names for testing purposes
## Usage

### Install Package

```bash
go get github.com/mmuoDev/names-generator
```

### Sample Usage

```go
package main

import (
	"log"

genNames "github.com/mmuoDev/names-generator/names"
)

func main() {
	names, err := genNames.GenerateRandomNames("yoruba", "male", 2)

	if err != nil {
		log.Fatal(err)
	}
	log.Print(names) // [mayowa ayotomiwa]
}
```
