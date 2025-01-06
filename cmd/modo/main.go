package main

import (
	"log/slog"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/elliot40404/modo/internal/cli"
	"github.com/elliot40404/modo/internal/parser"
	"github.com/elliot40404/modo/internal/renderer"
)

func main() {
	f, err := cli.ParseArgs()
	if err != nil {
		slog.Error("something went wrong", "error", err)
		return
	}
	defer f.Close()
	todos, err := parser.ParseTodos(f)
	if err != nil {
		slog.Error("something went wrong", "error", err)
	}
	p := tea.NewProgram(renderer.InitialModel(todos, f))
	_, err = p.Run()
	if err != nil {
		slog.Error("something went wrong", "error", err)
		return
	}
}
