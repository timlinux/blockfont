// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

package blockfont

import (
	"strings"
	"testing"
)

func TestRenderWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantLen  int
		wantRows int
	}{
		{
			name:     "empty string",
			input:    "",
			wantLen:  0,
			wantRows: LetterHeight,
		},
		{
			name:     "single letter",
			input:    "a",
			wantLen:  1,
			wantRows: LetterHeight,
		},
		{
			name:     "word",
			input:    "hello",
			wantLen:  5,
			wantRows: LetterHeight,
		},
		{
			name:     "with space",
			input:    "hi there",
			wantLen:  8,
			wantRows: LetterHeight,
		},
		{
			name:     "uppercase",
			input:    "ABC",
			wantLen:  3,
			wantRows: LetterHeight,
		},
		{
			name:     "mixed case",
			input:    "Hello",
			wantLen:  5,
			wantRows: LetterHeight,
		},
		{
			name:     "with numbers",
			input:    "test123",
			wantLen:  7,
			wantRows: LetterHeight,
		},
		{
			name:     "punctuation",
			input:    "hi!",
			wantLen:  3,
			wantRows: LetterHeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RenderWord(tt.input)
			if len(result) != tt.wantRows {
				t.Errorf("RenderWord() returned %d rows, want %d", len(result), tt.wantRows)
			}
			if len(result) > 0 && len(result[0]) != tt.wantLen {
				t.Errorf("RenderWord() returned %d letters, want %d", len(result[0]), tt.wantLen)
			}
		})
	}
}

func TestRenderText(t *testing.T) {
	result := RenderText("hi")
	lines := strings.Split(result, "\n")
	if len(lines) != LetterHeight {
		t.Errorf("RenderText() returned %d lines, want %d", len(lines), LetterHeight)
	}
}

func TestGetLetterWidth(t *testing.T) {
	tests := []struct {
		char rune
		want int
	}{
		{'a', 6},
		{' ', 3},
		{'w', 7},
		{'@', 7},
	}

	for _, tt := range tests {
		t.Run(string(tt.char), func(t *testing.T) {
			got := GetLetterWidth(tt.char)
			if got != tt.want {
				t.Errorf("GetLetterWidth(%q) = %d, want %d", tt.char, got, tt.want)
			}
		})
	}
}

func TestGetTotalWidth(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"a", 6},
		{"ab", 14}, // 6 + 1 (spacing) + 7 (b has trailing space)
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := GetTotalWidth(tt.input)
			if got != tt.want {
				t.Errorf("GetTotalWidth(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestBlockLettersCompleteness(t *testing.T) {
	// Test that all expected characters are present
	expected := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 .,;:!?-'\""

	for _, r := range expected {
		if _, ok := BlockLetters[r]; !ok {
			t.Errorf("BlockLetters missing character: %q", r)
		}
	}
}

func TestBlockLettersHeight(t *testing.T) {
	// Test that all characters have correct height
	for r, lines := range BlockLetters {
		if len(lines) != LetterHeight {
			t.Errorf("BlockLetters[%q] has %d lines, want %d", r, len(lines), LetterHeight)
		}
	}
}
