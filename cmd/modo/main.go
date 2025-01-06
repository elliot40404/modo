package main

import (
	"log/slog"

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
	if len(todos) == 0 {
		slog.Error("no todos found")
		return
	}
	renderer.NewRenderer(todos, f).Render()
}
