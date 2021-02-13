package actions

import (
	"github.com/LateDevonian/books/definitions"
)

func Read(b definitions.Book) int {
	pagesRead := b.Endingpage - b.Startingpage

	return pagesRead
}

func Locate(b definitions.Book) string {

	return "this is where your return will go"
}

func Volume(b definitions.Book) int {
	vol := b.Dimensions.Depth * b.Dimensions.Height * b.Dimensions.Width
	return vol

}
