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
	screenCharacter
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
	"Character",
	"Wrap",
	"Themes",
	"Credits",
}

// Kartoza brand colors
var (
	kartozaOrange = lipgloss.Color("#FF6B35")
	kartozaGold   = lipgloss.Color("#FFB347")
)

// Styles
var (
	headerStyle = lipgloss.NewStyle().
			Foreground(kartozaOrange).
			Bold(true).
			Align(lipgloss.Center)

	straplineStyle = lipgloss.NewStyle().
			Foreground(kartozaGold).
			Italic(true).
			Align(lipgloss.Center)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Align(lipgloss.Center)

	titleStyle = lipgloss.NewStyle().
			Foreground(kartozaOrange).
			Bold(true).
			Align(lipgloss.Center)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(kartozaGold).
			Italic(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#666666"))

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(kartozaOrange).
			Padding(1, 2)

	footerStyle = lipgloss.NewStyle().
			Foreground(kartozaGold).
			Bold(true).
			Align(lipgloss.Center)

	warningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true).
			Align(lipgloss.Center)
)

// Popular color scheme themes
var (
	// Catppuccin Mocha
	CatppuccinTheme = blockfont.Theme{
		Correct:   "\033[38;2;166;227;161m", // Green
		Incorrect: "\033[38;2;243;139;168m", // Red
		Cursor:    "\033[38;2;137;180;250m", // Blue
		Pending:   "\033[38;2;108;112;134m", // Overlay0
	}

	// Dracula
	DraculaTheme = blockfont.Theme{
		Correct:   "\033[38;2;80;250;123m",  // Green
		Incorrect: "\033[38;2;255;85;85m",   // Red
		Cursor:    "\033[38;2;139;233;253m", // Cyan
		Pending:   "\033[38;2;98;114;164m",  // Comment
	}

	// Nord
	NordTheme = blockfont.Theme{
		Correct:   "\033[38;2;163;190;140m", // Green
		Incorrect: "\033[38;2;191;97;106m",  // Red
		Cursor:    "\033[38;2;136;192;208m", // Frost
		Pending:   "\033[38;2;76;86;106m",   // Polar Night
	}

	// Gruvbox
	GruvboxTheme = blockfont.Theme{
		Correct:   "\033[38;2;184;187;38m",  // Green
		Incorrect: "\033[38;2;251;73;52m",   // Red
		Cursor:    "\033[38;2;131;165;152m", // Aqua
		Pending:   "\033[38;2;146;131;116m", // Gray
	}

	// Tokyo Night
	TokyoNightTheme = blockfont.Theme{
		Correct:   "\033[38;2;158;206;106m", // Green
		Incorrect: "\033[38;2;247;118;142m", // Red
		Cursor:    "\033[38;2;125;207;255m", // Cyan
		Pending:   "\033[38;2;86;95;137m",   // Comment
	}
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

	// For character animation demo
	characterAnimator *blockfont.CharacterAnimator
	characterAction   blockfont.AnimationAction
	characterFlipped  bool

	// Status message
	statusMsg string
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
		currentScreen:     screenWelcome,
		width:             80,
		height:            24,
		vimWidget:         vimWidget,
		animator:          blockfont.NewAnimator(),
		animatedText:      "fade",
		animationDone:     true,
		carousel:          blockfont.NewWordCarouselAnimator(),
		carouselWord:      0,
		words:             []string{"block", "font", "demo", "cool"},
		characterAnimator: blockfont.NewCharacterAnimator(),
		characterAction:   blockfont.ActionIdle,
		characterFlipped:  false,
		statusMsg:         "Welcome to blockfont demo",
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

func (m model) isVimInsertMode() bool {
	return m.currentScreen == screenVimEditing && m.vimWidget.Mode() != blockfont.ModeNormal
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		// If in vim insert mode, only handle escape and pass everything else to vim
		if m.isVimInsertMode() {
			// Let vim widget handle the input
			w, cmd := m.vimWidget.Update(msg)
			m.vimWidget = w
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
			return m, tea.Batch(cmds...)
		}

		// Normal navigation mode
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "left", "h":
			if m.currentScreen > 0 {
				m.currentScreen--
				m.statusMsg = fmt.Sprintf("Screen: %s", screenNames[m.currentScreen])
			}

		case "right", "l":
			if m.currentScreen < screenCredits {
				m.currentScreen++
				m.statusMsg = fmt.Sprintf("Screen: %s", screenNames[m.currentScreen])
			}

		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0":
			var idx int
			if msg.String() == "0" {
				idx = 9 // 0 = screen 10 (Credits)
			} else {
				idx = int(msg.Runes[0] - '1')
			}
			if idx >= 0 && idx <= int(screenCredits) {
				m.currentScreen = screen(idx)
				m.statusMsg = fmt.Sprintf("Screen: %s", screenNames[m.currentScreen])
			}

		case " ", "enter":
			// Trigger animation on animation screen
			if m.currentScreen == screenAnimations {
				m.animator.TriggerTransition(blockfont.TransitionFadeIn)
				m.animationDone = false
				m.statusMsg = "Animation triggered"
				cmds = append(cmds, tickCmd())
			}

		case "n":
			// Next word in carousel
			if m.currentScreen == screenAnimations {
				m.carouselWord = (m.carouselWord + 1) % len(m.words)
				m.carousel.TriggerTransition()
				m.statusMsg = fmt.Sprintf("Word: %s", m.words[m.carouselWord])
				cmds = append(cmds, tickCmd())
			}
			// Next action in character screen
			if m.currentScreen == screenCharacter {
				actions := blockfont.AllActions()
				for i, a := range actions {
					if a == m.characterAction {
						m.characterAction = actions[(i+1)%len(actions)]
						break
					}
				}
				m.characterAnimator.SetAction(m.characterAction)
				m.statusMsg = fmt.Sprintf("Action: %s", blockfont.GetActionName(m.characterAction))
			}

		case "p":
			// Previous action in character screen
			if m.currentScreen == screenCharacter {
				actions := blockfont.AllActions()
				for i, a := range actions {
					if a == m.characterAction {
						idx := (i - 1 + len(actions)) % len(actions)
						m.characterAction = actions[idx]
						break
					}
				}
				m.characterAnimator.SetAction(m.characterAction)
				m.statusMsg = fmt.Sprintf("Action: %s", blockfont.GetActionName(m.characterAction))
			}

		case "f":
			// Flip character direction
			if m.currentScreen == screenCharacter {
				m.characterFlipped = !m.characterFlipped
				m.characterAnimator.SetFlipped(m.characterFlipped)
				dir := "right"
				if m.characterFlipped {
					dir = "left"
				}
				m.statusMsg = fmt.Sprintf("Facing: %s", dir)
			}
		}

		// Handle vim editing on vim screen (normal mode keys)
		if m.currentScreen == screenVimEditing {
			w, cmd := m.vimWidget.Update(msg)
			m.vimWidget = w
			if cmd != nil {
				cmds = append(cmds, cmd)
			}
		}

	case tickMsg:
		// Always keep tick running for screens that need animation
		cmds = append(cmds, tickCmd())

		// Update animations
		if m.currentScreen == screenAnimations {
			if m.animator.Update() {
				// animation still running
			} else {
				m.animationDone = true
			}
			m.carousel.Update()
		}

		// Update character animation
		if m.currentScreen == screenCharacter {
			m.characterAnimator.Update()
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
	// Calculate available height for content
	headerHeight := 3  // title + strapline + status
	footerHeight := 3  // nav + help + branding
	contentHeight := m.height - headerHeight - footerHeight - 2 // -2 for spacing
	if contentHeight < 10 {
		contentHeight = 10
	}

	// Build header (always at top, centered)
	header := m.buildHeader()

	// Build content
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
	case screenCharacter:
		content = m.viewCharacter()
	case screenWordWrap:
		content = m.viewWordWrap()
	case screenThemes:
		content = m.viewThemes()
	case screenCredits:
		content = m.viewCredits()
	}

	// Center content both horizontally and vertically
	content = m.centerContent(content, contentHeight)

	// Build footer (always at bottom, centered)
	footer := m.buildFooter()

	return fmt.Sprintf("%s\n%s\n%s", header, content, footer)
}

