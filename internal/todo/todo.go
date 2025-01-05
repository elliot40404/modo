package todo

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

type Todo struct {
	Line    string
	Done    bool
	Content string
	Offset  int64
}

func (t *Todo) ToggleChecked(f *os.File) {
	t.Done = !t.Done
	_, err := f.Seek(t.Offset, 0)
	if err != nil {
		slog.Error("something went wrong", "error", err)
		os.Exit(1)
	}
	if t.Done {
		t.Line = strings.Replace(t.Line, "[ ]", "[x]", 1)
	} else {
		t.Line = strings.Replace(t.Line, "[x]", "[ ]", 1)
	}
	_, err = f.WriteString(t.Line)
	if err != nil {
		slog.Error("something went wrong", "error", err)
		os.Exit(1)
	}
}

func ParseTodos(f *os.File) ([]Todo, error) {
	todos := make([]Todo, 0)
	scanner := bufio.NewScanner(f)
	offset := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		if todo, ok := ParseTodo(line, offset); ok {
			todos = append(todos, todo)
		}
		offset += int64(len(line) + 1)
	}
	return todos, nil
}

func ParseTodo(line string, offset int64) (Todo, bool) {
	ogLine := line
	line = strings.Trim(line, " ")
	if strings.HasPrefix(line, "- [ ]") || strings.HasPrefix(strings.ToLower(line), "- [x]") {
		todo := Todo{
			Line:    ogLine,
			Done:    strings.Contains(strings.ToLower(line), "[x]"),
			Content: strings.Split(line, "] ")[1],
			Offset:  offset,
		}
		return todo, true
	}
	return Todo{}, false
}
