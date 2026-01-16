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
	highlight = color.New(color.BgRed, color.FgWhite, color.Bold).SprintFunc()
	written   = color.FgGreen
	unwritten = color.New(color.FgHiWhite).SprintFunc()
)

type Model struct {
	Book    model.Book
	Chapter int

	Text      string
	CursorPos int

	Pager Pager
}

type Pager struct {
	OffsetStep      int
	Offsets         []int
	UpperOffset     int
	UpperOffsetIdx  int
	BottomOffset    int
	BottomOffsetIdx int
}

func (m *Model) Init() tea.Cmd {

	tmp := utils.ResizeByWidth(m.Book.Chapters[m.Chapter].Text, 80)

	m.Text = utils.ConcatenateStrings(tmp)

	UpdateOffsets(m, &m.Pager)

	m.Pager.OffsetStep = 20

	m.CursorPos = 0

	m.Pager.UpperOffsetIdx = 0

	UpdateUpperOffsetIdx(m, &m.Pager)
	UpdateBottomOffsetIdx(&m.Pager)

	m.Pager.UpperOffset = m.Pager.Offsets[m.Pager.UpperOffsetIdx]
	m.Pager.BottomOffset = m.Pager.Offsets[m.Pager.BottomOffsetIdx]

	// Just return `nil`, which means "no I/O right now, please." //
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		time.Sleep(100 * time.Millisecond)

		m.Pager.OffsetStep = int(float32(msg.Height) * 0.75)
		tmp := utils.ResizeByWidth(m.Book.Chapters[m.Chapter].Text, int(float32(msg.Width)*0.75))

		m.Text = utils.ConcatenateStrings(tmp)

		UpdateOffsets(m, &m.Pager)

		UpdateUpperOffsetIdx(m, &m.Pager)
		UpdateBottomOffsetIdx(&m.Pager)

		m.Pager.UpperOffset = m.Pager.Offsets[m.Pager.UpperOffsetIdx]
		m.Pager.BottomOffset = m.Pager.Offsets[m.Pager.BottomOffsetIdx]

		return m, tea.ClearScreen

	// Check if key press event //
	case tea.KeyMsg:
		switch msg.String() {

		// These keys should exit the program //
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.CursorPos+2 > len(m.Text) {
				return m, tea.Quit
			}

			if m.Text[m.CursorPos] == '\n' {
				m.CursorPos += 1
			}

			if m.CursorPos+1 > m.Pager.BottomOffset {
				ScrollDown(&m.Pager)

			}

			return m, nil

		case "backspace":
			if m.CursorPos > 0 {
				m.CursorPos -= 1
			}

			if m.CursorPos < m.Pager.UpperOffset {
				ScrollUp(&m.Pager)

			}

			return m, nil

		default:
			if msg.String()[0] == m.Text[m.CursorPos] {
				m.CursorPos += 1
			}

			if m.CursorPos+1 >= len(m.Text) {
				return m, tea.Quit
			}
		}

	}

	return m, nil
}

func (m *Model) View() string {

	color.Set(written)

	return fmt.Sprintf("%s%s%s",
		m.Text[m.Pager.UpperOffset:m.CursorPos],
		highlight(m.Text[m.CursorPos:m.CursorPos+1]),
		unwritten(m.Text[m.CursorPos+1:m.Pager.BottomOffset]))
}
