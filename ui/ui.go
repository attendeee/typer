package ui

import (
	"time"

	"github.com/attendeee/typer/model"
	"github.com/attendeee/typer/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Book    model.Book
	Chapter int
}

func (m Model) Init() tea.Cmd {
	utils.ResizeByWidth(&m.Book.Chapters[0].Text, 80)

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

	case tea.WindowSizeMsg:
		time.Sleep(50 * time.Millisecond) // Maybe there is a better solution with channels //
		utils.ResizeByWidth(&m.Book.Chapters[0].Text, msg.Width)

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
