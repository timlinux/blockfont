# Basic Usage

## Core Functions

### RenderText

The simplest way to render block text:

```go
text := blockfont.RenderText("hello")
fmt.Println(text)
```

### RenderWord

For more control, use `RenderWord` which returns a 2D array:

```go
// Returns [][]string - 6 rows, each containing character renderings
lines := blockfont.RenderWord("hi")

// lines[0] = ["██  ██", "██████"] (first row of 'h' and 'i')
// lines[1] = ["██  ██", "  ██  "] (second row)
// ... etc
```

### Getting Dimensions

```go
// Width of a single character
width := blockfont.GetLetterWidth('w')  // 7

// Total width of a word (including spacing)
totalWidth := blockfont.GetTotalWidth("hello")  // ~35
```

## Layout

### Centering

```go
lines := blockfont.RenderWord("hi")
joined := blockfont.JoinBlockLines(lines, blockfont.LetterSpacing)
centered := blockfont.CenterLines(joined, 80)
```

### Alignment

```go
// Center, Left, or Right
aligned := blockfont.AlignLines(lines, blockfont.AlignCenter, 80)
```

### Word Wrapping

```go
wrapped := blockfont.WrapOnWordBoundaries("hello world", 40)
// Returns []string{"hello", "world"} if width is too small
```

## Styling

### ANSI Colors

```go
colored := blockfont.WrapWithColor(text, blockfont.ANSIRed)
fmt.Println(colored)  // Red text
```

### Inversion

```go
inverted := blockfont.InvertLine(line)
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
