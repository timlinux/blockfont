# Basic Usage

This guide covers the core rendering functions in blockfont.

## Core Functions

### RenderText

The simplest way to render block text as a string:

```go
package main

import (
    "fmt"
    "github.com/timlinux/blockfont"
)

func main() {
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

### RenderWord

For more control, use `RenderWord` which returns a 2D array of strings:

```go
// Returns [][]string - 6 rows, each containing character columns
lines := blockfont.RenderWord("hi")

// lines[0] = []string{"█  █", "  █"}  // Row 0 of 'h' and 'i'
// lines[1] = []string{"█  █", "   "}  // Row 1
// ... 6 rows total

// Join with spacing to get output
for _, row := range lines {
    fmt.Println(strings.Join(row, " "))
}
```

### RenderWithCursor

Full-featured rendering with cursor position, highlights, and word wrapping:

```go
text := "hello"
cursorPos := 2          // Cursor at 'l'
highlights := []blockfont.CharHighlight{
    blockfont.HighlightCorrect,   // h - green
    blockfont.HighlightCorrect,   // e - green
    blockfont.HighlightPending,   // l - dim
    blockfont.HighlightPending,   // l - dim
    blockfont.HighlightPending,   // o - dim
}
isInsertMode := false
maxWidth := 80
theme := blockfont.DefaultTheme

lines := blockfont.RenderWithCursor(
    text, cursorPos, highlights, isInsertMode, maxWidth, theme)

for _, line := range lines {
    fmt.Println(line)
}
```

### RenderPlainText

Simple colored rendering without cursor:

```go
// Render in red
lines := blockfont.RenderPlainText("error", blockfont.ANSIRed)
fmt.Println(strings.Join(lines, "\n"))

// Render without color
lines = blockfont.RenderPlainText("text", "")
fmt.Println(strings.Join(lines, "\n"))
```

## Getting Dimensions

```go
// Width of a single character in runes
width := blockfont.GetLetterWidth('w')  // Returns 7

// Total width of a word (including letter spacing)
totalWidth := blockfont.GetTotalWidth("hello")  // Returns ~35

// Width including cursor position
runes := []rune("hello")
widthWithCursor := blockfont.CalculateTotalWidth(runes, 2, false)

// Display width of rendered text
displayWidth := blockfont.GetDisplayWidth("hello")
```

## Layout Functions

### Centering Lines

```go
text := blockfont.RenderText("hello")
lines := strings.Split(text, "\n")

// Center within 80 columns
centered := blockfont.CenterLines(lines, 80)
fmt.Println(strings.Join(centered, "\n"))
```

### Left and Right Justify

```go
lines := strings.Split(blockfont.RenderText("hi"), "\n")

// Left justify with 4-space margin
leftAligned := blockfont.LeftJustify(lines, 4)

// Right justify within 80 columns
rightAligned := blockfont.RightJustify(lines, 80)
```

### Using Alignment Enum

```go
lines := strings.Split(blockfont.RenderText("hi"), "\n")

// Use alignment enum
aligned := blockfont.AlignLines(lines, 80, blockfont.AlignCenter)
// Also: blockfont.AlignLeft, blockfont.AlignRight
```

### Word Wrapping

Split text into multiple lines that fit within a width:

```go
// Get words that fit in width
words := blockfont.WrapOnWordBoundaries("hello world demo text", 40)
// Returns: []string{"hello", "world", "demo", "text"} or grouped by line
```

### Utility Functions

```go
// Pad a line to specific width
padded := blockfont.PadToWidth(line, 80)

// Find maximum line width
maxWidth := blockfont.MaxLineWidth(lines)

// Join block letter arrays horizontally
blocks := [][][]string{
    blockfont.RenderWord("hello"),
    blockfont.RenderWord("world"),
}
joined := blockfont.JoinBlockLines(blocks)
```

## Styling and Colors

### ANSI Color Constants

```go
blockfont.ANSIReset    // Reset all formatting
blockfont.ANSIRed      // Bold red
blockfont.ANSIGreen    // Bold green
blockfont.ANSIOrange   // Bold orange/yellow
blockfont.ANSIBlue     // Bold blue
blockfont.ANSICyan     // Bold cyan
blockfont.ANSIMagenta  // Bold magenta
blockfont.ANSIWhite    // Bold white
blockfont.ANSIDim      // Dimmed/faded
blockfont.ANSIInverse  // Inverse video
```

### Wrapping Text with Color

```go
text := blockfont.RenderText("success")
colored := blockfont.WrapWithColor(text, blockfont.ANSIGreen)
fmt.Println(colored)
```

### Inverse Video

```go
line := blockfont.RenderText("cursor")[0]
inverted := blockfont.InvertLine(line)
fmt.Println(inverted)
```

### Gradient Colors

Get colors from a red-to-green gradient:

```go
// Get color for value between 0.0 and 1.0
color := blockfont.GetGradientColor(0.5)  // Middle of gradient

// Get color based on WPM (0-200)
wpmColor := blockfont.GetWPMColor(85.0)  // Color for 85 WPM

// Get color based on progress (0.0-1.0)
progressColor := blockfont.GetProgressColor(0.75)  // 75% complete
```

## ANSI Utilities

### Strip ANSI Codes

```go
// Remove all ANSI escape codes from a string
cleaned := blockfont.RemoveANSI(coloredText)
```

### Visible Width

```go
// Get string width excluding ANSI codes
width := blockfont.VisibleStringWidth(coloredLine)
```

### Insert at Position

```go
// Insert overlay text at position in base string
result := blockfont.InsertAt(baseLine, overlay, 10)
```

## Complete Example

```go
package main

import (
    "fmt"
    "strings"
    "github.com/timlinux/blockfont"
)

func main() {
    // Render some text
    text := blockfont.RenderText("blockfont")

    // Split into lines
    lines := strings.Split(text, "\n")

    // Center it
    centered := blockfont.CenterLines(lines, 80)

    // Add some color
    result := blockfont.WrapWithColor(
        strings.Join(centered, "\n"),
        blockfont.ANSICyan,
    )

    fmt.Println(result)
    fmt.Println()

    // Show character highlighting
    highlights := []blockfont.CharHighlight{
        blockfont.HighlightCorrect,
        blockfont.HighlightCorrect,
        blockfont.HighlightCorrect,
        blockfont.HighlightIncorrect,
        blockfont.HighlightPending,
    }

    typingLines := blockfont.RenderWithCursor(
        "hello", -1, highlights, false, 0, blockfont.DefaultTheme)

    fmt.Println("Typing game example:")
    for _, line := range typingLines {
        fmt.Println(line)
    }
}
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
