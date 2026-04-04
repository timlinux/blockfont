// SPDX-FileCopyrightText: 2026 Tim Sutton / Kartoza
// SPDX-License-Identifier: MIT

package blockfont

import (
	"strings"
	"time"
)

// CharacterHeight is the height of character sprites (2x font height)
const CharacterHeight = 12

// CharacterWidth is the typical width of character sprites
const CharacterWidth = 12

// AnimationAction represents different character actions
type AnimationAction int

const (
	// ActionIdle is the standing still pose
	ActionIdle AnimationAction = iota
	// ActionWalk is walking animation
	ActionWalk
	// ActionRun is running animation
	ActionRun
	// ActionDuck is crouching/ducking pose
	ActionDuck
	// ActionJump is jumping animation
	ActionJump
	// ActionWave is waving/greeting animation
	ActionWave
)

// CharacterFrame represents a single animation frame
type CharacterFrame struct {
	Lines []string
}

// CharacterAnimation holds all frames for an action
type CharacterAnimation struct {
	Frames   []CharacterFrame
	Duration time.Duration // time per frame
}

// CharacterSprite holds all animations for a character
type CharacterSprite struct {
	Animations map[AnimationAction]*CharacterAnimation
}

// Character frames using block elements (◢ ◣ ◤ ◥ ██)
// Each frame is 12 lines tall (2x the font height)
// Tall thin style - side profile like "La Linea" cartoon
var (
	// Idle pose - standing still
	idleFrames = []CharacterFrame{
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
		}},
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
		}},
	}

	// Walking animation - 4 frames
	walkFrames = []CharacterFrame{
		// Frame 1 - legs apart, arm forward
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"   ██████◣  ",
			"    ████ ◥◣ ",
			"    ████    ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
			"  ◢◤    ██  ",
		}},
		// Frame 2 - legs crossing
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
		}},
		// Frame 3 - legs apart other way, arm back
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"  ◢◣████    ",
			" ◢◤ ████    ",
			"    ████    ",
			"     ██     ",
			"    ◢◣◥◣    ",
			"   ◢◤  ◥◣   ",
			"   ██   ◥◣  ",
		}},
		// Frame 4 - legs crossing
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
		}},
	}

	// Running animation - 4 frames
	runFrames = []CharacterFrame{
		// Frame 1 - big stride, leaning forward
		{Lines: []string{
			"            ",
			"     ◢██◣   ",
			"     ████   ",
			"     ◥██◤   ",
			"      ██    ",
			"   ◢██████◣ ",
			"     ████ ◥◣",
			"     ████   ",
			"      ██    ",
			"    ◢◤ ◥◣   ",
			"   ◢◤   ◥◣  ",
			"  ◢◤     ██ ",
		}},
		// Frame 2 - flight phase
		{Lines: []string{
			"            ",
			"     ◢██◣   ",
			"     ████   ",
			"     ◥██◤   ",
			"      ██    ",
			"    ██████◣ ",
			"     ████◥◣ ",
			"     ████   ",
			"      ██    ",
			"     ◢◤◥◣   ",
			"    ◢◤  ◥◣  ",
			"            ",
		}},
		// Frame 3 - big stride other leg
		{Lines: []string{
			"            ",
			"     ◢██◣   ",
			"     ████   ",
			"     ◥██◤   ",
			"      ██    ",
			"  ◢◣██████  ",
			" ◢◤  ████   ",
			"     ████   ",
			"      ██    ",
			"    ◢◣ ◥◣   ",
			"   ◢◤   ◥◣  ",
			"   ██    ◥◣ ",
		}},
		// Frame 4 - flight phase
		{Lines: []string{
			"            ",
			"     ◢██◣   ",
			"     ████   ",
			"     ◥██◤   ",
			"      ██    ",
			"  ◢◣██████  ",
			" ◢◤  ████   ",
			"     ████   ",
			"      ██    ",
			"     ◢◤◥◣   ",
			"    ◢◤  ◥◣  ",
			"            ",
		}},
	}

	// Duck/crouch pose
	duckFrames = []CharacterFrame{
		{Lines: []string{
			"            ",
			"            ",
			"            ",
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"    ████    ",
			"   ██████   ",
			"    ████    ",
			"   ◢◤  ◥◣   ",
			"  ◢◤    ◥◣  ",
		}},
	}

	// Jump animation - 3 frames
	jumpFrames = []CharacterFrame{
		// Crouch
		{Lines: []string{
			"            ",
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
			"  ◢◤    ◥◣  ",
		}},
		// Peak - arms up
		{Lines: []string{
			"    ◢◣◢◣    ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"            ",
			"            ",
		}},
		// Landing
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤    ",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"   ◢◤  ◥◣   ",
			"  ◢◤    ◥◣  ",
			" ◢◤      ◥◣ ",
		}},
	}

	// Wave animation - 3 frames
	waveFrames = []CharacterFrame{
		// Arm up
		{Lines: []string{
			"         ◢◣ ",
			"    ◢██◣ ██ ",
			"    ████◢█◤ ",
			"    ◥██████ ",
			"     ██ ◥◤  ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
		}},
		// Arm middle
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████◢███",
			"    ◥██████◤",
			"     ██     ",
			"    ████    ",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
		}},
		// Arm down
		{Lines: []string{
			"            ",
			"    ◢██◣    ",
			"    ████    ",
			"    ◥██◤◥◣  ",
			"     ██  ◥◣ ",
			"    ████  ██",
			"    ████    ",
			"    ████    ",
			"     ██     ",
			"     ██     ",
			"    ◢◤◥◣    ",
			"   ◢◤  ◥◣   ",
		}},
	}
)

