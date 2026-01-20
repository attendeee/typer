package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/adrg/xdg"
	"github.com/attendeee/typer/model"
	"github.com/attendeee/typer/ui"
	"github.com/attendeee/typer/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	if _, err := os.Stat(xdg.DataHome + "/typer/"); os.IsNotExist(err) {
		err := os.Mkdir(xdg.DataHome+"/typer/", 0755)
		if err != nil {
			fmt.Println("Unable to create typer folder", err)
			time.Sleep(5 * time.Second)
		}

	}

	path := flag.String("path", "", "Path to json file")
	state := flag.Bool("s", false, "Use saved state")
	chapter := flag.Int("c", 1, "Number of chapter")

	flag.Parse()

	book := utils.MustParseJsonToBook(*path)
	if *chapter-1 >= len(book.Chapters) {
		fmt.Println("Chapter is not present in a book")
		os.Exit(1)
	}

	s := &model.State{Chapter: uint32(*chapter - 1), CursorPos: 0}

	if *state == true {
		s = utils.GetStateFromJson(*path)
	}

	m := ui.Model{Book: *book, Path: *path, State: *s}

	p := tea.NewProgram(&m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
