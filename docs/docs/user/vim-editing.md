# Vim-Style Editing

blockfont includes a full vim-style text editing buffer for interactive text manipulation.

## Enabling Vim Mode

### In Widget

```go
opts := blockfont.DefaultWidgetOptions()
opts.VimMode = true

widget := blockfont.NewWidget(opts)
widget.SetText("edit me")
widget.Focus()  // Required to receive key events
```

### Using Buffer Directly

```go
buffer := blockfont.NewBuffer("hello world")
buffer.SetMode(blockfont.ModeNormal)
```

## Editing Modes

| Mode | Description | Activation |
|------|-------------|------------|
| `ModeNormal` | Navigation and commands | `Esc` from insert mode |
| `ModeInsert` | Text input | `i`, `a`, `o`, etc. |
| `ModeVisual` | Character selection | `v` (future) |
| `ModeVisualLine` | Line selection | `V` (future) |
| `ModeVisualBlock` | Block selection | `Ctrl+v` (future) |
| `ModeCommand` | Command input | `:` (future) |

## Normal Mode Commands

### Navigation

| Key | Action |
|-----|--------|
| `h`, `←` | Move cursor left |
| `l`, `→` | Move cursor right |
| `j`, `↓` | Move cursor down (multi-line) |
| `k`, `↑` | Move cursor up (multi-line) |
| `0`, `Home` | Move to line start |
| `$`, `End` | Move to line end |
| `g` | Move to first line |
| `G` | Move to last line |
| `w` | Move to next word |
| `b` | Move to previous word |

### Entering Insert Mode

| Key | Action |
|-----|--------|
| `i` | Insert before cursor |
| `a` | Insert after cursor |
| `I` | Insert at line start |
| `A` | Insert at line end |
| `o` | Open new line below and insert |
| `O` | Open new line above and insert |

### Editing

| Key | Action |
|-----|--------|
| `x` | Delete character at cursor |
| `X` | Delete character before cursor |
| `dd` | Delete entire line |
| `D` | Delete from cursor to end of line |
| `r` + char | Replace character at cursor |
| `u` | Undo (if implemented) |

## Insert Mode Commands

| Key | Action |
|-----|--------|
| `Esc` | Return to normal mode |
| `Backspace` | Delete character before cursor |
| `Delete` | Delete character at cursor |
| `←` `→` `↑` `↓` | Move cursor |
| Any character | Insert at cursor |

## Mode Indicator

Display the current mode in your UI:

```go
func (m model) View() string {
    mode := m.widget.Mode()

    var modeText string
    switch mode {
    case blockfont.ModeNormal:
        modeText = "-- NORMAL --"
    case blockfont.ModeInsert:
        modeText = "-- INSERT --"
    case blockfont.ModeVisual:
        modeText = "-- VISUAL --"
    }

    return fmt.Sprintf("%s\n\n%s", modeText, m.widget.View())
}
```

## Using the Buffer API

For programmatic editing:

```go
buffer := blockfont.NewBuffer("hello world")

// Check initial state
fmt.Println(buffer.Text())      // "hello world"
fmt.Println(buffer.Mode())      // ModeNormal
x, y := buffer.CursorPosition() // 0, 0

// Navigate
buffer.MoveRight()      // Move one character right
buffer.MoveWordForward() // Move to next word
buffer.MoveToLineEnd()   // Move to end of line

// Enter insert mode and type
buffer.SetMode(blockfont.ModeInsert)
buffer.Insert("!")

// Get result
fmt.Println(buffer.Text()) // "hello world!"

// Delete operations
buffer.SetMode(blockfont.ModeNormal)
buffer.MoveToLineStart()
buffer.Delete(6)           // Delete "hello "
fmt.Println(buffer.Text()) // "world!"
```

## Complete Widget Example

```go
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

    boxStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#FF6B35")).
        Padding(1, 2)

    helpStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("#666666"))
)

type model struct {
    widget *blockfont.Widget
}

func initialModel() model {
    opts := blockfont.DefaultWidgetOptions()
    opts.Width = 60
    opts.VimMode = true
    opts.Theme = blockfont.KartozaTheme

    widget := blockfont.NewWidget(opts)
    widget.SetText("edit me")
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

    w, cmd := m.widget.Update(msg)
    m.widget = w
    return m, cmd
}

func (m model) View() string {
    mode := modeStyle.Render(m.widget.Mode().String())

    editor := boxStyle.Render(m.widget.View())

    var help string
    if m.widget.Mode() == blockfont.ModeNormal {
        help = helpStyle.Render(
            "[i] Insert  [a] Append  [x] Delete  [h/l] Move  [0/$] Start/End")
    } else {
        help = helpStyle.Render("[Esc] Normal mode  Type to insert")
    }

    return fmt.Sprintf("%s\n\n%s\n\n%s\n\nPress Ctrl+C to quit", mode, editor, help)
}

func main() {
    p := tea.NewProgram(initialModel(), tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

## Cursor Appearance

The cursor appears as an underline below the current character in normal mode, and as a vertical line between characters in insert mode.

```go
// Block cursor (normal mode default)
opts.CursorStyle = blockfont.CursorBlock

// Line cursor (insert mode style)
opts.CursorStyle = blockfont.CursorLine

// Underline cursor
opts.CursorStyle = blockfont.CursorUnderline
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
