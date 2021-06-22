# names-generator
Generate Nigerian names for testing purposes. This was written in Go
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
	numOfNames := 2
	names, err := genNames.GenerateRandomNames("yoruba", "male", numOfNames)

	if err != nil {
		log.Fatal(err)
	}
	log.Print(names) // [mayowa ayotomiwa]
}
```
