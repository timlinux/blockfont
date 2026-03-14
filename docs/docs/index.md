# blockfont

Unicode block letter rendering for terminal applications.

![blockfont example](assets/example.png)

## Features

- **Block Letter Rendering**: Convert text to beautiful Unicode block letters
- **Vim-Style Editing**: Full vim buffer with normal/insert modes and common operations
- **Animations**: Smooth spring-based transitions using harmonica
- **Lipgloss Integration**: Full charmbracelet/lipgloss styling support
- **Bubbletea Widget**: Ready-to-use tea.Model for easy integration
- **Character Highlighting**: Color individual characters for typing games, diffs, etc.
- **Layout Utilities**: Centering, alignment, and word wrapping

## Quick Start

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

## Installation

```bash
go get github.com/timlinux/blockfont
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

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, m.widget.Update(msg)
}

func (m model) View() string {
    return m.widget.View()
}

func main() {
    widget := blockfont.NewWidget(blockfont.WidgetOptions{
        Width:     80,
        Alignment: blockfont.AlignCenter,
        VimMode:   true,
    })
    widget.SetText("hello")

    p := tea.NewProgram(model{widget: widget})
    p.Run()
}
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
