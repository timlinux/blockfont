# API Reference

This document provides a complete reference for all public types, functions, and constants in the blockfont library.

## Constants

### Font Constants

```go
const LetterHeight = 6       // Height of all block letters in lines
const LetterSpacing = 1      // Space between letters in characters
```

### Animation Constants

```go
const AnimationInterval = 50 * time.Millisecond  // Default animation tick interval
```

### ANSI Color Constants

```go
const ANSIReset   = "\033[0m"      // Reset all formatting
const ANSIRed     = "\033[1;31m"   // Bold red
const ANSIGreen   = "\033[1;32m"   // Bold green
const ANSIOrange  = "\033[1;33m"   // Bold orange/yellow
const ANSIBlue    = "\033[1;34m"   // Bold blue
const ANSIMagenta = "\033[1;35m"   // Bold magenta
const ANSICyan    = "\033[1;36m"   // Bold cyan
const ANSIWhite   = "\033[1;37m"   // Bold white
const ANSIDim     = "\033[2m"      // Dim/faded
const ANSIInverse = "\033[7m"      // Inverse video
```

---

## Types

### Alignment

Text alignment options for layout functions.

```go
type Alignment int

const (
    AlignLeft Alignment = iota   // Left-align text
    AlignCenter                   // Center text
    AlignRight                    // Right-align text
)
```

### Mode

Vim editing modes for the Buffer.

```go
type Mode int

const (
    ModeNormal Mode = iota    // Normal mode (navigation)
    ModeInsert                // Insert mode (typing)
    ModeVisual                // Visual mode (selection)
    ModeVisualLine            // Visual line mode
    ModeVisualBlock           // Visual block mode
    ModeCommand               // Command mode
)

// String returns the mode name
func (m Mode) String() string
```

### CharHighlight

Character highlight states for typing games, diffs, and visual feedback.

```go
type CharHighlight int

const (
    HighlightNone CharHighlight = iota      // No highlight
    HighlightCorrect                         // Correct input (green)
    HighlightIncorrect                       // Incorrect input (red)
    HighlightCursor                          // Cursor position (inverted)
    HighlightDelete                          // Deleted character
    HighlightChange                          // Changed character
    HighlightTarget                          // Target character
    HighlightPending                         // Pending input (dimmed)
)
```

### CursorStyle

Cursor appearance options.

```go
type CursorStyle int

const (
    CursorBlock CursorStyle = iota    // Block cursor (default)
    CursorLine                         // Vertical line cursor
    CursorUnderline                    // Underline cursor
)
```

### TransitionType

Animation transition types.

```go
type TransitionType int

const (
    TransitionSlideUp TransitionType = iota    // Slide up animation
    TransitionSlideDown                         // Slide down animation
    TransitionFadeIn                            // Fade in animation
    TransitionFadeOut                           // Fade out animation
    TransitionScale                             // Scale animation
)
```

### Theme

Color theme configuration.

```go
type Theme struct {
    Correct     string    // Color for correct characters
    Incorrect   string    // Color for incorrect characters
    Cursor      string    // Color for cursor
    Pending     string    // Color for pending characters
    Delete      string    // Color for deleted characters
    Change      string    // Color for changed characters
    Target      string    // Color for target characters
    Normal      string    // Color for normal text
}
```

**Pre-defined Themes:**

```go
var DefaultTheme Theme   // Standard green/red/cyan theme
var KartozaTheme Theme   // Kartoza-branded orange/blue theme
```

### StyleConfig

Styling configuration for the Widget.

```go
type StyleConfig struct {
    Theme           Theme
    CursorStyle     CursorStyle
    ShowCursor      bool
    ShowUnderline   bool
}
```

---

## Widget

High-level component for bubbletea integration.

### WidgetOptions

Configuration options for creating a Widget.

```go
type WidgetOptions struct {
    Width         int           // Maximum width for rendering
    Alignment     Alignment     // Text alignment
    VimMode       bool          // Enable vim-style editing
    Animate       bool          // Enable animations
    WordWrap      bool          // Enable word wrapping
    CursorStyle   CursorStyle   // Cursor appearance
    Theme         Theme         // Color theme
}
```

