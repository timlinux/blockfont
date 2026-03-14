# Animations

blockfont uses spring-based physics for smooth, natural animations.

## Enabling Animations

```go
opts := blockfont.DefaultWidgetOptions()
opts.Animate = true
widget := blockfont.NewWidget(opts)
```

## Triggering Transitions

```go
// Trigger a fade-in animation
cmd := widget.TriggerAnimation(blockfont.TransitionFadeIn)

// Trigger a slide-up animation
cmd := widget.TriggerAnimation(blockfont.TransitionSlideUp)
```

## Transition Types

| Type | Description |
|------|-------------|
| TransitionSlideUp | Slides content upward |
| TransitionSlideDown | Slides content downward |
| TransitionFadeIn | Fades content in |
| TransitionFadeOut | Fades content out |
| TransitionScale | Scales content from small to full |

## Using the Animator Directly

```go
animator := blockfont.NewAnimator()

// Trigger animation
animator.TriggerTransition(blockfont.TransitionSlideUp)

// In your update loop
if animator.Update() {
    // Animation still running
    offset := animator.GetOffset(10)
    opacity := animator.GetOpacityLevel(1.0)
    scale := animator.GetScaleFactor(0.7)
}
```

## Word Carousel

For speed reading applications:

```go
carousel := blockfont.NewWordCarouselAnimator()
carousel.TriggerTransition()

// Get per-word values
prevOffset := carousel.GetPrevOffset()
prevOpacity := carousel.GetPrevOpacity()
currentOffset := carousel.GetCurrentOffset()
currentScale := carousel.GetCurrentScale()
nextOffset := carousel.GetNextOffset()
nextOpacity := carousel.GetNextOpacity()
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
