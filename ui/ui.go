package ui

import (
	"github.com/attendeee/typer/model"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Book model.Book
}

func (m Model) Init() tea.Cmd {

	// Just return `nil`, which means "no I/O right now, please." //
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Check if key press event //
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program //
		case "ctrl+c":
			return m, tea.Quit

		}
	}

	return m, nil
}

func (m Model) View() string {
	var s string

	for _, v := range m.Book.Chapters[0].Text {
		s += v
		s += "\n"
	}

	s += "\n"
	s += "Press or 'ctrl+c'\n\n"

	return s
}
