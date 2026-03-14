# Widget Integration

The Widget provides a high-level interface for bubbletea applications. It wraps all blockfont functionality into a ready-to-use `tea.Model` component.

## Creating a Widget

### With Default Options

```go
opts := blockfont.DefaultWidgetOptions()
widget := blockfont.NewWidget(opts)
widget.SetText("hello")
```

### With Custom Options

```go
opts := blockfont.WidgetOptions{
    Width:       80,
    Alignment:   blockfont.AlignCenter,
    VimMode:     true,
    WordWrap:    true,
    Theme:       blockfont.KartozaTheme,
    CursorStyle: blockfont.CursorBlock,
}

widget := blockfont.NewWidget(opts)
widget.SetText("edit me")
widget.Focus()  // Enable keyboard input
```

## Widget Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `Width` | `int` | 80 | Maximum render width in characters |
| `Height` | `int` | 0 | Target render height (0 = auto) |
| `Alignment` | `Alignment` | `AlignLeft` | Text alignment (Left, Center, Right) |
| `VimMode` | `bool` | `false` | Enable vim-style editing |
| `Animate` | `bool` | `false` | Enable spring animations |
| `WordWrap` | `bool` | `false` | Enable word-boundary wrapping |
| `CursorStyle` | `CursorStyle` | `CursorBlock` | Cursor appearance |
| `Theme` | `Theme` | `DefaultTheme` | Color theme |

### Alignment Options

```go
blockfont.AlignLeft    // Left-align text (default)
blockfont.AlignCenter  // Center text
blockfont.AlignRight   // Right-align text
```

### Cursor Styles

```go
blockfont.CursorBlock      // Block cursor (solid rectangle)
blockfont.CursorLine       // Vertical line cursor
blockfont.CursorUnderline  // Underline cursor
```

## Bubbletea Integration

### Basic Example

```go
package main

import (
    "fmt"
    "os"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/timlinux/blockfont"
)

type model struct {
    widget *blockfont.Widget
}

func initialModel() model {
    opts := blockfont.DefaultWidgetOptions()
    opts.VimMode = true

    widget := blockfont.NewWidget(opts)
    widget.SetText("hello")
    widget.Focus()

    return model{widget: widget}
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
    }

    // Pass messages to widget
    w, cmd := m.widget.Update(msg)
    m.widget = w
    return m, cmd
}

func (m model) View() string {
    mode := m.widget.Mode().String()
    return fmt.Sprintf(
        "Mode: %s\n\n%s\n\nPress Ctrl+C to quit",
        mode, m.widget.View(),
    )
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

### With Animations

```go
package main

import (
    "time"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/timlinux/blockfont"
)

type model struct {
    widget *blockfont.Widget
}

// Animation tick message
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
    case tickMsg:
        // Update animations
        w, cmd := m.widget.Update(blockfont.AnimationTickMsg(time.Time(msg)))
        m.widget = w
        if cmd != nil {
            cmds = append(cmds, cmd)
        }
        cmds = append(cmds, tickCmd())

    case tea.KeyMsg:
        if msg.String() == "ctrl+c" {
            return m, tea.Quit
        }
        w, cmd := m.widget.Update(msg)
        m.widget = w
        if cmd != nil {
            cmds = append(cmds, cmd)
        }
    }

    return m, tea.Batch(cmds...)
}

func (m model) View() string {
    return m.widget.View()
}
```

## Widget Methods

### Text Management

```go
// Set text content
widget.SetText("new text")

// Get current text
text := widget.Text()
```

### Focus Management

```go
// Give keyboard focus
widget.Focus()

// Remove focus
widget.Blur()

// Check focus state
if widget.Focused() {
    // Widget has focus
}
```

### Mode and Cursor

```go
// Get current vim mode
mode := widget.Mode()  // ModeNormal, ModeInsert, etc.

// Get cursor position
pos := widget.CursorPosition()
```

## Character Highlighting

Perfect for typing games, code diffs, and visual feedback:

```go
// Define highlights for each character
highlights := []blockfont.CharHighlight{
    blockfont.HighlightCorrect,    // Green
    blockfont.HighlightCorrect,    // Green
    blockfont.HighlightIncorrect,  // Red
    blockfont.HighlightPending,    // Dim gray
    blockfont.HighlightPending,    // Dim gray
}

// Apply highlights
widget.SetHighlights(highlights)
```

### Highlight Types

| Highlight | Color | Use Case |
|-----------|-------|----------|
| `HighlightNone` | Default | No highlighting |
| `HighlightCorrect` | Green | Correct input |
| `HighlightIncorrect` | Red | Wrong input |
| `HighlightCursor` | Cyan (inverted) | Cursor position |
| `HighlightPending` | Dim | Not yet typed |
| `HighlightDelete` | Theme delete color | Deleted text |
| `HighlightChange` | Theme change color | Changed text |
| `HighlightTarget` | Theme target color | Target character |

## Themes

### Built-in Themes

```go
// Default theme (green/red/cyan)
opts.Theme = blockfont.DefaultTheme

// Kartoza theme (orange/blue)
opts.Theme = blockfont.KartozaTheme
```

### Custom Theme

```go
opts.Theme = blockfont.Theme{
    Correct:   "\033[1;32m",  // Green
    Incorrect: "\033[1;31m",  // Red
    Cursor:    "\033[1;36m",  // Cyan
    Pending:   "\033[2m",     // Dim
    Delete:    "\033[1;33m",  // Yellow
    Change:    "\033[1;35m",  // Magenta
    Target:    "\033[1;34m",  // Blue
    Normal:    "",            // Default
}
```

## Styling with Lipgloss

Combine with lipgloss for advanced styling:

```go
import "github.com/charmbracelet/lipgloss"

// Create a styled container
boxStyle := lipgloss.NewStyle().
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#FF6B35")).
    Padding(1, 2)

func (m model) View() string {
    return boxStyle.Render(m.widget.View())
}
```

## Word Wrapping

Enable word wrapping for long text:

```go
opts := blockfont.DefaultWidgetOptions()
opts.Width = 60       // Maximum width
opts.WordWrap = true  // Enable wrapping

widget := blockfont.NewWidget(opts)
widget.SetText("hello world this is a long sentence")
```

The text will wrap at word boundaries to fit within the specified width.

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
