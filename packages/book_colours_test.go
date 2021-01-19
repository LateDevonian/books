package books

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Book_Colour_Extracted(t *testing.T) {

	fmt.Println("test one")
	expected := "one"
	actual := "two"

	assert.Equal(t, expected, actual, "testing the test")

}

func Test_Test(t *testing.T) {

	fmt.Println("test one")
	expected := "one"
	actual := "two"

	assert.Equal(t, expected, actual, "testing the test")

}
