package main

import (
	"fmt"
	"book"
)


func (b Book) Volume() int {
	vol := b.dimensions.depth * b.dimensions.height * b.dimensions.depth
	return vol

}

func main() {

	newBook := Book{
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

	volumeBook := newBook.Volume
	fmt.Println("the book volume is":, volumeBook)

	pagesRead := newBook.Read()

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
