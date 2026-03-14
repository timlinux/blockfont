// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

package blockfont

import (
	"strings"
	"testing"
)

func TestCenterLines(t *testing.T) {
	lines := []string{"hello", "hi"}
	result := CenterLines(lines, 10)

	// "hello" (5 chars) in width 10 should have 2 spaces padding
	if !strings.HasPrefix(result[0], "  ") {
		t.Errorf("CenterLines did not center 'hello' correctly: %q", result[0])
	}

	// "hi" (2 chars) in width 10 should have 4 spaces padding
	if !strings.HasPrefix(result[1], "    ") {
		t.Errorf("CenterLines did not center 'hi' correctly: %q", result[1])
	}
}

func TestLeftJustify(t *testing.T) {
	lines := []string{"hello", "world"}
	result := LeftJustify(lines, 3)

	for i, line := range result {
		if !strings.HasPrefix(line, "   ") {
			t.Errorf("LeftJustify line %d = %q, expected 3-space prefix", i, line)
		}
	}
}

func TestRightJustify(t *testing.T) {
	lines := []string{"hi"}
	result := RightJustify(lines, 10)

	// "hi" (2 chars) in width 10 should have 8 spaces padding
	if !strings.HasPrefix(result[0], "        ") {
		t.Errorf("RightJustify did not right-justify correctly: %q", result[0])
	}
}

func TestWrapOnWordBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		maxWidth int
		wantLen  int
	}{
		{
			name:     "single word fits",
			text:     "hello",
			maxWidth: 100,
			wantLen:  1,
		},
		{
			name:     "wraps on space",
			text:     "hello world",
			maxWidth: 40,
			wantLen:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WrapOnWordBoundaries(tt.text, tt.maxWidth)
			if len(result) != tt.wantLen {
				t.Errorf("WrapOnWordBoundaries() returned %d lines, want %d", len(result), tt.wantLen)
			}
		})
	}
}

func TestPadToWidth(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		expected int
	}{
		{"hello", 10, 10},
		{"hi", 5, 5},
		{"test", 4, 4},
	}

	for _, tt := range tests {
		result := PadToWidth(tt.input, tt.width)
		got := VisibleStringWidth(result)
		if got != tt.expected {
			t.Errorf("PadToWidth(%q, %d) visible width = %d, want %d", tt.input, tt.width, got, tt.expected)
		}
	}
}

func TestMaxLineWidth(t *testing.T) {
	lines := []string{"short", "a bit longer", "hi"}
	got := MaxLineWidth(lines)
	want := 12 // "a bit longer"
	if got != want {
		t.Errorf("MaxLineWidth() = %d, want %d", got, want)
	}
}

func TestAlignLines(t *testing.T) {
	lines := []string{"hi"}

	// Test center alignment
	centered := AlignLines(lines, AlignCenter, 10)
	if !strings.HasPrefix(centered[0], "    ") {
		t.Errorf("AlignLines(Center) = %q, expected centered", centered[0])
	}

	// Test right alignment
	rightAligned := AlignLines(lines, AlignRight, 10)
	if !strings.HasPrefix(rightAligned[0], "        ") {
		t.Errorf("AlignLines(Right) = %q, expected right aligned", rightAligned[0])
	}

	// Test left alignment (no change)
	leftAligned := AlignLines(lines, AlignLeft, 10)
	if leftAligned[0] != "hi" {
		t.Errorf("AlignLines(Left) = %q, expected no change", leftAligned[0])
	}
}
