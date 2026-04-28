package prompt

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle    = lipgloss.NewStyle().Bold(true)
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
)

func (m Model) View() string {
	lines := []string{titleStyle.Render("Are you actually working right now?"), ""}
	for i, choice := range m.choices {
		prefix := "  "
		if i == m.selected {
			prefix = "> "
			choice = selectedStyle.Render(choice)
		}
		lines = append(lines, prefix+choice)
	}

	return strings.Join(lines, "\n")
}