// centerContent centers the content both horizontally and vertically within the available space
func (m model) centerContent(content string, availableHeight int) string {
	lines := strings.Split(content, "\n")

	// Center each line horizontally
	var centeredLines []string
	for _, line := range lines {
		centered := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, line)
		centeredLines = append(centeredLines, centered)
	}

	// Calculate vertical padding
	contentHeight := len(centeredLines)
	if contentHeight < availableHeight {
		topPadding := (availableHeight - contentHeight) / 2
		bottomPadding := availableHeight - contentHeight - topPadding

		// Add top padding
		paddedLines := make([]string, 0, availableHeight)
		for range topPadding {
			paddedLines = append(paddedLines, "")
		}

		// Add content
		paddedLines = append(paddedLines, centeredLines...)

		// Add bottom padding
		for range bottomPadding {
			paddedLines = append(paddedLines, "")
		}

		return strings.Join(paddedLines, "\n")
	}

	return strings.Join(centeredLines, "\n")
}

func (m model) buildHeader() string {
	// Title
	title := headerStyle.Width(m.width).Render("blockfont")

	// Strapline
	strapline := straplineStyle.Width(m.width).Render("Block letter rendering using ASCII block characters")

	// Status
	status := statusStyle.Width(m.width).Render(m.statusMsg)

	return fmt.Sprintf("%s\n%s\n%s", title, strapline, status)
}

