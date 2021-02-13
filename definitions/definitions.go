package definitions

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
	Startingpage int
	Endingpage   int
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

type Colours struct {
	// colours string `json:"colours"`
	// [
	//   {
	//     colour string `json:"colour"`
	//     category string `json:"category"`
	//     colortype string `json:"type"`
	// 	code string `json:"code"`
	// 		{
	//       	rgba []int `json:"rgba"`
	//       	hex string `json:"hex"`
	// 		}
	// 	}
	// ]
}
