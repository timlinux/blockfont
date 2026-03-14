// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

// Animated example demonstrating blockfont widget with animations.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/timlinux/blockfont"
)

type model struct {
	widget  *blockfont.Widget
	words   []string
	current int
}

func initialModel() model {
	opts := blockfont.DefaultWidgetOptions()
	opts.Width = 80
	opts.Alignment = blockfont.AlignCenter
	opts.Animate = true

	widget := blockfont.NewWidget(opts)

	words := []string{"hello", "world", "blockfont", "demo"}
	widget.SetText(words[0])

	return model{
		widget:  widget,
		words:   words,
		current: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case " ", "enter", "n":
			// Next word
			m.current = (m.current + 1) % len(m.words)
			m.widget.SetText(m.words[m.current])
			return m, m.widget.TriggerAnimation(blockfont.TransitionFadeIn)
		case "p":
			// Previous word
			m.current = (m.current - 1 + len(m.words)) % len(m.words)
			m.widget.SetText(m.words[m.current])
			return m, m.widget.TriggerAnimation(blockfont.TransitionSlideUp)
		}
	}

	w, cmd := m.widget.Update(msg)
	m.widget = w
	return m, cmd
}

func (m model) View() string {
	s := "\n"
	s += "  ╭─────────────────────────────────────────────────────────────────────────────╮\n"
	s += "  │                        blockfont Animation Demo                            │\n"
	s += "  ╰─────────────────────────────────────────────────────────────────────────────╯\n"
	s += "\n"
	s += m.widget.View()
	s += "\n\n"
	s += "  [space/n] Next word • [p] Previous word • [q] Quit\n"
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
