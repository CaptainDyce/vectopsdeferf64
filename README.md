# vectopsf64
Vector operation ([]float64) library -  specifications of deferred operations on vectors

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
	op := v.Op().Ident().Divl(2 * math.Pi).ApplyOp(math.Sin).Rev()
	a := op.OnSize(100) // apply it on a new [100]float64
	fmt.Println(a)

	b := make([]float64, 100)
	op.On(b) // apply it on b
	fmt.Println(b)
}
```
