package main

import (
	"fmt"

	"github.com/LateDevonian/books/actions"
	"github.com/LateDevonian/books/definitions"
)

func main() {
	//Describes a book in detail as an example of using golang

	newBook := definitions.Book{
		Title:  "Jane Eyre",
		Author: "Charlotte Bronte",
		Dimensions: definitions.Dimensions{
			Height: 34,
			Width:  44,
			Depth:  1,
		},
		Booktype:     false,
		Startingpage: 1,
		Endingpage:   8,
		Colour:       "Penguin orange classic style",
	}

	volumeBook := actions.Volume(newBook)
	fmt.Println("the book volume is", volumeBook)

	pagesRead := actions.Read(newBook)

	if pagesRead > 0 {
		fmt.Println("Hurrah, you have read this many pages:", pagesRead)
	} else {
		err := fmt.Errorf("oh dear you didn't read any pages today")
		fmt.Println(err)
	}
}
