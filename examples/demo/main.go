// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

// Demo application showcasing all blockfont library features.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/timlinux/blockfont"
)

// Demo screens
type screen int

const (
	screenWelcome screen = iota
	screenLowercase
	screenUppercase
	screenBasicRender
	screenHighlights
	screenVimEditing
	screenAnimations
	screenWordWrap
	screenThemes
	screenCredits
)

var screenNames = []string{
	"Welcome",
	"Lowercase",
	"Uppercase",
	"Mixed",
	"Highlights",
	"Vim",
	"Animations",
	"Wrap",
	"Themes",
	"Credits",
}

// Styles
var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B35")).
			Bold(true).
			Align(lipgloss.Center)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFB347")).
			Italic(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF6B35")).
			Padding(1, 2)

	highlightStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFB347")).
			Bold(true)
)

type model struct {
	currentScreen screen
	width         int
	height        int

	// For vim editing demo
	vimWidget *blockfont.Widget

	// For animation demo
	animator      *blockfont.Animator
	animatedText  string
	animationDone bool

	// For carousel demo
	carousel     *blockfont.WordCarouselAnimator
	carouselWord int
	words        []string
}

func initialModel() model {
	// Create vim widget
	opts := blockfont.DefaultWidgetOptions()
	opts.Width = 60
	opts.VimMode = true
	opts.Theme = blockfont.KartozaTheme
	vimWidget := blockfont.NewWidget(opts)
	vimWidget.SetText("edit me")
	vimWidget.Focus()

	return model{
		currentScreen: screenWelcome,
		width:         80,
		height:        24,
		vimWidget:     vimWidget,
		animator:      blockfont.NewAnimator(),
		animatedText:  "fade",
		animationDone: true,
		carousel:      blockfont.NewWordCarouselAnimator(),
		carouselWord:  0,
		words:         []string{"block", "font", "demo", "cool"},
	}
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(blockfont.AnimationInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return tickCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "left", "h":
			if m.currentScreen > 0 {
				m.currentScreen--
			}

		case "right", "l":
			if m.currentScreen < screenCredits {
				m.currentScreen++
			}

		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			var idx int
			if msg.String() == "0" {
				idx = 9
			} else {
				idx = int(msg.Runes[0] - '1')
			}
			if idx >= 0 && idx <= int(screenCredits) {
				m.currentScreen = screen(idx)
			}

		case " ", "enter":
			// Trigger animation on animation screen
			if m.currentScreen == screenAnimations {
				m.animator.TriggerTransition(blockfont.TransitionFadeIn)
				m.animationDone = false
				cmds = append(cmds, tickCmd())
			}

		case "n":
			// Next word in carousel
			if m.currentScreen == screenAnimations {
				m.carouselWord = (m.carouselWord + 1) % len(m.words)
				m.carousel.TriggerTransition()
				cmds = append(cmds, tickCmd())
			}
		}

		// Handle vim editing on vim screen
		if m.currentScreen == screenVimEditing {
			w, cmd := m.vimWidget.Update(msg)
			m.vimWidget = w
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case tickMsg:
		// Update animations
		if m.currentScreen == screenAnimations {
			if m.animator.Update() {
				cmds = append(cmds, tickCmd())
			} else {
				m.animationDone = true
			}

			if m.carousel.Update() {
				cmds = append(cmds, tickCmd())
			}
		}

	case blockfont.AnimationTickMsg:
		if m.currentScreen == screenVimEditing {
			w, cmd := m.vimWidget.Update(msg)
			m.vimWidget = w
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var content string

	switch m.currentScreen {
	case screenWelcome:
		content = m.viewWelcome()
	case screenLowercase:
		content = m.viewLowercase()
	case screenUppercase:
		content = m.viewUppercase()
	case screenBasicRender:
		content = m.viewBasicRender()
	case screenHighlights:
		content = m.viewHighlights()
	case screenVimEditing:
		content = m.viewVimEditing()
	case screenAnimations:
		content = m.viewAnimations()
	case screenWordWrap:
		content = m.viewWordWrap()
	case screenThemes:
		content = m.viewThemes()
	case screenCredits:
		content = m.viewCredits()
	}

	// Build navigation
	nav := m.buildNavigation()

	// Build help footer
	help := helpStyle.Render("[←/→] Navigate • [1-9,0] Jump to screen • [q] Quit")

	return fmt.Sprintf("%s\n\n%s\n\n%s", nav, content, help)
}

func (m model) buildNavigation() string {
	var tabs []string
	for i, name := range screenNames {
		style := lipgloss.NewStyle().Padding(0, 1)
		if screen(i) == m.currentScreen {
			style = style.
				Foreground(lipgloss.Color("#000000")).
				Background(lipgloss.Color("#FF6B35")).
				Bold(true)
		} else {
			style = style.Foreground(lipgloss.Color("#888888"))
		}
		tabs = append(tabs, style.Render(fmt.Sprintf("%d:%s", i+1, name)))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func (m model) viewWelcome() string {
	// Render "blockfont" in block letters
	title := blockfont.RenderText("blockfont")

	// Center it
	lines := strings.Split(title, "\n")
	centered := blockfont.CenterLines(lines, m.width)
	title = strings.Join(centered, "\n")

	subtitle := subtitleStyle.Render("Unicode block letter rendering for terminal applications")
	subtitle = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, subtitle)

	features := []string{
		"• Full ASCII character set (a-z, A-Z, 0-9, punctuation)",
		"• Vim-style text editing with cursor support",
		"• Character-level highlighting for typing games",
		"• Spring-based animations with harmonica",
		"• Word wrapping with cursor tracking",
		"• Lipgloss integration for styling",
		"• Bubbletea Widget (tea.Model)",
	}

	featureBox := boxStyle.Render(strings.Join(features, "\n"))
	featureBox = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, featureBox)

	return fmt.Sprintf("%s\n\n%s\n\n%s", title, subtitle, featureBox)
}

func (m model) viewLowercase() string {
	s := titleStyle.Width(m.width).Render("Lowercase Alphabet")
	s += "\n\n"

	s += subtitleStyle.Render("a - m:") + "\n"
	s += blockfont.RenderText("abcdefghijklm") + "\n\n"

	s += subtitleStyle.Render("n - z:") + "\n"
	s += blockfont.RenderText("nopqrstuvwxyz") + "\n"

	return s
}

func (m model) viewUppercase() string {
	s := titleStyle.Width(m.width).Render("Uppercase Alphabet")
	s += "\n\n"

	s += subtitleStyle.Render("A - M:") + "\n"
	s += blockfont.RenderText("ABCDEFGHIJKLM") + "\n\n"

	s += subtitleStyle.Render("N - Z:") + "\n"
	s += blockfont.RenderText("NOPQRSTUVWXYZ") + "\n"

	return s
}

func (m model) viewBasicRender() string {
	s := titleStyle.Width(m.width).Render("Mixed Characters")
	s += "\n\n"

	// Show mixed case
	s += subtitleStyle.Render("Mixed case:") + "\n"
	s += blockfont.RenderText("Hello World") + "\n\n"

	// Show numbers
	s += subtitleStyle.Render("Numbers:") + "\n"
	s += blockfont.RenderText("0123456789") + "\n\n"

	// Show punctuation
	s += subtitleStyle.Render("Punctuation:") + "\n"
	s += blockfont.RenderText("!?.,;:-'\"") + "\n"

	return s
}

func (m model) viewHighlights() string {
	s := titleStyle.Width(m.width).Render("Character Highlights")
	s += "\n\n"
	s += subtitleStyle.Render("Perfect for typing games and diffs:") + "\n\n"

	// Demo text with highlights
	text := "typing"
	highlights := []blockfont.CharHighlight{
		blockfont.HighlightCorrect,   // t - green
		blockfont.HighlightCorrect,   // y - green
		blockfont.HighlightIncorrect, // p - red
		blockfont.HighlightPending,   // i - gray
		blockfont.HighlightPending,   // n - gray
		blockfont.HighlightPending,   // g - gray
	}

	lines := blockfont.RenderWithCursor(text, -1, highlights, false, 0, blockfont.DefaultTheme)
	s += strings.Join(lines, "\n") + "\n\n"

	// Legend
	legend := []string{
		blockfont.ANSIGreen + "██" + blockfont.ANSIReset + " Correct",
		blockfont.ANSIRed + "██" + blockfont.ANSIReset + " Incorrect",
		blockfont.ANSIDim + "██" + blockfont.ANSIReset + " Pending",
		blockfont.ANSICyan + "██" + blockfont.ANSIReset + " Cursor",
	}
	s += boxStyle.Render(strings.Join(legend, "  "))

	return s
}

func (m model) viewVimEditing() string {
	s := titleStyle.Width(m.width).Render("Vim-Style Editing")
	s += "\n\n"

	// Mode indicator
	mode := m.vimWidget.Mode().String()
	modeStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#FFB347")).
		Bold(true).
		Padding(0, 1)
	s += modeStyle.Render(mode) + "\n\n"

	// The editor widget
	s += boxStyle.Render(m.vimWidget.View()) + "\n\n"

	// Help based on mode
	if m.vimWidget.Mode() == blockfont.ModeNormal {
		s += helpStyle.Render("[i] Insert • [a] Append • [x] Delete • [h/j/k/l] Move • [0/$] Line start/end")
	} else {
		s += helpStyle.Render("[Esc] Normal mode • Type to insert text")
	}

	return s
}

func (m model) viewAnimations() string {
	s := titleStyle.Width(m.width).Render("Spring Animations")
	s += "\n\n"

	// Animated text with fade
	s += subtitleStyle.Render("Fade animation (press Space to trigger):") + "\n"

	opacity := m.animator.GetOpacityLevel(1.0)
	var animText string
	if opacity > 0.7 {
		animText = blockfont.RenderText(m.animatedText)
	} else if opacity > 0.3 {
		animText = blockfont.WrapWithColor(blockfont.RenderText(m.animatedText), blockfont.ANSIDim)
	} else {
		animText = "" // Faded out
	}
	s += animText + "\n\n"

	// Word carousel
	s += subtitleStyle.Render("Word carousel (press N for next):") + "\n"

	// Current word with position based on animation
	currentIdx := m.carouselWord
	prevIdx := (currentIdx - 1 + len(m.words)) % len(m.words)
	nextIdx := (currentIdx + 1) % len(m.words)

	// Previous word (dimmed)
	prevOpacity := m.carousel.GetPrevOpacity()
	if prevOpacity > 0.2 {
		prevText := blockfont.WrapWithColor(blockfont.RenderText(m.words[prevIdx]), blockfont.ANSIDim)
		s += prevText + "\n"
	}

	// Current word (bright)
	currentText := blockfont.RenderText(m.words[currentIdx])
	s += currentText + "\n"

	// Next word (dimmed)
	nextOpacity := m.carousel.GetNextOpacity()
	if nextOpacity > 0.2 {
		nextText := blockfont.WrapWithColor(blockfont.RenderText(m.words[nextIdx]), blockfont.ANSIDim)
		s += nextText + "\n"
	}

	return s
}

func (m model) viewWordWrap() string {
	s := titleStyle.Width(m.width).Render("Word Wrapping")
	s += "\n\n"
	s += subtitleStyle.Render("Text wraps at word boundaries with cursor tracking:") + "\n\n"

	// Demo word wrap with cursor
	text := "hello world demo"
	lines := blockfont.RenderWithCursor(text, 6, nil, false, 50, blockfont.DefaultTheme)
	s += boxStyle.Render(strings.Join(lines, "\n"))

	return s
}

func (m model) viewThemes() string {
	s := titleStyle.Width(m.width).Render("Theme Support")
	s += "\n\n"

	// Default theme
	s += subtitleStyle.Render("Default Theme:") + "\n"
	defaultHighlights := []blockfont.CharHighlight{
		blockfont.HighlightCorrect,
		blockfont.HighlightIncorrect,
		blockfont.HighlightPending,
	}
	lines := blockfont.RenderWithCursor("abc", -1, defaultHighlights, false, 0, blockfont.DefaultTheme)
	s += strings.Join(lines, "\n") + "\n\n"

	// Kartoza theme
	s += subtitleStyle.Render("Kartoza Theme:") + "\n"
	lines = blockfont.RenderWithCursor("abc", -1, defaultHighlights, false, 0, blockfont.KartozaTheme)
	s += strings.Join(lines, "\n") + "\n\n"

	// Color options
	s += subtitleStyle.Render("ANSI Color Constants:") + "\n"
	colors := []struct {
		name  string
		color string
	}{
		{"Red", blockfont.ANSIRed},
		{"Green", blockfont.ANSIGreen},
		{"Orange", blockfont.ANSIOrange},
		{"Blue", blockfont.ANSIBlue},
		{"Cyan", blockfont.ANSICyan},
		{"Magenta", blockfont.ANSIMagenta},
	}

	var colorDemo []string
	for _, c := range colors {
		colorDemo = append(colorDemo, c.color+"██"+blockfont.ANSIReset+" "+c.name)
	}
	s += boxStyle.Render(strings.Join(colorDemo, "  "))

	return s
}

func (m model) viewCredits() string {
	// Render "thanks" in block letters
	title := blockfont.RenderText("thanks")
	lines := strings.Split(title, "\n")
	centered := blockfont.CenterLines(lines, m.width)
	title = strings.Join(centered, "\n")

	credits := []string{
		"blockfont v0.1.0",
		"",
		"A Unicode block letter rendering library",
		"for terminal applications.",
		"",
		"Built with:",
		"  • Go",
		"  • charmbracelet/bubbletea",
		"  • charmbracelet/lipgloss",
		"  • charmbracelet/harmonica",
		"",
		"https://github.com/timlinux/blockfont",
	}

	creditsBox := boxStyle.Render(strings.Join(credits, "\n"))
	creditsBox = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, creditsBox)

	footer := highlightStyle.Render("Made with ❤️ by Kartoza")
	footer = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, footer)

	return fmt.Sprintf("%s\n\n%s\n\n%s", title, creditsBox, footer)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
