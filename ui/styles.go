package ui

import "github.com/charmbracelet/lipgloss"

var (
	SelectColor     = lipgloss.Color("#a983f7")
	DescSelectColor = lipgloss.Color("#8d6dcf")
)

var (
	// Audio Category
	BaseSelectStyle = lipgloss.NewStyle().
			Bold(true).
			BorderStyle(lipgloss.BlockBorder()).
			BorderForeground(SelectColor).
			BorderLeft(true).
			PaddingLeft(1)

	SelectStyle     = BaseSelectStyle.Foreground(SelectColor)
	DescSelectStyle = BaseSelectStyle.Foreground(DescSelectColor)
	UnselectStyle   = lipgloss.NewStyle().PaddingLeft(2)

	// Audio Destination
	TextBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1)
)
