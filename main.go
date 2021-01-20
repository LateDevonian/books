package main

import (
	"fmt"
)

func main() {
	//Describes a book in detail as an example of using golang

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
	fmt.Println("the book volume is", volumeBook)

	pagesRead := newBook.Read()

	if pagesRead > 0 {
		fmt.Println("Hurrah, you have read this many pages:", pagesRead)
	} else {
		err := fmt.Errorf("oh dear you didn't read any pages today")
		fmt.Println(err)
	}
}
