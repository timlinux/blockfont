# API Reference

## Constants

```go
const LetterHeight = 6
const LetterSpacing = 1
const AnimationInterval = 50 * time.Millisecond
```

## Types

### Alignment

```go
type Alignment int

const (
    AlignLeft Alignment = iota
    AlignCenter
    AlignRight
)
```

### Mode

```go
type Mode int

const (
    ModeNormal Mode = iota
    ModeInsert
    ModeVisual
    ModeVisualLine
    ModeVisualBlock
    ModeCommand
)
```

### CharHighlight

```go
type CharHighlight int

const (
    HighlightNone CharHighlight = iota
    HighlightCorrect
    HighlightIncorrect
    HighlightCursor
    HighlightDelete
    HighlightChange
    HighlightTarget
    HighlightPending
)
```

### TransitionType

```go
type TransitionType int

const (
    TransitionSlideUp TransitionType = iota
    TransitionSlideDown
    TransitionFadeIn
    TransitionFadeOut
    TransitionScale
)
```

## Functions

### Font Functions

| Function | Description |
|----------|-------------|
| `RenderWord(word string) [][]string` | Render word as 2D array |
| `RenderText(text string) string` | Render text as string |
| `GetLetterWidth(char rune) int` | Get character width |
| `GetTotalWidth(word string) int` | Get total word width |

### Layout Functions

| Function | Description |
|----------|-------------|
| `CenterLines(lines []string, width int) []string` | Center lines |
| `LeftJustify(lines []string, margin int) []string` | Left-justify |
| `RightJustify(lines []string, width int) []string` | Right-justify |
| `WrapOnWordBoundaries(text string, maxWidth int) []string` | Word wrap |
| `PadToWidth(line string, width int) string` | Pad to width |
| `MaxLineWidth(lines []string) int` | Find max width |

### ANSI Functions

| Function | Description |
|----------|-------------|
| `RemoveANSI(s string) string` | Strip ANSI codes |
| `VisibleStringWidth(s string) int` | Width without ANSI |
| `WrapWithColor(text, color string) string` | Add color |
| `InvertLine(line string) string` | Apply inverse |

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
