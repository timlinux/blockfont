# Vim-Style Editing

Enable vim-style editing for interactive text manipulation.

## Enabling Vim Mode

```go
opts := blockfont.DefaultWidgetOptions()
opts.VimMode = true
widget := blockfont.NewWidget(opts)
```

## Modes

| Mode | Description |
|------|-------------|
| NORMAL | Navigation and commands |
| INSERT | Text input |
| VISUAL | Selection (future) |

## Normal Mode Keys

| Key | Action |
|-----|--------|
| h, ← | Move left |
| l, → | Move right |
| j, ↓ | Move down |
| k, ↑ | Move up |
| 0, Home | Line start |
| $, End | Line end |
| g | First line |
| G | Last line |
| i | Insert before cursor |
| a | Insert after cursor |
| A | Insert at line end |
| I | Insert at line start |
| o | Open line below |
| O | Open line above |
| x | Delete character |
| dd | Delete line |
| D | Delete to end of line |

## Insert Mode Keys

| Key | Action |
|-----|--------|
| Esc | Return to normal mode |
| Backspace | Delete before cursor |
| Delete | Delete at cursor |
| ←→↑↓ | Move cursor |
| Any char | Insert text |

## Using the Buffer Directly

```go
buffer := blockfont.NewBuffer("hello world")

// Navigate
buffer.MoveRight(5)

// Edit
buffer.SetMode(blockfont.ModeInsert)
buffer.Insert(" there")

// Get result
text := buffer.Text()  // "hello there world"
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
