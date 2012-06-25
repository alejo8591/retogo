// strings_split_join
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Lastima que las cosas sean así"
	sl := strings.Fields(str)
	fmt.Printf("Splitter en un slice %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s -", val)
	}
	fmt.Println()
	str2 := "GO|ABC de GO|232323"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitter en un slice %v\n",sl2)
	for _, val := range sl2{
		fmt.Printf("%s -", val)
	}
	fmt.Println()
	
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined con ; quedaria: %s\n", str3)
}