package renderer

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/elliot40404/modo/internal/todo"
)

type Model struct {
	fd       *os.File
	choices  []todo.Todo
	cursor   int
	selected map[int]struct{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel(todos []todo.Todo, fd *os.File) Model {
	selected := make(map[int]struct{})
	for i, todo := range todos {
		if todo.Done {
			selected[i] = struct{}{}
		}
	}
	return Model{
		choices:  todos,
		selected: selected,
		fd:       fd,
	}
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
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			m.choices[m.cursor].ToggleChecked(m.fd)
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
	for i, choice := range m.choices {
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
