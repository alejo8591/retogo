// local_scope
package main

import (
	"fmt"
)
var a = "G"
func main() {
	n()
	m()
	o()
}
func n(){fmt.Printf(a)}

func m(){fmt.Printf("Chau Soda")}

func o(){
	a := "o"
	fmt.Printf(a)
}