func (m model) buildFooter() string {
	// Navigation tabs
	nav := m.buildNavigation()
	nav = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, nav)

	// Help text
	var helpText string
	if m.isVimInsertMode() {
		helpText = warningStyle.Width(m.width).Render("-- INSERT MODE -- Press [Esc] to return to navigation")
	} else {
		helpText = helpStyle.Width(m.width).Render("[←/→] Navigate • [1-9,0] Jump • [q] Quit")
	}

	// Kartoza branding
	branding := footerStyle.Width(m.width).Render("Made with ❤️ by Kartoza | https://kartoza.com")

	return fmt.Sprintf("\n%s\n%s\n%s", nav, helpText, branding)
}

func (m model) buildNavigation() string {
	var tabs []string
	for i, name := range screenNames {
		style := lipgloss.NewStyle().Padding(0, 1)

		// Display key: 1-9 for first 9, 0 for 10th
		key := fmt.Sprintf("%d", i+1)
		if i == 9 {
			key = "0"
		}

		if screen(i) == m.currentScreen {
			style = style.
				Foreground(lipgloss.Color("#000000")).
				Background(kartozaOrange).
				Bold(true)
		} else {
			style = style.Foreground(lipgloss.Color("#888888"))
		}
		tabs = append(tabs, style.Render(fmt.Sprintf("%s:%s", key, name)))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, tabs...)
}

func (m model) viewWelcome() string {
	// Render "blockfont" in block letters
	title := blockfont.RenderText("blockfont")

	features := []string{
		"• Lowercase letters (a-z)",
		"• Uppercase letters (A-Z)",
		"• Numbers (0-9)",
		"• Common punctuation (. , ; : ! ? - ' \" etc.)",
		"• Vim-style text editing with cursor support",
		"• Character-level highlighting for typing games",
		"• Spring-based animations with harmonica",
		"• Word wrapping with cursor tracking",
		"• Lipgloss integration for styling",
		"• Bubbletea Widget (tea.Model)",
	}

	featureBox := boxStyle.Render(strings.Join(features, "\n"))

	return fmt.Sprintf("%s\n\n%s", title, featureBox)
}

func (m model) viewLowercase() string {
	s := titleStyle.Render("Lowercase Alphabet")
	s += "\n\n"

	s += subtitleStyle.Render("a - m:") + "\n"
	s += blockfont.RenderText("abcdefghijklm") + "\n\n"

	s += subtitleStyle.Render("n - z:") + "\n"
	s += blockfont.RenderText("nopqrstuvwxyz")

	return s
}

func (m model) viewUppercase() string {
	s := titleStyle.Render("Uppercase Alphabet")
	s += "\n\n"

	s += subtitleStyle.Render("A - M:") + "\n"
	s += blockfont.RenderText("ABCDEFGHIJKLM") + "\n\n"

	s += subtitleStyle.Render("N - Z:") + "\n"
	s += blockfont.RenderText("NOPQRSTUVWXYZ")

	return s
}

func (m model) viewBasicRender() string {
	s := titleStyle.Render("Mixed Characters")
	s += "\n\n"

	// Show mixed case
	s += subtitleStyle.Render("Mixed case:") + "\n"
	s += blockfont.RenderText("Hello World") + "\n\n"

	// Show numbers
	s += subtitleStyle.Render("Numbers:") + "\n"
	s += blockfont.RenderText("0123456789") + "\n\n"

	// Show punctuation
	s += subtitleStyle.Render("Punctuation:") + "\n"
	s += blockfont.RenderText("!?.,;:-'\"")

	return s
}

