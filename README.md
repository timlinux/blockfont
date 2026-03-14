# blockfont

Unicode block letter rendering for terminal applications.

```
██████ ██     ◢████◣ ◢████◣ ██ ◢█◤ ██████ ◢████◣ ██◣ ██ ◥█  █◤
██  ██ ██     ██  ██ ██  ██ ██◢█◤  ██     ██  ██ ███◣██  ◥██◤
██████ ██     ██  ██ ██     ███◣   ████   ██  ██ ██◥███   ██
██  ██ ██     ██  ██ ██  ◢█ ██◥█◣  ██     ██  ██ ██ ◥██   ██
██████ ██████ ◥████◤ ◥████◤ ██ ◥█◣ ██     ◥████◤ ██  ██   ██
```

## Features

- **Block Letter Rendering**: Convert text to beautiful Unicode block letters with smooth rounded corners
- **Vim-Style Editing**: Full vim buffer with normal/insert modes and common operations
- **Spring Animations**: Smooth physics-based transitions using harmonica
- **Lipgloss Integration**: Full charmbracelet/lipgloss styling support
- **Bubbletea Widget**: Ready-to-use `tea.Model` for easy integration
- **Character Highlighting**: Color individual characters for typing games, diffs, etc.
- **Layout Utilities**: Centering, alignment, and word wrapping

## Installation

```bash
go get github.com/timlinux/blockfont
```

## Quick Start

### Simple Rendering

```go
package main

import (
    "fmt"
    "github.com/timlinux/blockfont"
)

func main() {
    // Render text as block letters
    text := blockfont.RenderText("hello")
    fmt.Println(text)
}
```

### Widget with Bubbletea

```go
package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/timlinux/blockfont"
)

type model struct {
    widget *blockfont.Widget
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    w, cmd := m.widget.Update(msg)
    m.widget = w
    return m, cmd
}

func (m model) View() string {
    return m.widget.View()
}

func main() {
    opts := blockfont.DefaultWidgetOptions()
    opts.Width = 80
    opts.Alignment = blockfont.AlignCenter
    opts.VimMode = true

    widget := blockfont.NewWidget(opts)
    widget.SetText("hello world")

    p := tea.NewProgram(model{widget: widget})
    p.Run()
}
```

### Character Highlighting

```go
widget := blockfont.NewWidget(blockfont.DefaultWidgetOptions())
widget.SetText("typing")

// Set highlights for a typing game
highlights := []blockfont.CharHighlight{
    blockfont.HighlightCorrect,   // 't' - correct
    blockfont.HighlightCorrect,   // 'y' - correct
    blockfont.HighlightIncorrect, // 'p' - wrong
    blockfont.HighlightPending,   // 'i' - not typed yet
    blockfont.HighlightPending,   // 'n' - not typed yet
    blockfont.HighlightPending,   // 'g' - not typed yet
}
widget.SetHighlights(highlights)
```

## API Reference

### Core Functions

```go
// Render a word as block letters (returns 6 lines of character arrays)
func RenderWord(word string) [][]string

// Render text as a single string with newlines
func RenderText(text string) string

// Get the width of a single character
func GetLetterWidth(char rune) int

// Get the total width of a word
func GetTotalWidth(word string) int
```

### Widget

```go
// Create a new widget with options
func NewWidget(opts WidgetOptions) *Widget

// Set the text content
func (w *Widget) SetText(text string)

// Get the current text
func (w *Widget) Text() string

// Set character highlights
func (w *Widget) SetHighlights(highlights []CharHighlight)

// Render the widget (implements tea.Model.View)
func (w *Widget) View() string
```

### Buffer (Vim Mode)

```go
// Create a new buffer
func NewBuffer(text string) *Buffer

// Text manipulation
func (b *Buffer) Insert(text string)
func (b *Buffer) Delete(n int) string
func (b *Buffer) DeleteLine() string

// Cursor movement
func (b *Buffer) MoveLeft(n int)
func (b *Buffer) MoveRight(n int)
func (b *Buffer) MoveUp(n int)
func (b *Buffer) MoveDown(n int)

// Mode control
func (b *Buffer) Mode() Mode
func (b *Buffer) SetMode(mode Mode)
```

### Layout

```go
// Center lines within a width
func CenterLines(lines []string, width int) []string

// Left-justify with margin
func LeftJustify(lines []string, margin int) []string

// Right-justify within width
func RightJustify(lines []string, width int) []string

// Word wrap text
func WrapOnWordBoundaries(text string, maxWidth int) []string
```

### Animation

```go
// Create a new animator
func NewAnimator() *Animator

// Trigger a transition
func (a *Animator) TriggerTransition(t TransitionType)

// Update animation (call each frame)
func (a *Animator) Update() bool

// Get animation values
func (a *Animator) GetOffset(maxOffset int) int
func (a *Animator) GetOpacityLevel(maxOpacity float64) float64
func (a *Animator) GetScaleFactor(minScale float64) float64
```

## Supported Characters

- Lowercase letters: a-z
- Uppercase letters: A-Z
- Digits: 0-9
- Punctuation: `. , ; : ! ? - ' " ( ) [ ] { } < > / \ | _ + = * & @ # $ % ^ ` ~`
- Space

## Themes

```go
// Use default theme
widget.SetTheme(blockfont.DefaultTheme)

// Use Kartoza-branded theme
widget.SetTheme(blockfont.KartozaTheme)

// Create custom theme
customTheme := blockfont.Theme{
    Primary:   lipgloss.Color("#FFFFFF"),
    Correct:   lipgloss.Color("#00FF00"),
    Incorrect: lipgloss.Color("#FF0000"),
    // ... etc
}
widget.SetTheme(customTheme)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
