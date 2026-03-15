// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

// Preview all lowercase letters.
package main

import (
	"fmt"

	"github.com/timlinux/blockfont"
)

func main() {
	fmt.Println("=== Lowercase Letters Preview ===")
	fmt.Println()
	fmt.Println("a b c d e:")
	fmt.Println(blockfont.RenderText("abcde"))
	fmt.Println()
	fmt.Println("f g h i j:")
	fmt.Println(blockfont.RenderText("fghij"))
	fmt.Println()
	fmt.Println("k l m n o:")
	fmt.Println(blockfont.RenderText("klmno"))
	fmt.Println()
	fmt.Println("p q r s t:")
	fmt.Println(blockfont.RenderText("pqrst"))
	fmt.Println()
	fmt.Println("u v w x y z:")
	fmt.Println(blockfont.RenderText("uvwxyz"))
	fmt.Println()
	fmt.Println("=== Full Alphabet ===")
	fmt.Println()
	fmt.Println(blockfont.RenderText("abcdefghijklm"))
	fmt.Println()
	fmt.Println(blockfont.RenderText("nopqrstuvwxyz"))
}
