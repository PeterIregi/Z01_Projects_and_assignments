# Ascii-art

A small Go utility that renders text as ASCII art using a banner file (similar to FIGlet-style fonts). The project reads a banner definition from `standard.txt` and prints words as multi-line ASCII characters.

## Features

- Load an ASCII banner file (`standard.txt`) into memory.
- Render input words as ASCII art using the loaded banner.
- Simple CLI in `cmd/` and reusable package in `pkg/` for programmatic use.

## Quick start

1. From the project root, run the CLI with a quoted input string:

   ```sh
   go run ./cmd "Hello"
   ```

   To render multiple words separated by literal `\n` sequences (interpreted by the program as separators), pass them inside the same quoted string. For example:

   ```sh
   go run ./cmd "Hello\nThere"
   ```

   The CLI expects the banner file `standard.txt` to be in the current working directory (the CLI calls `LoadBanner("./standard.txt")`).

## How it works

- `pkg.LoadBanner(filename string) map[rune][]string` reads the banner file and returns a mapping from printable ASCII runes (32..126) to an 8-line slice representing that character's ASCII-art rows.
- `pkg.PrintAscii(banner map[rune][]string, words []string)` prints each word using the banner map. In the CLI the input string is split into words based on literal `\n` or `\t` tokens and passed to `PrintAscii`.

## Usage examples

- Single word:

  ```sh
  go run ./cmd "Hi"
  ```

- Two words separated by a literal `\n` sequence (rendered one after another):

  ```sh
  go run ./cmd "Hello\nThere"
  ```

- If you want to feed text from a file or pipeline, you can call the package functions from your own Go code by loading the banner and calling `PrintAscii`.

## Programmatic example

```go
package main

import (
    "ascii_art/pkg"
)

func main() {
    banner := pkg.LoadBanner("./standard.txt")
    words := []string{"Hello", "There"}
    pkg.PrintAscii(banner, words)
}
```

## Testing

- Run the unit tests for the whole module:

  ```sh
  go test ./...
  ```

## Notes and caveats

- The project expects the banner file to be formatted as a sequence of 9-line blocks per ASCII character (some implementations include a blank line per character). The loader in this repository slices the file in blocks to map printable ASCII characters to their 8-line representations.
- This project is intentionally small and educational. Some tests and CLI/test usage might assume slightly different parameter ordering or behavior; inspect `cmd/main.go` and `pkg/` to understand the exact function signatures used by the CLI.

## Repository layout

- `cmd/` — CLI `main.go` and CLI tests
- `pkg/` — library functions `loadbanner.go`, `print.go` and package tests
- `standard.txt` — ASCII banner definition file used by the loader
- `go.mod` — Go module file
