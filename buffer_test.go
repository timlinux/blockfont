// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

package blockfont

import (
	"testing"
)

func TestNewBuffer(t *testing.T) {
	b := NewBuffer("hello\nworld")
	if b.Text() != "hello\nworld" {
		t.Errorf("NewBuffer text = %q, want %q", b.Text(), "hello\nworld")
	}
	if len(b.Lines()) != 2 {
		t.Errorf("NewBuffer lines = %d, want %d", len(b.Lines()), 2)
	}
}

func TestBufferInsert(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		insert   string
		cursorX  int
		cursorY  int
		expected string
	}{
		{
			name:     "insert at start",
			initial:  "world",
			insert:   "hello ",
			cursorX:  0,
			cursorY:  0,
			expected: "hello world",
		},
		{
			name:     "insert at end",
			initial:  "hello",
			insert:   " world",
			cursorX:  5,
			cursorY:  0,
			expected: "hello world",
		},
		{
			name:     "insert in middle",
			initial:  "helo",
			insert:   "l",
			cursorX:  3,
			cursorY:  0,
			expected: "hello",
		},
		{
			name:     "insert newline",
			initial:  "helloworld",
			insert:   "\n",
			cursorX:  5,
			cursorY:  0,
			expected: "hello\nworld",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuffer(tt.initial)
			b.SetMode(ModeInsert)
			b.SetCursorPosition(tt.cursorX, tt.cursorY)
			b.Insert(tt.insert)
			if b.Text() != tt.expected {
				t.Errorf("Insert() = %q, want %q", b.Text(), tt.expected)
			}
		})
	}
}

func TestBufferDelete(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		count    int
		cursorX  int
		cursorY  int
		expected string
		deleted  string
	}{
		{
			name:     "delete single char",
			initial:  "hello",
			count:    1,
			cursorX:  0,
			cursorY:  0,
			expected: "ello",
			deleted:  "h",
		},
		{
			name:     "delete multiple chars",
			initial:  "hello",
			count:    3,
			cursorX:  0,
			cursorY:  0,
			expected: "lo",
			deleted:  "hel",
		},
		{
			name:     "delete at middle",
			initial:  "hello",
			count:    2,
			cursorX:  2,
			cursorY:  0,
			expected: "heo",
			deleted:  "ll",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuffer(tt.initial)
			b.SetCursorPosition(tt.cursorX, tt.cursorY)
			deleted := b.Delete(tt.count)
			if b.Text() != tt.expected {
				t.Errorf("Delete() text = %q, want %q", b.Text(), tt.expected)
			}
			if deleted != tt.deleted {
				t.Errorf("Delete() deleted = %q, want %q", deleted, tt.deleted)
			}
		})
	}
}

func TestBufferCursorIndex(t *testing.T) {
	b := NewBuffer("hello\nworld")

	tests := []struct {
		cursorX int
		cursorY int
		want    int
	}{
		{0, 0, 0},
		{5, 0, 5},
		{0, 1, 6},
		{5, 1, 11},
	}

	for _, tt := range tests {
		b.SetMode(ModeInsert)
		b.SetCursorPosition(tt.cursorX, tt.cursorY)
		got := b.CursorIndex()
		if got != tt.want {
			t.Errorf("CursorIndex() at (%d,%d) = %d, want %d", tt.cursorX, tt.cursorY, got, tt.want)
		}
	}
}

func TestBufferMode(t *testing.T) {
	b := NewBuffer("test")

	if b.Mode() != ModeNormal {
		t.Errorf("Initial mode = %v, want %v", b.Mode(), ModeNormal)
	}

	b.SetMode(ModeInsert)
	if b.Mode() != ModeInsert {
		t.Errorf("After SetMode(Insert) = %v, want %v", b.Mode(), ModeInsert)
	}

	if !b.IsInsertMode() {
		t.Error("IsInsertMode() = false, want true")
	}
}

func TestBufferMovement(t *testing.T) {
	b := NewBuffer("hello\nworld")
	b.SetMode(ModeInsert)

	// Move right
	b.MoveRight(3)
	x, y := b.CursorPosition()
	if x != 3 || y != 0 {
		t.Errorf("After MoveRight(3) = (%d,%d), want (3,0)", x, y)
	}

	// Move down
	b.MoveDown(1)
	x, y = b.CursorPosition()
	if x != 3 || y != 1 {
		t.Errorf("After MoveDown(1) = (%d,%d), want (3,1)", x, y)
	}

	// Move to line start
	b.MoveToLineStart()
	x, _ = b.CursorPosition()
	if x != 0 {
		t.Errorf("After MoveToLineStart() x = %d, want 0", x)
	}

	// Move to line end
	b.MoveToLineEnd()
	x, _ = b.CursorPosition()
	if x != 5 {
		t.Errorf("After MoveToLineEnd() x = %d, want 5", x)
	}
}

func TestModeString(t *testing.T) {
	tests := []struct {
		mode Mode
		want string
	}{
		{ModeNormal, "NORMAL"},
		{ModeInsert, "INSERT"},
		{ModeVisual, "VISUAL"},
		{ModeVisualLine, "V-LINE"},
		{ModeVisualBlock, "V-BLOCK"},
		{ModeCommand, "COMMAND"},
	}

	for _, tt := range tests {
		if got := tt.mode.String(); got != tt.want {
			t.Errorf("%v.String() = %q, want %q", tt.mode, got, tt.want)
		}
	}
}
