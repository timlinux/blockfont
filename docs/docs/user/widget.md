# Widget Integration

The Widget provides a high-level interface for bubbletea applications.

## Creating a Widget

```go
opts := blockfont.DefaultWidgetOptions()
opts.Width = 80
opts.Alignment = blockfont.AlignCenter

widget := blockfont.NewWidget(opts)
widget.SetText("hello")
```

## Widget Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| Width | int | 80 | Target render width |
| Height | int | 0 | Target render height |
| Alignment | Alignment | AlignLeft | Text alignment |
| VimMode | bool | false | Enable vim editing |
| Animate | bool | false | Enable animations |
| WordWrap | bool | false | Enable word wrapping |
| CursorStyle | CursorStyle | CursorBlock | Cursor appearance |
| Theme | Theme | DefaultTheme | Color theme |

## Bubbletea Integration

```go
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
```

## Character Highlighting

For typing games or diffs:

```go
highlights := []blockfont.CharHighlight{
    blockfont.HighlightCorrect,   // Green
    blockfont.HighlightIncorrect, // Red
    blockfont.HighlightPending,   // Gray
}
widget.SetHighlights(highlights)
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
