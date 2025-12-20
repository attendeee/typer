package ui

import (
	"fmt"
	"time"

	"github.com/attendeee/typer/model"
	"github.com/attendeee/typer/utils"
	"github.com/fatih/color"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	highlight = color.New(color.BgRed, color.Underline).SprintFunc()
	written   = color.New(color.FgHiWhite).SprintFunc()
	unwritten = color.New(color.FgWhite).SprintFunc()
)

type Model struct {
	Book      model.Book
	Chapter   int
	Text      string
	CursorPos int
}

func (m *Model) Init() tea.Cmd {
	m.CursorPos = 1990

	utils.ResizeByWidth(&m.Book.Chapters[m.Chapter].Text, 120)

	m.Text = utils.ConcatenateStrings(&m.Book.Chapters[m.Chapter].Text)

	// Just return `nil`, which means "no I/O right now, please." //
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Check if key press event //
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program //
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.Text[m.CursorPos] == '\n' {
				m.CursorPos += 1
			}

			if m.CursorPos+1 >= len(m.Text) {
				return m, tea.Quit
			}

		case "backspace":
			if m.CursorPos > 0 {
				m.CursorPos -= 1
			}

		default:
			if msg.String()[0] == m.Text[m.CursorPos] {
				m.CursorPos += 1
			}

			if m.CursorPos+1 >= len(m.Text) {
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		time.Sleep(50 * time.Millisecond) // Maybe there is a better solution with channels //
		utils.ResizeByWidth(&m.Book.Chapters[m.Chapter].Text, msg.Width)

	}

	return m, nil
}

func (m *Model) View() string {
	return fmt.Sprintf("%s%s%s", written(m.Text[:m.CursorPos]), highlight(m.Text[m.CursorPos:m.CursorPos+1]), unwritten(m.Text[m.CursorPos+1:]))
}
