# Getting Started

## Installation

Install blockfont using `go get`:

```bash
go get github.com/timlinux/blockfont
```

## Quick Example

```go
package main

import (
    "fmt"
    "github.com/timlinux/blockfont"
)

func main() {
    text := blockfont.RenderText("hello")
    fmt.Println(text)
}
```

## Output

```
██  ██ ██████ ██     ██     ◢████◣
██  ██ ██     ██     ██     ██  ██
██████ ████   ██     ██     ██  ██
██  ██ ██     ██     ██     ██  ██
██  ██ ██████ ██████ ██████ ◥████◤
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
