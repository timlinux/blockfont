# Contributing

We welcome contributions to blockfont!

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/timlinux/blockfont.git
cd blockfont
```

2. Enter the Nix development shell:
```bash
nix develop
```

3. Install pre-commit hooks:
```bash
pre-commit install
```

## Running Tests

```bash
go test -v ./...
```

## Building

```bash
go build ./...
```

## Code Style

- Follow standard Go conventions
- Run `gofmt` before committing
- Add tests for new functionality
- Document public APIs with godoc comments

## Pull Request Process

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linting
5. Submit a PR with clear description

## Adding New Characters

To add a new character to the font:

1. Edit `font.go`
2. Add entry to `BlockLetters` map
3. Character must be 6 lines tall
4. Use ASCII block characters: █ ◢ ◣ ◤ ◥

Example:
```go
'★': {
    "  ██  ",
    " ◢██◣ ",
    "██████",
    " ◥██◤ ",
    " ◢██◣ ",
    "      ",
},
```

---

Made with :heart: by [Kartoza](https://kartoza.com) | [Donate!](https://github.com/sponsors/timlinux) | [GitHub](https://github.com/timlinux/blockfont)
