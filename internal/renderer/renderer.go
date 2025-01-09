package renderer

import (
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
	for i, todo := range r.List {
		options = append(options, huh.NewOption(todo.Content, i).Selected(todo.Done))
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
		if err == huh.ErrUserAborted {
			return
		}
		slog.Error("something went wrong", "error", err)
		return
	}
	for _, i := range out {
		todo := r.List[i]
		if !todo.Done {
			r.List[i].ToggleChecked(r.File)
		}
	}
}
