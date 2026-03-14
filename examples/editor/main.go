// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

// Editor example demonstrating vim-style editing with blockfont.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/timlinux/blockfont"
)

var (
	modeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#FFB347")).
			Bold(true).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))

	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF6B35")).
			Padding(1, 2)
)

type model struct {
	widget *blockfont.Widget
}

func initialModel() model {
	opts := blockfont.DefaultWidgetOptions()
	opts.Width = 70
	opts.Alignment = blockfont.AlignCenter
	opts.VimMode = true
	opts.Theme = blockfont.KartozaTheme

	widget := blockfont.NewWidget(opts)
	widget.SetText("edit me")
	widget.Focus()

	return model{
		widget: widget,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	w, cmd := m.widget.Update(msg)
	m.widget = w
	return m, cmd
}

func (m model) View() string {
	s := "\n"
	s += "  ╭─────────────────────────────────────────────────────────────────────────────╮\n"
	s += "  │                         blockfont Vim Editor Demo                          │\n"
	s += "  ╰─────────────────────────────────────────────────────────────────────────────╯\n"
	s += "\n"

	// Mode indicator
	mode := m.widget.Mode().String()
	s += "  " + modeStyle.Render(mode) + "\n\n"

	// Editor content
	content := borderStyle.Render(m.widget.View())
	s += content
	s += "\n\n"

	// Help
	var help string
	if m.widget.Mode() == blockfont.ModeNormal {
		help = "Normal: [i]nsert [a]ppend [x]delete [h/j/k/l]move"
	} else {
		help = "Insert: [esc]normal mode [type]insert text"
	}
	s += "  " + helpStyle.Render(help) + "\n"
	s += "  " + helpStyle.Render("[ctrl+c] Quit") + "\n"
	s += "\n"
	s += "  Made with ❤️ by Kartoza | https://kartoza.com\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
