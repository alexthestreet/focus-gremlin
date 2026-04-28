package prompt

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	choices  []string
	selected int
	done     bool
	choice   string
}

func NewModel(choices []string) Model {
	model := Model{choices: append([]string(nil), choices...)}
	if len(model.choices) > 0 {
		model.choice = model.choices[0]
	}

	return model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown:
			if m.selected < len(m.choices)-1 {
				m.selected++
				m.choice = m.choices[m.selected]
			}
		case tea.KeyUp:
			if m.selected > 0 {
				m.selected--
				m.choice = m.choices[m.selected]
			}
		case tea.KeyEnter:
			m.done = true
		}
		if len(m.choices) > 0 {
			m.choice = m.choices[m.selected]
		}
	}

	return m, nil
}

func (m Model) SelectedIndex() int {
	return m.selected
}

func (m Model) SelectedChoice() string {
	return m.choice
}

func (m Model) Done() bool {
	return m.done
}
