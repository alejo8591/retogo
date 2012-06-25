// ----- PLAY WITH STRINGS -----------
package main
import (
        "fmt"
	    "strings"
)
func main() {
	var str string = "Hola mocomo como van las cosas. Hola!"
	fmt.Printf("La frase es: Hola mocomo como van las cosas. Hola!")
	fmt.Printf("La posición de \"mocomo\" es: ")
	fmt.Printf("%d\n", strings.Index(str,"mocomo"))
	fmt.Printf("La posición de la primera instancia de \"Hola\" es: ") 
	fmt.Printf("%d\n", strings.Index(str, "Hola"))
	fmt.Printf("La posición de la ultima instancia de \"Hola\" es: ") 
	fmt.Printf("%d\n", strings.LastIndex(str, "Hola")) 
	fmt.Printf("La posición del \"Perro\" es: ") 
	fmt.Printf("%d\n", strings.Index(str, "Perro"))
	
	// second part count chars
	var str1 string = "hola me voy pa donde yopo?"
    var manyG = "ppppppppp"
    fmt.Printf("Numero de O`s %s es: ", str1) 
    fmt.Printf("%d\n", strings.Count(str, "o"))
    fmt.Printf("números de p’s in %s ie: ", manyG)
    fmt.Printf("%d\n", strings.Count(manyG, "gg"))

    // third part repeat strings
    var origins string = "hola mundillo"
    var repeated string
    repeated = strings.Repeat(origins, 3)
    fmt.Printf("El nuevo string repetido es: %s\n", repeated)
￼￼￼}