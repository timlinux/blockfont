# Animations

blockfont uses spring-based physics (via charmbracelet/harmonica) for smooth, natural animations.

## Enabling Animations in Widget

```go
opts := blockfont.DefaultWidgetOptions()
opts.Animate = true

widget := blockfont.NewWidget(opts)
widget.SetText("animate me")
```

## Animation Interval

All animations use a consistent tick interval:

```go
// Default interval (50ms = 20fps)
interval := blockfont.AnimationInterval

// Or get it via function
interval := blockfont.GetAnimationInterval()
```

## Transition Types

| Type | Description |
|------|-------------|
| `TransitionSlideUp` | Slides content upward from below |
| `TransitionSlideDown` | Slides content downward from above |
| `TransitionFadeIn` | Fades content in from transparent |
| `TransitionFadeOut` | Fades content out to transparent |
| `TransitionScale` | Scales content from small to full size |

## Using the Animator

### Creating an Animator

```go
animator := blockfont.NewAnimator()
```

### Triggering Animations

```go
// Trigger a fade-in animation
animator.TriggerTransition(blockfont.TransitionFadeIn)

// Trigger a slide animation
animator.TriggerTransition(blockfont.TransitionSlideUp)
```

### Animation Loop

The `Update()` method advances the animation by one frame and returns `true` while the animation is still running:

```go
for animator.Update() {
    // Animation is still running
    // Get current values
    offset := animator.GetOffset()
    opacity := animator.GetOpacity()
    scale := animator.GetScale()

    // Render your content with these values
    // ...

    time.Sleep(blockfont.AnimationInterval)
}
// Animation complete
```

### Animator Methods

```go
// Check if animation is active
if animator.IsAnimating() {
    // ...
}

// Get position offset (for slide animations)
offset := animator.GetOffset()  // Returns int

// Get opacity (for fade animations)
opacity := animator.GetOpacity()  // Returns 0.0 to 1.0

// Get opacity in discrete levels (useful for text effects)
level := animator.GetOpacityLevel(1.0)  // Scale factor

// Get scale (for scale animations)
scale := animator.GetScale()  // Returns 0.0 to 1.0

// Reset animator to initial state
animator.Reset()
```

## Bubbletea Integration

### Animation Tick Message

```go
import (
    "time"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/timlinux/blockfont"
)

type tickMsg time.Time

func tickCmd() tea.Cmd {
    return tea.Tick(blockfont.AnimationInterval, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}
```

### Complete Animation Example

```go
package main

import (
    "fmt"
    "os"
    "time"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/timlinux/blockfont"
)

type tickMsg time.Time

type model struct {
    animator     *blockfont.Animator
    text         string
    animating    bool
}

func initialModel() model {
    return model{
        animator:  blockfont.NewAnimator(),
        text:      "fade",
        animating: false,
    }
}

func tickCmd() tea.Cmd {
    return tea.Tick(blockfont.AnimationInterval, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            return m, tea.Quit
        case " ", "enter":
            // Trigger animation
            m.animator.TriggerTransition(blockfont.TransitionFadeIn)
            m.animating = true
            return m, tickCmd()
        }

    case tickMsg:
        if m.animator.Update() {
            // Still animating
            return m, tickCmd()
        }
        m.animating = false
    }

    return m, nil
}

func (m model) View() string {
    var content string

    opacity := m.animator.GetOpacityLevel(1.0)
    rendered := blockfont.RenderText(m.text)

    if opacity > 0.7 {
        // Full opacity
        content = rendered
    } else if opacity > 0.3 {
        // Dimmed
        content = blockfont.WrapWithColor(rendered, blockfont.ANSIDim)
    } else {
        // Hidden
        content = ""
    }

    help := "\nPress SPACE to animate, Q to quit"
    return content + help
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v\n", err)
        os.Exit(1)
    }
}
```

## Word Carousel Animator

The `WordCarouselAnimator` is specialized for speed reading applications where words cycle through view:

```go
carousel := blockfont.NewWordCarouselAnimator()
words := []string{"block", "font", "demo", "cool"}
currentWord := 0

// Trigger word change
carousel.TriggerTransition()
currentWord = (currentWord + 1) % len(words)
```

### Carousel Methods

```go
// Get opacity for previous word (fading out)
prevOpacity := carousel.GetPrevOpacity()

// Get opacity for current word (fading in)
currentOpacity := carousel.GetCurrentOpacity()

// Get opacity for next word (preview)
nextOpacity := carousel.GetNextOpacity()

// Update animation frame
if carousel.Update() {
    // Still animating
}
```

### Carousel Example

```go
func (m model) View() string {
    var s strings.Builder

    prevIdx := (m.currentWord - 1 + len(m.words)) % len(m.words)
    nextIdx := (m.currentWord + 1) % len(m.words)

    // Previous word (fading out)
    if m.carousel.GetPrevOpacity() > 0.2 {
        prev := blockfont.WrapWithColor(
            blockfont.RenderText(m.words[prevIdx]),
            blockfont.ANSIDim)
        s.WriteString(prev)
        s.WriteString("\n")
    }

    // Current word (fully visible)
    current := blockfont.RenderText(m.words[m.currentWord])
    s.WriteString(current)
    s.WriteString("\n")

    // Next word (preview)
    if m.carousel.GetNextOpacity() > 0.2 {
        next := blockfont.WrapWithColor(
            blockfont.RenderText(m.words[nextIdx]),
            blockfont.ANSIDim)
        s.WriteString(next)
    }

    return s.String()
}
```

## Combining Animations with Styling

Use animations with lipgloss for advanced effects:

```go
import "github.com/charmbracelet/lipgloss"

func (m model) View() string {
    opacity := m.animator.GetOpacity()

    // Fade effect using lipgloss
    style := lipgloss.NewStyle()

    if opacity > 0.7 {
        style = style.Foreground(lipgloss.Color("#FFFFFF"))
    } else if opacity > 0.3 {
        style = style.Foreground(lipgloss.Color("#888888"))
    } else {
        style = style.Foreground(lipgloss.Color("#333333"))
    }

    return style.Render(blockfont.RenderText(m.text))
}
```

## Spring Physics

The animations use spring physics from charmbracelet/harmonica, which provides:

- **Natural motion**: Objects accelerate and decelerate naturally
- **Overshoot**: Optional bounce effect at the end
- **Smooth interpolation**: No jarring steps

The spring parameters are tuned for a responsive feel appropriate for terminal UIs.

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
