// helloapp
package helloapp

import (
	"fmt"
	"net/http"
)

func init() {
         http.HandleFunc(“/”, handle)
}
func handle(w http.ResponseWriter, r *http.Request) {
	// some Chinese characters after World!
	fmt.Fprint(w, “<html><body>Hello, World! 세상아 안녕!! </body></html>”)
}