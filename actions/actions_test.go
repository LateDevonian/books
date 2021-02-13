package actions

import (
	"fmt"
	"testing"

	"github.com/LateDevonian/books/definitions"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {

	janeEyre := buildBook()

	pagesRead := Read(janeEyre)
	expectedPages := 7
	fmt.Println(pagesRead)

	assert.Equal(t, expectedPages, pagesRead, "the correct number of pates are read")

}

func TestVolume(t *testing.T) {

	janeEyre := definitions.Book{
		Dimensions: definitions.Dimensions{
			Height: 2,
			Width:  3,
			Depth:  4,
		},
	}

	returnedVolume := Volume(janeEyre)
	expectedVolume := 24
	fmt.Println(returnedVolume)

	assert.Equal(t, expectedVolume, returnedVolume, "the book volume is calculated")
}

func buildBook() definitions.Book {

	exampleBook := definitions.Book{
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
	return exampleBook

}
