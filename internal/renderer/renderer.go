package renderer

import (
	"errors"
	"log/slog"
	"os"

	"github.com/charmbracelet/huh"

	"github.com/elliot40404/modo/internal/parser"
)

type Renderer struct {
	List []parser.Todo
	File *os.File
}

func NewRenderer(list []parser.Todo, file *os.File) *Renderer {
	return &Renderer{
		List: list,
		File: file,
	}
}

func (r Renderer) Render() {
	options := make([]huh.Option[int], 0, len(r.List))
	preselected := make([]int, 0, len(r.List))
	for i, todo := range r.List {
		options = append(options, huh.NewOption(todo.Content, i).Selected(todo.Done))
		if todo.Done {
			preselected = append(preselected, i)
		}
	}
	out := make([]int, 0, len(r.List))
	keymap := huh.NewDefaultKeyMap()
	keymap.Quit.SetKeys("q", "ctrl+c")
	err := huh.NewForm(
		huh.NewGroup(huh.NewMultiSelect[int]().
			Title("Todos").
			Options(
				options...,
			).
			Height(10).
			Value(&out)),
	).
		WithKeyMap(keymap).
		Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			return
		}
		slog.Error("something went wrong", "error", err)
		return
	}
	done := make(map[int]struct{})
	for _, i := range out {
		done[i] = struct{}{}
		todo := r.List[i]
		if !todo.Done {
			todo.ToggleChecked(r.File)
		}
	}
	for _, i := range preselected {
		if _, ok := done[i]; !ok {
			r.List[i].ToggleChecked(r.File)
		}
	}
}
