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
	// In the handler we first need to make a Context 
	// object associated with the current request r
	c := appengine.NewContext(r)
	// Then from this context we test whether there is already a user 
	// logged in at this point with
	u := user.Current(c)
	// If this is the case user.Current returns a pointer 
	// to a user.User value for the user; otherwise it returns nil.
	if u == nil{
		// he 2nd parameter r.URL.String() is 
		// the currently requested url so that the Google account sign-in
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			// The LoginURL() function returns an error value as its 2nd argument. 
			http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	} 
	fmt.Fprint(w, "Hello, %v !", u)
}