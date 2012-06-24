// global_scope
package main

import (
	"fmt"
)
var a = "G"

func main() {
	n()
	m()
}

func n(){fmt.Printf("%s\n",a)}
func m(){
	a = "HOJO"
	fmt.Printf("%s\n",a)
}