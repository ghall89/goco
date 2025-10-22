package styles

import "github.com/charmbracelet/lipgloss"

var Success = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("32"))

var Error = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("9"))
