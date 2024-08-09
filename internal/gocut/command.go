package gocut

import "github.com/charmbracelet/bubbles/list"

var _ list.DefaultItem = Command{}

type Command struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
}

func (c Command) Title() string {
	return c.Name
}

func (c Command) Description() string {
	return c.Command
}

func (c Command) FilterValue() string {
	return c.Name
}
