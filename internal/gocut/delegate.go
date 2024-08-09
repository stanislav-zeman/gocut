package gocut

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()
	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		i, ok := m.SelectedItem().(Command)
		if !ok {
			return nil
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				args := strings.Split(i.Command, " ")
				cmd := exec.Command(args[0], args[1:]...)
				return tea.ExecProcess(cmd, func(err error) tea.Msg {
					return err
				})
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose}
	d.ShortHelpFunc = func() []key.Binding {
		return help
	}
	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "execute"),
		),
	}
}
