package actions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CatlogColors struct {
	CatlogColors []Catlog `json:"catlog_colors"`
}

type Catlog struct {
	//Colors string `json: "colors"`
}

func open_colors() {
	file, _ := ioutil.ReadFile("colours.json")

	data := CatlogColors{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.CatlogColors); i++ {
		fmt.Println("Product Id: ", data.CatlogColors[i].Colors)
	}
}