func (m model) viewHighlights() string {
	s := titleStyle.Render("Character Highlights")
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
	s := titleStyle.Render("Vim-Style Editing")
	s += "\n\n"

	// Mode indicator
	mode := m.vimWidget.Mode().String()
	var modeStyle lipgloss.Style
	if m.vimWidget.Mode() == blockfont.ModeNormal {
		modeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#98c379")).
			Bold(true).
			Padding(0, 1)
	} else {
		modeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(lipgloss.Color("#e5c07b")).
			Bold(true).
			Padding(0, 1)
	}
	s += modeStyle.Render(mode) + "\n\n"

	// The editor widget
	s += boxStyle.Render(m.vimWidget.View()) + "\n\n"

	// Help based on mode
	if m.vimWidget.Mode() == blockfont.ModeNormal {
		s += helpStyle.Render("[i] Insert • [a] Append • [x] Delete • [h/l] Move • [0/$] Start/End")
	} else {
		s += warningStyle.Render("Type to insert • [Esc] Return to normal mode")
	}

	return s
}

func (m model) viewAnimations() string {
	s := titleStyle.Render("Spring Animations")
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
		s += nextText
	}

	return s
}

func (m model) viewCharacter() string {
	s := titleStyle.Render("Character Animation")
	s += "\n\n"
	s += subtitleStyle.Render(fmt.Sprintf("Action: %s", blockfont.GetActionName(m.characterAction))) + "\n\n"

	// Render the character
	charColor := "\033[38;2;255;107;53m" // Kartoza orange
	s += m.characterAnimator.RenderWithColor(charColor) + "\n\n"

	// Action buttons
	actions := blockfont.AllActions()
	var actionButtons []string
	for _, action := range actions {
		name := blockfont.GetActionName(action)
		style := lipgloss.NewStyle().Padding(0, 1)
		if action == m.characterAction {
			style = style.
				Foreground(lipgloss.Color("#000000")).
				Background(kartozaOrange).
				Bold(true)
		} else {
			style = style.Foreground(lipgloss.Color("#888888"))
		}
		actionButtons = append(actionButtons, style.Render(name))
	}
	s += lipgloss.JoinHorizontal(lipgloss.Top, actionButtons...) + "\n\n"

	// Direction indicator
	dir := "→ Right"
	if m.characterFlipped {
		dir = "← Left"
	}
	dirStyle := lipgloss.NewStyle().
		Foreground(kartozaGold).
		Bold(true)
	s += dirStyle.Render("Direction: "+dir) + "\n\n"

	// Help text
	s += helpStyle.Render("[N] Next action • [P] Previous action • [F] Flip direction")

	return s
}

func (m model) viewWordWrap() string {
	s := titleStyle.Render("Word Wrapping")
	s += "\n\n"
	s += subtitleStyle.Render("Text wraps at word boundaries with cursor tracking:") + "\n\n"

	// Demo word wrap with cursor
	text := "hello world demo"
	lines := blockfont.RenderWithCursor(text, 6, nil, false, 50, blockfont.DefaultTheme)
	s += boxStyle.Render(strings.Join(lines, "\n"))

	return s
}

func (m model) viewThemes() string {
	s := titleStyle.Render("Popular Color Schemes")
	s += "\n\n"

	// Demo highlights for all themes
	demoHighlights := []blockfont.CharHighlight{
		blockfont.HighlightCorrect,
		blockfont.HighlightIncorrect,
		blockfont.HighlightPending,
	}

	themes := []struct {
		name  string
		theme blockfont.Theme
	}{
		{"Kartoza (Default)", blockfont.KartozaTheme},
		{"Catppuccin Mocha", CatppuccinTheme},
		{"Dracula", DraculaTheme},
		{"Nord", NordTheme},
		{"Gruvbox", GruvboxTheme},
		{"Tokyo Night", TokyoNightTheme},
	}

	for i, t := range themes {
		s += subtitleStyle.Render(t.name+":") + "\n"
		lines := blockfont.RenderWithCursor("abc", 1, demoHighlights, false, 0, t.theme)
		s += strings.Join(lines, "\n")
		if i < len(themes)-1 {
			s += "\n\n"
		}
	}

	return s
}

func (m model) viewCredits() string {
	// Render "thanks" in block letters
	title := blockfont.RenderText("thanks")

	credits := []string{
		"blockfont v0.1.0",
		"",
		"A block letter rendering library",
		"for terminal applications.",
		"",
		"Created by:",
		"  Tim Sutton (@timlinux)",
		"  https://github.com/timlinux",
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

	return fmt.Sprintf("%s\n\n%s", title, creditsBox)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
