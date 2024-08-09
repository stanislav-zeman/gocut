package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stanislav-zeman/gocut/internal/gocut"
)

func main() {
	model := gocut.NewModel()
	program := tea.NewProgram(model, tea.WithAltScreen())
	_, err := program.Run()
	if err != nil {
		fmt.Println("error running gocut:", err)
		os.Exit(1)
	}
}
