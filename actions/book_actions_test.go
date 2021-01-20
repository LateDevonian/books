package actions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {

	janeEyre := Book{
		Title:  "Jane Eyre",
		Author: "Charlotte Bronte",
		Dimensions: Dimensions{
			height: 34,
			width:  44,
			depth:  1,
		},
		Booktype:     false,
		startingpage: 1,
		endingpage:   8,
		Colour:       "Penguin orange classic style",
	}

	pagesRead := janeEyre.Read()
	expectedPages := 7
	fmt.Println(pagesRead)

	assert.Equal(t, expectedPages, pagesRead, "the correct number of pates are read")

}

func TestVolume(t *testing.T) {

	janeEyre := Book{
		Dimensions: Dimensions{
			height: 2,
			width:  3,
			depth:  4,
		},
	}

	returnedVolume := janeEyre.Volume()
	expectedVolume := 24
	fmt.Println(returnedVolume)

	assert.Equal(t, expectedVolume, returnedVolume, "the book volume is calculated")
}
