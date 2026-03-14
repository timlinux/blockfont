// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

// Simple example demonstrating basic blockfont usage.
package main

import (
	"fmt"

	"github.com/timlinux/blockfont"
)

func main() {
	// Basic text rendering
	fmt.Println("=== Simple Text Rendering ===")
	fmt.Println()
	text := blockfont.RenderText("hello")
	fmt.Println(text)
	fmt.Println()

	// With mixed case
	fmt.Println("=== Mixed Case ===")
	fmt.Println()
	text = blockfont.RenderText("Hello World")
	fmt.Println(text)
	fmt.Println()

	// Get dimensions
	fmt.Println("=== Dimensions ===")
	fmt.Printf("Letter 'a' width: %d\n", blockfont.GetLetterWidth('a'))
	fmt.Printf("Letter 'w' width: %d\n", blockfont.GetLetterWidth('w'))
	fmt.Printf("Word 'hello' total width: %d\n", blockfont.GetTotalWidth("hello"))
	fmt.Println()

	// Layout - centering
	fmt.Println("=== Centered Text (width 80) ===")
	lines := blockfont.RenderWord("hi")
	joined := blockfont.JoinBlockLines(lines, blockfont.LetterSpacing)
	centered := blockfont.CenterLines(joined, 80)
	for _, line := range centered {
		fmt.Println(line)
	}
	fmt.Println()

	// With color
	fmt.Println("=== Colored Text ===")
	coloredText := blockfont.WrapWithColor(blockfont.RenderText("red"), blockfont.ANSIRed)
	fmt.Println(coloredText)

	greenText := blockfont.WrapWithColor(blockfont.RenderText("green"), blockfont.ANSIGreen)
	fmt.Println(greenText)
	fmt.Println()

	fmt.Println("Made with ❤️ by Kartoza | https://kartoza.com")
}
