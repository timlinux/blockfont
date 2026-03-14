// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

package blockfont

import (
	"strings"
	"testing"
)

func TestRenderWithCursor(t *testing.T) {
	tests := []struct {
		name         string
		text         string
		cursorIdx    int
		isInsertMode bool
		maxWidth     int
		wantRows     int
	}{
		{
			name:         "simple text no cursor",
			text:         "hi",
			cursorIdx:    -1,
			isInsertMode: false,
			maxWidth:     0,
			wantRows:     LetterHeight + 1, // +1 for underline row
		},
		{
			name:         "cursor at start normal mode",
			text:         "hi",
			cursorIdx:    0,
			isInsertMode: false,
			maxWidth:     0,
			wantRows:     LetterHeight + 1,
		},
		{
			name:         "cursor at end insert mode",
			text:         "hi",
			cursorIdx:    2,
			isInsertMode: true,
			maxWidth:     0,
			wantRows:     LetterHeight + 1,
		},
		{
			name:         "empty text with cursor",
			text:         "",
			cursorIdx:    0,
			isInsertMode: true,
			maxWidth:     0,
			wantRows:     LetterHeight + 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RenderWithCursor(tt.text, tt.cursorIdx, nil, tt.isInsertMode, tt.maxWidth, DefaultTheme)
			if len(result) != tt.wantRows {
				t.Errorf("RenderWithCursor() returned %d rows, want %d", len(result), tt.wantRows)
			}
		})
	}
}

func TestRenderWithCursorHighlights(t *testing.T) {
	highlights := []CharHighlight{
		HighlightCorrect,
		HighlightIncorrect,
	}

	// Use cursor at -1 to not override highlights
	result := RenderWithCursor("hi", -1, highlights, false, 0, DefaultTheme)

	// Check that we have 7 rows (6 + underline)
	if len(result) != LetterHeight+1 {
		t.Errorf("RenderWithCursor() returned %d rows, want %d", len(result), LetterHeight+1)
	}

	// Check that first character has green color (correct)
	if !strings.Contains(result[0], ANSIGreen) {
		t.Error("First character should have green highlight")
	}

	// Check that second character has red color (incorrect)
	if !strings.Contains(result[0], ANSIRed) {
		t.Error("Second character should have red highlight")
	}
}

func TestRenderWithCursorUnderline(t *testing.T) {
	result := RenderWithCursor("hi", 0, nil, false, 0, DefaultTheme)

	// The underline row is the last row
	underlineRow := result[LetterHeight]

	// Should contain underline characters under cursor
	if !strings.Contains(underlineRow, "▔") {
		t.Error("Underline row should contain underline characters under cursor")
	}
}

func TestRenderWithCursorInsertMode(t *testing.T) {
	result := RenderWithCursor("hi", 1, nil, true, 0, DefaultTheme)

	// Should contain the insert cursor character
	foundInsertCursor := false
	for _, line := range result[:LetterHeight] {
		if strings.Contains(line, "|") {
			foundInsertCursor = true
			break
		}
	}

	if !foundInsertCursor {
		t.Error("Insert mode should show | cursor")
	}
}

func TestRenderPlainText(t *testing.T) {
	result := RenderPlainText("hi", "")

	if len(result) != LetterHeight {
		t.Errorf("RenderPlainText() returned %d rows, want %d", len(result), LetterHeight)
	}
}

func TestRenderPlainTextWithColor(t *testing.T) {
	result := RenderPlainText("hi", ANSIRed)

	// Check that all rows contain the color code
	for i, line := range result {
		if !strings.Contains(line, ANSIRed) {
			t.Errorf("Row %d should contain red color code", i)
		}
	}
}

func TestCalculateTotalWidth(t *testing.T) {
	tests := []struct {
		text         string
		cursorIdx    int
		isInsertMode bool
		wantMin      int // Minimum expected width
	}{
		{
			text:         "hi",
			cursorIdx:    0,
			isInsertMode: false,
			wantMin:      10, // At least some width
		},
		{
			text:         "hi",
			cursorIdx:    1,
			isInsertMode: true,
			wantMin:      11, // Should be slightly wider with insert cursor
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			runes := []rune(tt.text)
			got := CalculateTotalWidth(runes, tt.cursorIdx, tt.isInsertMode)
			if got < tt.wantMin {
				t.Errorf("CalculateTotalWidth() = %d, want at least %d", got, tt.wantMin)
			}
		})
	}
}

func TestRenderWithCursorWordWrap(t *testing.T) {
	// Test with a narrow width that forces wrapping
	result := RenderWithCursor("hello world", 0, nil, false, 40, DefaultTheme)

	// Should produce multiple "rows" of 7 lines each (6 + underline)
	// With wrapping, we expect at least 2 sets
	if len(result) < LetterHeight+1 {
		t.Errorf("Word wrap should produce at least one complete row")
	}
}

func TestGetDisplayWidth(t *testing.T) {
	width := GetDisplayWidth("hello")
	if width <= 0 {
		t.Errorf("GetDisplayWidth() = %d, want > 0", width)
	}

	// Width should increase with more characters
	width2 := GetDisplayWidth("hello world")
	if width2 <= width {
		t.Errorf("Longer text should have greater width: %d <= %d", width2, width)
	}
}
