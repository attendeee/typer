package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/attendeee/typer/model"
	"github.com/attendeee/typer/ui"
	"github.com/attendeee/typer/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	path := flag.String("path", "", "Path to json file")
	chapter := flag.Int("c", 1, "Number of chapter")

	flag.Parse()

	book := utils.MustParseJsonToBook(*path)
	if *chapter-1 >= len(book.Chapters) {
		fmt.Println("Chapter is not present in a book")
		os.Exit(1)
	}

	s := utils.GetStateFromJson()
	if s == nil {
		s = &model.State{Chapter: uint32(*chapter - 1), CursorPos: 0}
	}

	m := ui.Model{Book: *book, State: *s}

	p := tea.NewProgram(&m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
