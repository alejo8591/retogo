// my_os
package main

import (
	"fmt"
	"os"
)
var nombre string
var prompt = "Enter a digit, e.g. 3 "+  "or %s to quit."
func init() {
  if runtime.GOOS == "windows" {
          prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
  } else { // Unix-like
          prompt = fmt.Sprintf(prompt, "Ctrl+D")
  } 
}