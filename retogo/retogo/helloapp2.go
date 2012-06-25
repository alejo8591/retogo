// helloapp2
package helloapp2

import (
	"fmt"
	"appengine"
	"appengine/user"
	"net/http"
)
func init(){
	http.HandleFunc("/", handler)
}
// User API
	// loggin with User API with google gmail
func handler(w http.ResponseWriter, r *http.Request){
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil{
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	} 
	fmt.Fprint(w, "Hello, %v!", u)
}