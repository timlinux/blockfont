# blockfont

Block letter rendering for terminal applications using ASCII block characters.

![blockfont demo](images/blockfont-demo.gif)

## Features

- **Block Letter Rendering**: Convert text to beautiful block letters using ASCII block characters like `█ ◢ ◣ ◤ ◥`
- **Full Character Set**: Lowercase (a-z), uppercase (A-Z), numbers (0-9), and common punctuation
- **Vim-Style Editing**: Full vim buffer with normal/insert/visual modes and common operations
- **Character Highlighting**: Color individual characters for typing games, diffs, and visual feedback
- **Spring Animations**: Smooth transitions using charmbracelet/harmonica physics
- **Word Wrapping**: Intelligent word-boundary wrapping with cursor position tracking
- **Lipgloss Integration**: Full charmbracelet/lipgloss styling support
- **Bubbletea Widget**: Ready-to-use tea.Model for easy integration
- **Layout Utilities**: Centering, alignment, and positioning helpers

## Quick Start

```bash
go get github.com/timlinux/blockfont
```

```go
package main

import (
    "fmt"
    "github.com/timlinux/blockfont"
)

func main() {
    // Simple rendering
    text := blockfont.RenderText("hello")
    fmt.Println(text)
}
```

**Output:**

```
█  █ ████ █    █    ████
█  █ █    █    █    █  █
████ ███  █    █    █  █
█  █ █    █    █    █  █
█  █ █    █    █    █  █
█  █ ████ ████ ████ ████
```

## Widget Example

```go
package main

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/timlinux/blockfont"
)

type model struct {
    widget *blockfont.Widget
}

func (m model) Init() tea.Cmd { return nil }

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
    opts.VimMode = true
    opts.Theme = blockfont.KartozaTheme

    widget := blockfont.NewWidget(opts)
    widget.SetText("edit me")
    widget.Focus()

    tea.NewProgram(model{widget: widget}).Run()
}
```

## Character Highlighting

Perfect for typing games, code diffs, and visual feedback:

```go
text := "typing"
highlights := []blockfont.CharHighlight{
    blockfont.HighlightCorrect,    // t - green
    blockfont.HighlightCorrect,    // y - green
    blockfont.HighlightIncorrect,  // p - red
    blockfont.HighlightPending,    // i - dim
    blockfont.HighlightPending,    // n - dim
    blockfont.HighlightPending,    // g - dim
}

lines := blockfont.RenderWithCursor(
    text, -1, highlights, false, 0, blockfont.DefaultTheme)
```

## Themes

Built-in themes with customizable colors:

```go
// Use the default theme
theme := blockfont.DefaultTheme

// Or use the Kartoza theme
theme := blockfont.KartozaTheme

// Create custom theme
theme := blockfont.Theme{
    Correct:   blockfont.ANSIGreen,
    Incorrect: blockfont.ANSIRed,
    Cursor:    blockfont.ANSICyan,
    Pending:   blockfont.ANSIDim,
}
```

## Run the Demo

Try out all features interactively:

```bash
# Clone the repository
git clone https://github.com/timlinux/blockfont
cd blockfont

# Run the demo
make demo

# Or with nix
nix run
```

## Documentation

- [Getting Started](user/getting-started.md) - Installation and first steps
- [Basic Usage](user/basic-usage.md) - Core rendering functions
- [Widget Guide](user/widget.md) - Bubbletea integration
- [Vim Editing](user/vim-editing.md) - Buffer operations
- [Animations](user/animations.md) - Spring-based transitions
- [API Reference](developer/api.md) - Complete API documentation

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
