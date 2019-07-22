# vectopsf64
Vector ([]float64) library - immediate operations on float64 arrays

### Quick Example

```go
package main

import (
	"fmt"
	v "github.com/CaptainDyce/vectopsdeferf64"
	"math"
)

func main() {
	// defining the following operation:
	// identity vector, divide each entry by 2*pi, apply math.Sin, and reverse the vector
	op := v.Ident().Divl(2 * math.Pi).ApplyOp(math.Sin).Rev()
	a := op.OnSize(100) // []float64
	fmt.Println(a)

	b := make([]float64, 100)
	op.On(b) // apply the previously defined operation to b
	fmt.Println(b)
}
```