### Creating a Widget

```go
// DefaultWidgetOptions returns sensible defaults
func DefaultWidgetOptions() WidgetOptions

// NewWidget creates a new Widget with the given options
func NewWidget(opts WidgetOptions) *Widget
```

### Widget Methods

```go
// SetText sets the widget's text content
func (w *Widget) SetText(text string)

// Text returns the current text content
func (w *Widget) Text() string

// View returns the rendered string (implements tea.Model)
func (w *Widget) View() string

// Update handles input messages (implements tea.Model)
func (w *Widget) Update(msg tea.Msg) (*Widget, tea.Cmd)

// Focus gives keyboard focus to the widget
func (w *Widget) Focus()

// Blur removes keyboard focus from the widget
func (w *Widget) Blur()

// Focused returns whether the widget has focus
func (w *Widget) Focused() bool

// Mode returns the current vim mode
func (w *Widget) Mode() Mode

// CursorPosition returns the cursor X position
func (w *Widget) CursorPosition() int
```

### AnimationTickMsg

Message type for animation updates.

```go
type AnimationTickMsg time.Time
```

---

## Buffer

Vim-style text editing buffer.

### Creating a Buffer

```go
// NewBuffer creates a new buffer with initial text
func NewBuffer(text string) *Buffer
```

### Buffer Methods

**Text Operations:**

```go
// Text returns the current text content
func (b *Buffer) Text() string

// SetText replaces all text content
func (b *Buffer) SetText(text string)

// Insert inserts text at the cursor position
func (b *Buffer) Insert(text string)

// Delete removes n characters from cursor position
// Returns the deleted text
func (b *Buffer) Delete(n int) string

// DeleteLine deletes the current line
// Returns the deleted line
func (b *Buffer) DeleteLine() string

// DeleteToEndOfLine deletes from cursor to end of line
// Returns the deleted text
func (b *Buffer) DeleteToEndOfLine() string

// ReplaceChar replaces the character at cursor position
func (b *Buffer) ReplaceChar(r rune)
```

**Cursor Movement:**

```go
// CursorPosition returns (x, y) cursor position
func (b *Buffer) CursorPosition() (x, y int)

// CursorIndex returns the linear cursor index
func (b *Buffer) CursorIndex() int

// SetCursorPosition sets cursor to (x, y) position
func (b *Buffer) SetCursorPosition(x, y int)

// SetCursorIndex sets cursor to linear index
func (b *Buffer) SetCursorIndex(index int)

// MoveLeft moves cursor left
func (b *Buffer) MoveLeft()

// MoveRight moves cursor right
func (b *Buffer) MoveRight()

// MoveUp moves cursor up one line
func (b *Buffer) MoveUp()

// MoveDown moves cursor down one line
func (b *Buffer) MoveDown()

// MoveToLineStart moves cursor to start of line
func (b *Buffer) MoveToLineStart()

// MoveToLineEnd moves cursor to end of line
func (b *Buffer) MoveToLineEnd()

// MoveToStart moves cursor to start of buffer
func (b *Buffer) MoveToStart()

// MoveToEnd moves cursor to end of buffer
func (b *Buffer) MoveToEnd()

// MoveWordForward moves cursor to next word
func (b *Buffer) MoveWordForward()

// MoveWordBackward moves cursor to previous word
func (b *Buffer) MoveWordBackward()
```

**Mode Management:**

```go
// Mode returns the current editing mode
func (b *Buffer) Mode() Mode

// SetMode sets the editing mode
func (b *Buffer) SetMode(mode Mode)
```

**Utility:**

```go
// Clone creates a deep copy of the buffer
func (b *Buffer) Clone() *Buffer

// Length returns the total character count
func (b *Buffer) Length() int
```

---

## Animator

Spring-based animation system using harmonica.

### Creating an Animator

```go
// NewAnimator creates a new animator
func NewAnimator() *Animator
```

### Animator Methods

