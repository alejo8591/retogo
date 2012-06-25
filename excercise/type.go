// type
package main

import (
	"fmt"
)

type ENT int
func main() {
	var a,b ENT = 3, 4
	c:= a + b 
	fmt.Printf("el valor de c es: %d", c)
}
