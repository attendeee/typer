package utils

import (
	"encoding/json"
	"os"

	"github.com/attendeee/typer/model"
)

func MustParseJsonToBook(path string) *model.Book {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var book model.Book

	err = json.Unmarshal(file, &book)
	if err != nil {
		panic(err)
	}

	return &book

}
