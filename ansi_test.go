// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

package blockfont

import (
	"testing"
)

func TestRemoveANSI(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "no ANSI",
			input: "hello",
			want:  "hello",
		},
		{
			name:  "with color",
			input: "\033[1;31mhello\033[0m",
			want:  "hello",
		},
		{
			name:  "multiple codes",
			input: "\033[1;32mgreen\033[0m and \033[1;34mblue\033[0m",
			want:  "green and blue",
		},
		{
			name:  "inverse",
			input: "\033[7mreversed\033[0m",
			want:  "reversed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveANSI(tt.input)
			if got != tt.want {
				t.Errorf("RemoveANSI() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestVisibleStringWidth(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "plain text",
			input: "hello",
			want:  5,
		},
		{
			name:  "with ANSI color",
			input: "\033[1;31mhello\033[0m",
			want:  5,
		},
		{
			name:  "unicode",
			input: "hello 世界",
			want:  8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := VisibleStringWidth(tt.input)
			if got != tt.want {
				t.Errorf("VisibleStringWidth() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestInvertLine(t *testing.T) {
	input := "hello"
	result := InvertLine(input)

	// Should start with inverse code
	if result[:4] != ANSIInverse {
		t.Errorf("InvertLine() should start with inverse code")
	}

	// Should end with reset code
	if result[len(result)-4:] != ANSIReset {
		t.Errorf("InvertLine() should end with reset code")
	}

	// Should contain original text
	if RemoveANSI(result) != input {
		t.Errorf("InvertLine() content = %q, want %q", RemoveANSI(result), input)
	}
}

func TestWrapWithColor(t *testing.T) {
	text := "hello"
	colored := WrapWithColor(text, ANSIRed)

	// Should contain the color code
	if colored[:7] != ANSIRed {
		t.Errorf("WrapWithColor() should start with color code")
	}

	// Should contain reset
	if colored[len(colored)-4:] != ANSIReset {
		t.Errorf("WrapWithColor() should end with reset")
	}

	// Empty color should return unchanged text
	unchanged := WrapWithColor(text, "")
	if unchanged != text {
		t.Errorf("WrapWithColor with empty color = %q, want %q", unchanged, text)
	}
}

func TestInsertAt(t *testing.T) {
	// Note: InsertAt replaces characters at position, not inserts before
	tests := []struct {
		name    string
		base    string
		overlay string
		x       int
		want    string
	}{
		{
			name:    "overlay at start",
			base:    "xxxxx",
			overlay: "hello",
			x:       0,
			want:    "hello",
		},
		{
			name:    "overlay in middle",
			base:    "xxxx world",
			overlay: "hello",
			x:       0,
			want:    "helloworld",
		},
		{
			name:    "overlay past end",
			base:    "hi",
			overlay: "there",
			x:       5,
			want:    "hi   there",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InsertAt(tt.base, tt.overlay, tt.x)
			if got != tt.want {
				t.Errorf("InsertAt() = %q, want %q", got, tt.want)
			}
		})
	}
}
