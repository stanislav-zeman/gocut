package gocut

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	appStyle   = lipgloss.NewStyle().Padding(1, 2)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)
	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)

type Model struct {
	list         list.Model
	delegateKeys *delegateKeyMap
}

func NewModel() Model {
	delegateKeys := newDelegateKeyMap()
	commands := []list.Item{
		Command{
			Name:    "xd",
			Command: "xpanes {1..4}",
		},
		Command{
			Name:    "help",
			Command: "lvim help.go",
		},
	}

	delegate := newItemDelegate(delegateKeys)
	commandList := list.New(commands, delegate, 0, 0)
	commandList.Title = "Commands"
	commandList.Styles.Title = titleStyle
	return Model{
		list:         commandList,
		delegateKeys: delegateKeys,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}
	}

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel

	var cmds []tea.Cmd
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return appStyle.Render(m.list.View())
}
