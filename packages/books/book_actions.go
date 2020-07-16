package books

func (b Book) Read() int {
	pagesRead := b.endingpage - b.startingpage
	return pagesRead
}

func (b Book) Locate() string {

	return "this is where your return will go"
}