// NewCharacterSprite creates a new character sprite with all animations
func NewCharacterSprite() *CharacterSprite {
	return &CharacterSprite{
		Animations: map[AnimationAction]*CharacterAnimation{
			ActionIdle: {
				Frames:   idleFrames,
				Duration: 500 * time.Millisecond,
			},
			ActionWalk: {
				Frames:   walkFrames,
				Duration: 150 * time.Millisecond,
			},
			ActionRun: {
				Frames:   runFrames,
				Duration: 100 * time.Millisecond,
			},
			ActionDuck: {
				Frames:   duckFrames,
				Duration: 200 * time.Millisecond,
			},
			ActionJump: {
				Frames:   jumpFrames,
				Duration: 150 * time.Millisecond,
			},
			ActionWave: {
				Frames:   waveFrames,
				Duration: 200 * time.Millisecond,
			},
		},
	}
}

// GetFrame returns the animation frame at the given index for an action
func (cs *CharacterSprite) GetFrame(action AnimationAction, frameIndex int) *CharacterFrame {
	anim, ok := cs.Animations[action]
	if !ok {
		return nil
	}
	if len(anim.Frames) == 0 {
		return nil
	}
	idx := frameIndex % len(anim.Frames)
	return &anim.Frames[idx]
}

// GetFrameCount returns the number of frames for an action
func (cs *CharacterSprite) GetFrameCount(action AnimationAction) int {
	anim, ok := cs.Animations[action]
	if !ok {
		return 0
	}
	return len(anim.Frames)
}

// GetFrameDuration returns the duration per frame for an action
func (cs *CharacterSprite) GetFrameDuration(action AnimationAction) time.Duration {
	anim, ok := cs.Animations[action]
	if !ok {
		return 100 * time.Millisecond
	}
	return anim.Duration
}

// RenderFrame renders a character frame as a string
func (cf *CharacterFrame) Render() string {
	return strings.Join(cf.Lines, "\n")
}

