package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/attendeee/typer/model"
)

func GetStateFromJson() *model.State {

	var s model.State

	file, err := os.ReadFile("./state.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &s)

	return &s
}

func SaveStateToJson(s *model.State) {

	// Open (or create) the file with write permission, truncating it
	file, err := os.OpenFile("./state.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}

	defer file.Close()

	// Encode the struct to JSON and overwrite the file
	encoder := json.NewEncoder(file)

	err = encoder.Encode(s)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func ResizeByWidth(s []string, width int) []string {
	o := make([]string, len(s))
	for i := 0; i < len(s); i += 1 {
		o[i] = WrapText(s[i], width)
	}

	return o
}

func ConcatenateStrings(s []string) string {
	var t string

	for _, v := range s {
		t += v
		t += "\n"
		t += "\n"
	}

	t += "\n"

	return t
}

func WrapText(input string, width int) string {
	if width < 1 {
		return input
	}

	words := strings.Fields(input) // Split input into words
	var wrapped strings.Builder
	currentLineLength := 0

	for _, word := range words {

		if currentLineLength+len(word)+1 > width {
			wrapped.WriteString("\n")
			currentLineLength = 0
		} else if currentLineLength > 0 {
			wrapped.WriteString(" ")
			currentLineLength++
		}

		wrapped.WriteString(word)
		currentLineLength += len(word)
	}

	return wrapped.String()
}

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
