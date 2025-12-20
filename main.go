package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/attendeee/typer/model"
	"github.com/attendeee/typer/ui"

	tea "github.com/charmbracelet/bubbletea"
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

func main() {
	path := flag.String("path", "", "Path to json file")

	flag.Parse()

	book := MustParseJsonToBook(*path)

	p := tea.NewProgram(&ui.Model{Book: *book})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
