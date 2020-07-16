package books

// Dimensions holds the physical volume of a book
type Dimensions struct {
	Height   int
	Width    int
	Depth    int
	Filesize int
}

// Book is a thingy whatsitwhat is a book?
type Book struct {
	Title        string
	Author       string
	startingpage int
	endingpage   int
	Dimensions   Dimensions
	Booktype     bool // is it an ebook, yes.
	Colour       string
	BookLocation bookLocation
}

type bookLocation struct {
	// isbn
	// deweyDecimalNumber
	// onShelf - should be a boolean
	// available via ereader true or false
}