```go
// TriggerTransition starts a transition animation
func (a *Animator) TriggerTransition(t TransitionType)

// Update advances the animation by one frame
// Returns true if animation is still in progress
func (a *Animator) Update() bool

// IsAnimating returns whether animation is active
func (a *Animator) IsAnimating() bool

// GetOffset returns the current position offset
func (a *Animator) GetOffset() int

// GetOpacity returns the current opacity (0.0-1.0)
func (a *Animator) GetOpacity() float64

// GetOpacityLevel returns opacity in discrete levels
func (a *Animator) GetOpacityLevel(scale float64) float64

// GetScale returns the current scale factor
func (a *Animator) GetScale() float64

// Reset resets the animator to initial state
func (a *Animator) Reset()
```

### WordCarouselAnimator

Specialized animator for word carousel effects.

```go
// NewWordCarouselAnimator creates a carousel animator
func NewWordCarouselAnimator() *WordCarouselAnimator

// TriggerTransition starts the carousel transition
func (w *WordCarouselAnimator) TriggerTransition()

// Update advances the animation
// Returns true if still animating
func (w *WordCarouselAnimator) Update() bool

// GetPrevOpacity returns opacity of previous word
func (w *WordCarouselAnimator) GetPrevOpacity() float64

// GetCurrentOpacity returns opacity of current word
func (w *WordCarouselAnimator) GetCurrentOpacity() float64

// GetNextOpacity returns opacity of next word
func (w *WordCarouselAnimator) GetNextOpacity() float64
```

### Animation Utility

```go
// GetAnimationInterval returns recommended tick interval
func GetAnimationInterval() time.Duration
```

---

## Font Functions

Core rendering functions.

### RenderWord

Renders a word as a 2D array of strings.

```go
func RenderWord(word string) [][]string
```

**Returns:** Array of 6 rows, each containing character columns.

**Example:**

```go
lines := blockfont.RenderWord("hi")
// lines[0] = ["█  █", "  █"]  // Row 0 of 'h' and 'i'
// lines[1] = ["█  █", "   "]  // Row 1
// ... etc
```

### RenderText

Renders text as a single string with newlines.

```go
func RenderText(text string) string
```

**Example:**

```go
output := blockfont.RenderText("hello")
fmt.Println(output)
```

### RenderWithCursor

Full-featured rendering with cursor, highlights, and word wrapping.

```go
func RenderWithCursor(
    text string,
    cursorIdx int,
    highlights []CharHighlight,
    isInsertMode bool,
    maxWidth int,
    theme Theme,
) []string
```

**Parameters:**

| Parameter | Description |
|-----------|-------------|
| `text` | Text to render |
| `cursorIdx` | Cursor position (-1 for no cursor) |
| `highlights` | Per-character highlights (can be nil) |
| `isInsertMode` | True for insert mode cursor |
| `maxWidth` | Maximum width (0 for no wrapping) |
| `theme` | Color theme to use |

**Returns:** Array of rendered lines including underline row.

### RenderPlainText

Simple colored rendering without cursor or highlights.

```go
func RenderPlainText(text string, color string) []string
```

**Example:**

```go
lines := blockfont.RenderPlainText("error", blockfont.ANSIRed)
```

### Width Functions

```go
// GetLetterWidth returns the width of a character in runes
func GetLetterWidth(char rune) int

// GetTotalWidth returns the total width of a word
func GetTotalWidth(word string) int

// CalculateTotalWidth calculates width including cursor
func CalculateTotalWidth(runes []rune, cursorIdx int, isInsertMode bool) int

// GetDisplayWidth returns visible width of rendered text
func GetDisplayWidth(text string) int
```

---

## Layout Functions

Text alignment and positioning utilities.

### Alignment Functions

```go
// CenterLines centers lines within a width
func CenterLines(lines []string, width int) []string

// LeftJustify adds left margin to lines
func LeftJustify(lines []string, margin int) []string

// RightJustify right-aligns lines within a width
func RightJustify(lines []string, width int) []string

// AlignLines applies alignment to lines
func AlignLines(lines []string, width int, align Alignment) []string
```

### Word Wrapping

```go
// WrapOnWordBoundaries wraps text at word boundaries
func WrapOnWordBoundaries(text string, maxWidth int) []string
```

### Utility Functions

