# Architecture

## Overview

blockfont is organized as a single Go package with clear separation of concerns.

```
blockfont/
├── font.go        # Core font data and rendering
├── widget.go      # High-level bubbletea widget
├── buffer.go      # Vim-style text buffer
├── layout.go      # Alignment and layout utilities
├── style.go       # Styling and themes
├── animation.go   # Spring-based animations
└── ansi.go        # ANSI escape code utilities
```

## Component Relationships

```
Widget (tea.Model)
    │
    ├── Buffer (vim editing)
    │       └── lines, cursor, mode
    │
    ├── Animator (transitions)
    │       └── springs, position, opacity, scale
    │
    └── Theme (styling)
            └── colors, highlights

Core Functions
    │
    ├── RenderWord() ──► BlockLetters map
    │
    └── Layout Utils ──► ANSI Utils
```

## Data Flow

1. **Input**: User sets text via `SetText()`
2. **Processing**: Text stored in buffer (if vim mode) or directly
3. **Rendering**: `View()` calls `RenderWord()` for each character
4. **Styling**: Highlights and colors applied per character
5. **Layout**: Alignment and centering applied
6. **Output**: Final string returned

## Key Design Decisions

- **Single Package**: Keeps API simple and reduces import complexity
- **Lipgloss Integration**: Leverages existing styling ecosystem
- **Spring Physics**: Natural, polished feel for animations
- **Vim Buffer**: Familiar editing paradigm for power users

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
