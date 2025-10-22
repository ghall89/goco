# goco

A small utility for making and pushing git commits, written in Go.

## Building from Source

This Readme assumes you're using [task](https://taskfile.dev/docs/installation).

### Install Dependencies

Before building, ensure you have all dependencies installed with `task install-deps`.

### Building

Simply run `task build` to build. The binaries will appear in the `out/` directory.

You can also build for your specific platform with:

- `task build-darwin` for macOS
- `task build-linux` for Linux
- `task build-windows` for Windows

## Dependencies

- [go-git](https://github.com/go-git/go-git)
- [huh](https://github.com/charmbracelet/huh)
- [lipgloss](https://github.com/charmbracelet/lipgloss)
