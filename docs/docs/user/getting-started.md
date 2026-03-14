# Getting Started

This guide will help you get up and running with blockfont quickly.

## Installation

Install blockfont using `go get`:

```bash
go get github.com/timlinux/blockfont
```

## Requirements

- Go 1.21 or later
- Terminal with Unicode support

### Optional Dependencies

For bubbletea integration:
```bash
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/lipgloss
```

For animations:
```bash
go get github.com/charmbracelet/harmonica
```

## Quick Example

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

## Supported Characters

blockfont supports a comprehensive character set:

| Category | Characters |
|----------|------------|
| Lowercase | `a b c d e f g h i j k l m n o p q r s t u v w x y z` |
| Uppercase | `A B C D E F G H I J K L M N O P Q R S T U V W X Y Z` |
| Numbers | `0 1 2 3 4 5 6 7 8 9` |
| Punctuation | `! @ # $ % ^ & * ( ) - _ = + [ ] { } \| ; : ' " , . < > / ? ~` |
| Special | `space` |

## Unicode Block Elements

blockfont uses these Unicode characters:

| Character | Name | Code Point |
|-----------|------|------------|
| `█` | Full Block | U+2588 |
| `◢` | Lower Right Triangle | U+25E2 |
| `◣` | Lower Left Triangle | U+25E3 |
| `◤` | Upper Left Triangle | U+25E4 |
| `◥` | Upper Right Triangle | U+25E5 |

## Centering Text

```go
package main

import (
    "fmt"
    "strings"
    "github.com/timlinux/blockfont"
)

func main() {
    // Render text
    text := blockfont.RenderText("hello")

    // Split into lines
    lines := strings.Split(text, "\n")

    // Center within 80 columns
    centered := blockfont.CenterLines(lines, 80)

    fmt.Println(strings.Join(centered, "\n"))
}
```

## Adding Color

```go
package main

import (
    "fmt"
    "strings"
    "github.com/timlinux/blockfont"
)

func main() {
    // Render with color
    lines := blockfont.RenderPlainText("error", blockfont.ANSIRed)
    fmt.Println(strings.Join(lines, "\n"))

    // Or wrap existing text
    text := blockfont.RenderText("warning")
    colored := blockfont.WrapWithColor(text, blockfont.ANSIOrange)
    fmt.Println(colored)
}
```

## Running the Demo

To see all features in action:

```bash
# Clone the repository
git clone https://github.com/timlinux/blockfont
cd blockfont

# Enter nix development shell (recommended)
nix develop

# Run the demo
make demo
```

The demo includes screens for:

1. **Welcome** - Overview and features
2. **Basic Rendering** - Character sets
3. **Highlights** - Character coloring
4. **Vim Editing** - Interactive editing
5. **Animations** - Transitions
6. **Word Wrap** - Layout features
7. **Themes** - Color themes
8. **Credits**

## Next Steps

- [Basic Usage](basic-usage.md) - Learn the core rendering functions
- [Widget Guide](widget.md) - Integrate with bubbletea
- [Vim Editing](vim-editing.md) - Enable vim-style editing
- [Animations](animations.md) - Add smooth transitions
- [API Reference](../developer/api.md) - Complete API documentation

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
