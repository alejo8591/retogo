// converting
package main

import (
	"fmt"
	"math"
)
func main() {
	fmt.Println(Uint8FromInt(30))
	fmt.Println(IntFromFloat64(30.4))
}
func Uint8FromInt(n int) (uint8, error){
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the range int32 range", x))
}