```go
// PadToWidth pads a string to specified width
func PadToWidth(line string, width int) string

// MaxLineWidth finds the maximum line width
func MaxLineWidth(lines []string) int

// JoinBlockLines joins block letter arrays horizontally
func JoinBlockLines(blocks [][][]string) []string
```

---

## ANSI Functions

ANSI escape code utilities.

```go
// RemoveANSI strips all ANSI escape codes from a string
func RemoveANSI(s string) string

// VisibleStringWidth returns string width excluding ANSI codes
func VisibleStringWidth(s string) int

// WrapWithColor wraps text with color codes
func WrapWithColor(text, color string) string

// InvertLine applies inverse video to a line
func InvertLine(line string) string

// InsertAt inserts overlay text at position in base string
func InsertAt(base, overlay string, x int) string
```

---

## Style Functions

### Gradient Colors

```go
// GradientColors is a red-to-green gradient array
var GradientColors = []string{
    "\033[38;5;196m",  // Red
    "\033[38;5;202m",  // Orange-red
    "\033[38;5;208m",  // Orange
    "\033[38;5;214m",  // Orange-yellow
    "\033[38;5;220m",  // Yellow
    "\033[38;5;226m",  // Yellow-green
    "\033[38;5;190m",  // Light green
    "\033[38;5;154m",  // Green
    "\033[38;5;118m",  // Bright green
    "\033[38;5;46m",   // Pure green
}

// GetGradientColor returns color for value 0.0-1.0
func GetGradientColor(value float64) string

// GetWPMColor returns color based on WPM (0-200)
func GetWPMColor(wpm float64) string

// GetProgressColor returns color based on progress (0.0-1.0)
func GetProgressColor(progress float64) string
```

---

## Font Data

### BlockLetters

Map of rune to 6-line string arrays containing all supported characters.

```go
var BlockLetters map[rune][]string
```

**Supported Characters:**

- Lowercase: `a-z`
- Uppercase: `A-Z`
- Numbers: `0-9`
- Punctuation: `!@#$%^&*()-_=+[]{}|;:'",.<>/?~` and space

**Example:**

```go
// Access letter data directly
letterH := blockfont.BlockLetters['h']
// letterH[0] = "█  █"  // Row 0
// letterH[1] = "█  █"  // Row 1
// letterH[2] = "████"  // Row 2
// ...
```

---

## Complete Examples

### Simple Rendering

```go
package main

import (
    "fmt"
    "github.com/timlinux/blockfont"
)

func main() {
    // Render text
    output := blockfont.RenderText("hello")
    fmt.Println(output)

    // Center it
    lines := strings.Split(output, "\n")
    centered := blockfont.CenterLines(lines, 80)
    fmt.Println(strings.Join(centered, "\n"))
}
```

### Typing Game Highlights

```go
package main

import (
    "fmt"
    "strings"
    "github.com/timlinux/blockfont"
)

func main() {
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
    fmt.Println(strings.Join(lines, "\n"))
}
```

### Bubbletea Widget with Vim Editing

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
    return m.widget.View() + "\n\nPress Ctrl+C to quit"
}

func main() {
    opts := blockfont.DefaultWidgetOptions()
    opts.Width = 60
    opts.VimMode = true
    opts.Theme = blockfont.KartozaTheme

    widget := blockfont.NewWidget(opts)
    widget.SetText("edit me")
    widget.Focus()

    p := tea.NewProgram(model{widget: widget}, tea.WithAltScreen())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

### Animated Text

```go
package main

import (
    "fmt"
    "time"
    "github.com/timlinux/blockfont"
)

func main() {
    animator := blockfont.NewAnimator()
    animator.TriggerTransition(blockfont.TransitionFadeIn)

    for animator.Update() {
        opacity := animator.GetOpacityLevel(1.0)
        fmt.Printf("\033[H\033[2J") // Clear screen

        text := blockfont.RenderText("fade")
        if opacity < 0.3 {
            text = "" // Hidden
        } else if opacity < 0.7 {
            text = blockfont.WrapWithColor(text, blockfont.ANSIDim)
        }
        fmt.Println(text)

        time.Sleep(blockfont.AnimationInterval)
    }
}
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
