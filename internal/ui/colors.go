package ui

import "github.com/charmbracelet/lipgloss"

var (
	DocStyle   = lipgloss.NewStyle().Margin(1, 2)
	BrandColor = lipgloss.Color("#00D7FF")
	PillColor  = lipgloss.Color("#FF00D7")

	TitleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Background(BrandColor).Padding(0, 1).Bold(true)

	CommandStyle = lipgloss.NewStyle().Foreground(BrandColor).Bold(true)
	PillStyle    = lipgloss.NewStyle().Foreground(PillColor).Italic(true)
	DescStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#88888"))
)
