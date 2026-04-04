// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

package blockfont

import (
	"strings"
	"testing"
	"time"
)

func TestNewCharacterSprite(t *testing.T) {
	sprite := NewCharacterSprite()
	if sprite == nil {
		t.Fatal("NewCharacterSprite returned nil")
	}

	// Check all actions have animations
	for _, action := range AllActions() {
		anim, ok := sprite.Animations[action]
		if !ok {
			t.Errorf("Missing animation for action %s", GetActionName(action))
			continue
		}
		if len(anim.Frames) == 0 {
			t.Errorf("No frames for action %s", GetActionName(action))
		}
	}
}

func TestCharacterFrameHeight(t *testing.T) {
	sprite := NewCharacterSprite()

	for _, action := range AllActions() {
		anim := sprite.Animations[action]
		for i, frame := range anim.Frames {
			if len(frame.Lines) != CharacterHeight {
				t.Errorf("Action %s frame %d has %d lines, expected %d",
					GetActionName(action), i, len(frame.Lines), CharacterHeight)
			}
		}
	}
}

func TestCharacterAnimator(t *testing.T) {
	animator := NewCharacterAnimator()

	// Test initial state
	if animator.GetAction() != ActionIdle {
		t.Errorf("Expected initial action to be Idle, got %s", GetActionName(animator.GetAction()))
	}
	if !animator.IsPlaying() {
		t.Error("Expected animator to be playing initially")
	}
	if animator.IsFlipped() {
		t.Error("Expected animator to not be flipped initially")
	}

	// Test action change
	animator.SetAction(ActionWalk)
	if animator.GetAction() != ActionWalk {
		t.Errorf("Expected action to be Walk, got %s", GetActionName(animator.GetAction()))
	}

	// Test flip
	animator.SetFlipped(true)
	if !animator.IsFlipped() {
		t.Error("Expected animator to be flipped after SetFlipped(true)")
	}

	// Test pause/play
	animator.Pause()
	if animator.IsPlaying() {
		t.Error("Expected animator to be paused")
	}
	animator.Play()
	if !animator.IsPlaying() {
		t.Error("Expected animator to be playing")
	}
}

func TestCharacterFrameRender(t *testing.T) {
	sprite := NewCharacterSprite()
	frame := sprite.GetFrame(ActionIdle, 0)

	if frame == nil {
		t.Fatal("GetFrame returned nil")
	}

	rendered := frame.Render()
	lines := strings.Split(rendered, "\n")

	if len(lines) != CharacterHeight {
		t.Errorf("Rendered frame has %d lines, expected %d", len(lines), CharacterHeight)
	}
}

func TestCharacterFrameRenderWithColor(t *testing.T) {
	sprite := NewCharacterSprite()
	frame := sprite.GetFrame(ActionIdle, 0)

	if frame == nil {
		t.Fatal("GetFrame returned nil")
	}

	color := "\033[38;2;255;107;53m"
	rendered := frame.RenderWithColor(color)

	// Check that color codes are present
	if !strings.Contains(rendered, color) {
		t.Error("Rendered frame does not contain color code")
	}
	if !strings.Contains(rendered, ANSIReset) {
		t.Error("Rendered frame does not contain reset code")
	}
}

func TestCharacterAnimatorUpdate(t *testing.T) {
	animator := NewCharacterAnimator()
	animator.SetAction(ActionWalk)

	// Manually set lastUpdate to simulate time passing
	animator.Play()

	// Initial update should not change frame immediately
	_ = animator.Update()
	// Frame may or may not change depending on time elapsed

	// Simulate waiting for frame duration
	time.Sleep(200 * time.Millisecond)
	_ = animator.Update()
	// After sleep, frame should have changed
}

func TestCharacterFlipFrame(t *testing.T) {
	animator := NewCharacterAnimator()
	// Use wave action which has asymmetric frames
	animator.SetAction(ActionWave)

	// Get unflipped frame
	animator.SetFlipped(false)
	unflipped := animator.GetCurrentFrame()

	// Get flipped frame
	animator.SetFlipped(true)
	flipped := animator.GetCurrentFrame()

	if unflipped == nil || flipped == nil {
		t.Fatal("GetCurrentFrame returned nil")
	}

	// Frames should be different (one is reversed)
	unflippedStr := unflipped.Render()
	flippedStr := flipped.Render()

	if unflippedStr == flippedStr {
		t.Error("Flipped frame should be different from unflipped")
	}
}

func TestAllActions(t *testing.T) {
	actions := AllActions()

	expectedActions := []AnimationAction{
		ActionIdle,
		ActionWalk,
		ActionRun,
		ActionDuck,
		ActionJump,
		ActionWave,
	}

	if len(actions) != len(expectedActions) {
		t.Errorf("Expected %d actions, got %d", len(expectedActions), len(actions))
	}

	for i, action := range actions {
		if action != expectedActions[i] {
			t.Errorf("Action %d: expected %v, got %v", i, expectedActions[i], action)
		}
	}
}

func TestGetActionName(t *testing.T) {
	tests := []struct {
		action   AnimationAction
		expected string
	}{
		{ActionIdle, "Idle"},
		{ActionWalk, "Walk"},
		{ActionRun, "Run"},
		{ActionDuck, "Duck"},
		{ActionJump, "Jump"},
		{ActionWave, "Wave"},
	}

	for _, tt := range tests {
		name := GetActionName(tt.action)
		if name != tt.expected {
			t.Errorf("GetActionName(%v) = %s, expected %s", tt.action, name, tt.expected)
		}
	}
}

func TestGetFrameCount(t *testing.T) {
	sprite := NewCharacterSprite()

	tests := []struct {
		action        AnimationAction
		expectedCount int
	}{
		{ActionIdle, 2},
		{ActionWalk, 4},
		{ActionRun, 4},
		{ActionDuck, 1},
		{ActionJump, 3},
		{ActionWave, 3},
	}

	for _, tt := range tests {
		count := sprite.GetFrameCount(tt.action)
		if count != tt.expectedCount {
			t.Errorf("GetFrameCount(%s) = %d, expected %d",
				GetActionName(tt.action), count, tt.expectedCount)
		}
	}
}

func TestCharacterAnimatorRender(t *testing.T) {
	animator := NewCharacterAnimator()

	rendered := animator.Render()
	if rendered == "" {
		t.Error("Render returned empty string")
	}

	colorRendered := animator.RenderWithColor("\033[32m")
	if colorRendered == "" {
		t.Error("RenderWithColor returned empty string")
	}
	if !strings.Contains(colorRendered, "\033[32m") {
		t.Error("RenderWithColor does not contain color code")
	}
}
