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
	form := huh.NewForm(
		huh.NewGroup(huh.NewMultiSelect[int]().
			Title("Todos").
			Options(
				options...,
			).
			Height(10).
			Value(&out)),
	)
	err := form.Run()
	if err != nil {
		slog.Error("something went wrong", "error", err)
		return
	}
	for _, i := range out {
		r.List[i].ToggleChecked(r.File)
	}
}
