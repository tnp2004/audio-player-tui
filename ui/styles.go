package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Color Code
	LightPurple = "#a983f7"
	DarkPurple  = "#8d6dcf"
)

var (
	// Audio Category
	BaseSelectStyle = lipgloss.NewStyle().
			Bold(true).
			BorderStyle(lipgloss.BlockBorder()).
			BorderForeground(lipgloss.Color(LightPurple)).
			BorderLeft(true).
			PaddingLeft(1)

	SelectStyle     = BaseSelectStyle.Foreground(lipgloss.Color(LightPurple))
	DescSelectStyle = BaseSelectStyle.Foreground(lipgloss.Color(DarkPurple))
	UnselectStyle   = lipgloss.NewStyle().PaddingLeft(2)

	// Audio Destination
	TextBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(LightPurple)).
			Padding(0, 1)

	// Audio Player
	ProgressBarStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(LightPurple))
	ProgressBarBoxStyle = lipgloss.NewStyle().MarginTop(1)
)
