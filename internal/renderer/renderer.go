package renderer

import (
	"fmt"
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/elliot40404/modo/internal/parser"
)

type Renderer struct {
	List []parser.Todo
	File *os.File
}

type Model struct {
	Renderer
	cursor   int
	selected map[int]struct{}
}

func NewRenderer(list []parser.Todo, file *os.File) *Renderer {
	return &Renderer{
		List: list,
		File: file,
	}
}

func (r Renderer) Render() {
	selected := make(map[int]struct{})
	for i, todo := range r.List {
		if todo.Done {
			selected[i] = struct{}{}
		}
	}
	model := Model{
		Renderer: r,
		selected: selected,
	}
	p := tea.NewProgram(model, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		slog.Error("something went wrong", "error", err)
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.List)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			m.List[m.cursor].ToggleChecked(m.File)
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := ""
	for i, choice := range m.List {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.Content)
	}
	return s
}
