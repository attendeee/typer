package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/attendeee/typer/ui"
	"github.com/attendeee/typer/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	path := flag.String("path", "", "Path to json file")

	flag.Parse()

	book := utils.MustParseJsonToBook(*path)

	p := tea.NewProgram(&ui.Model{Book: *book})

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
