package actions

func (b Book) Read() int {
	pagesRead := b.endingpage - b.startingpage
	return pagesRead
}

func (b Book) Locate() string {

	return "this is where your return will go"
}

func (b Book) Volume() int {
	vol := b.Dimensions.depth * b.Dimensions.height * b.Dimensions.width
	return vol

}
