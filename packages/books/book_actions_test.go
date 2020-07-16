package books

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
			Height: 34,
			Width:  44,
			Depth:  1,
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