// RenderFrameWithColor renders a frame with ANSI color
func (cf *CharacterFrame) RenderWithColor(color string) string {
	var result strings.Builder
	for i, line := range cf.Lines {
		result.WriteString(color)
		result.WriteString(line)
		result.WriteString(ANSIReset)
		if i < len(cf.Lines)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}

// CharacterAnimator handles character animation state
type CharacterAnimator struct {
	sprite       *CharacterSprite
	action       AnimationAction
	currentFrame int
	lastUpdate   time.Time
	isPlaying    bool
	looping      bool
	flipped      bool // horizontal flip for facing direction
}

// NewCharacterAnimator creates a new character animator
func NewCharacterAnimator() *CharacterAnimator {
	return &CharacterAnimator{
		sprite:       NewCharacterSprite(),
		action:       ActionIdle,
		currentFrame: 0,
		lastUpdate:   time.Now(),
		isPlaying:    true,
		looping:      true,
		flipped:      false,
	}
}

// SetAction changes the current animation action
func (ca *CharacterAnimator) SetAction(action AnimationAction) {
	if ca.action != action {
		ca.action = action
		ca.currentFrame = 0
		ca.lastUpdate = time.Now()
	}
}

// GetAction returns the current action
func (ca *CharacterAnimator) GetAction() AnimationAction {
	return ca.action
}

// SetFlipped sets whether the character is flipped horizontally
func (ca *CharacterAnimator) SetFlipped(flipped bool) {
	ca.flipped = flipped
}

// IsFlipped returns whether the character is flipped
func (ca *CharacterAnimator) IsFlipped() bool {
	return ca.flipped
}

// SetLooping sets whether the animation should loop
func (ca *CharacterAnimator) SetLooping(looping bool) {
	ca.looping = looping
}

// Play starts the animation
func (ca *CharacterAnimator) Play() {
	ca.isPlaying = true
}

// Pause pauses the animation
func (ca *CharacterAnimator) Pause() {
	ca.isPlaying = false
}

// IsPlaying returns whether the animation is playing
func (ca *CharacterAnimator) IsPlaying() bool {
	return ca.isPlaying
}

// Update advances the animation based on elapsed time
// Returns true if a frame change occurred
func (ca *CharacterAnimator) Update() bool {
	if !ca.isPlaying {
		return false
	}

	elapsed := time.Since(ca.lastUpdate)
	frameDuration := ca.sprite.GetFrameDuration(ca.action)
	frameCount := ca.sprite.GetFrameCount(ca.action)

	if elapsed >= frameDuration {
		ca.lastUpdate = time.Now()
		if ca.looping {
			ca.currentFrame = (ca.currentFrame + 1) % frameCount
		} else if ca.currentFrame < frameCount-1 {
			ca.currentFrame++
		} else {
			ca.isPlaying = false
		}
		return true
	}
	return false
}

// GetCurrentFrame returns the current animation frame
func (ca *CharacterAnimator) GetCurrentFrame() *CharacterFrame {
	frame := ca.sprite.GetFrame(ca.action, ca.currentFrame)
	if frame == nil {
		return nil
	}
	if ca.flipped {
		return ca.flipFrame(frame)
	}
	return frame
}

// flipFrame creates a horizontally flipped version of a frame
func (ca *CharacterAnimator) flipFrame(frame *CharacterFrame) *CharacterFrame {
	flipped := &CharacterFrame{
		Lines: make([]string, len(frame.Lines)),
	}
	for i, line := range frame.Lines {
		runes := []rune(line)
		// Reverse the runes
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		// Swap triangle characters
		for j, r := range runes {
			switch r {
			case '◢':
				runes[j] = '◣'
			case '◣':
				runes[j] = '◢'
			case '◤':
				runes[j] = '◥'
			case '◥':
				runes[j] = '◤'
			}
		}
		flipped.Lines[i] = string(runes)
	}
	return flipped
}

// Render returns the current frame as a string
func (ca *CharacterAnimator) Render() string {
	frame := ca.GetCurrentFrame()
	if frame == nil {
		return ""
	}
	return frame.Render()
}

// RenderWithColor returns the current frame with ANSI color
func (ca *CharacterAnimator) RenderWithColor(color string) string {
	frame := ca.GetCurrentFrame()
	if frame == nil {
		return ""
	}
	return frame.RenderWithColor(color)
}

// GetActionName returns a human-readable name for an action
func GetActionName(action AnimationAction) string {
	switch action {
	case ActionIdle:
		return "Idle"
	case ActionWalk:
		return "Walk"
	case ActionRun:
		return "Run"
	case ActionDuck:
		return "Duck"
	case ActionJump:
		return "Jump"
	case ActionWave:
		return "Wave"
	default:
		return "Unknown"
	}
}

// AllActions returns all available animation actions
func AllActions() []AnimationAction {
	return []AnimationAction{
		ActionIdle,
		ActionWalk,
		ActionRun,
		ActionDuck,
		ActionJump,
		ActionWave,
	}
}
