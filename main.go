package main

import (
	"fmt"
)

//Book struct tutorial

type dimensions struct {
	height int
	width  int
	depth  int
}

// Book is a thingy whatsitwhat is a book?
type Book struct {
	title        string
	author       string
	startingpage int
	endingpage   int
	dimensions   dimensions
	booktype     bool // is it an ebook, yes
	colour       string
	bookLocation bookLocation
}

type bookLocation struct {

	// isbn
	// deweyDecimalNumber
	// onShelf - should be a boolean
	// available via ereader true or false
}

func (b Book) Read() int {
	pagesRead := b.endingpage - b.startingpage
	return pagesRead
}

func (b Book) Volume() int {
	vol := b.dimensions.depth * b.dimensions.height * b.dimensions.depth
	return vol

}

func main() {

	janeEyre := Book{
		title:  "Jane Eyre",
		author: "Charlotte Bronte",
		dimensions: dimensions{
			height: 34,
			width:  44,
			depth:  1,
		},
		booktype:     false,
		startingpage: 1,
		endingpage:   8,
		colour:       "Penguin orange classic style",
	}

	volumeJane := janeEyre.Volume
	fmt.Println(volumeJane)
	fmt.Println(janeEyre.Volume)

	pagesRead := janeEyre.Read()
	fmt.Println(pagesRead)

	if pagesRead > 0 {
		fmt.Println("Hurrah, you have read this many pages:", pagesRead)
	} else {
		err := fmt.Errorf("oh dear you didn't read any pages today")
		fmt.Println(err)
	}
}

func (b Book) Locate() string {

	return "this is where your return will go"
}
