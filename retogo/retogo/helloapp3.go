// helloapp3
package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct{
	Title string
	Body []byte
}

// This is a method named save that takes as its receiver p, a pointer to Page . 
// It takes no parameters, and returns a value of type error.
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// 0600, passed as the third parameter to WriteFile, 
	// indicates that the file should be created with read-write 
	// permissions for the current user only.
	return ioutil.WriteFile(filename, p.Body, 0600)
}
// The function loadPage constructs the file name from Title, 
// reads the file's contents into a new Page, and returns a pointer to that new page.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body,err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
		return &Page{Title: title, Body: body}, nil
}

func main(){
	p1 := &Page{Title: "TestPage", Body: []byte("This page is example")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}