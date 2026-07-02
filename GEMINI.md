# GoRef Workspace Instructions & Memory

This file serves as the long-lived, team-shared memory and architectural instruction guide for **GoRef**, a fast, local-first Go reference CLI and interactive TUI. It is automatically loaded by the Gemini CLI to ensure context continuity.

---

## 1. Project Vision & Architecture
* **Goal:** A lightning-fast, offline lookup tool for practical Go programming tasks, not a replacement for official docs or a verbose tutorial.
* **Storage:** Markdown reference cards live in a local `refs/` directory and are embedded into the Go executable at compile time using `//go:embed refs/*.md`.
* **Hybrid Loader:** At runtime, GoRef merges the embedded file catalog with the local `refs/` filesystem. If a local file exists with the same name, it overrides the embedded snapshot. This keeps the binary portable while remaining fully user-extensible.

---

## 2. Reference Card Naming & Formatting
To keep cards action-oriented and highly scannable, we enforce strict naming and format constraints.

### Naming Convention
* **Format:** `namespace.action.md`
* **Good Examples:** `map.delete-key.md`, `json.decode-request.md`, `cmd.build.md`, `setup.flat-layout.md`
* **Bad Examples:** `common-commands.md`, `pkg-basics.md`, `delete-key.md`

### Markdown Template
Each card must strictly follow this 5-tier header structure. Keep the `## Card` section short and screen-dense; push detailed explanations to `## Deep`.

```md
# namespace.action

## Card
Quick, screen-dense answer with immediate code snippets.

### Subcard: Small Action
Optional minor related sub-actions.

## Example
A complete, step-by-step, runnable shell walkthrough or Go file example showing expected inputs and outputs. Do not leave this section sparse or empty.

## Deep
In-depth mechanical explanations (garbage collection, mutex safety, stream buffers, etc.).

## Gotchas
Common pitfalls, silent failures, compilation errors, or memory leaks.

## Related
- list
- of
- sibling-references
```

---

## 3. CLI Design & Commands
The GoRef CLI supports smart normalization and query fallbacks:
* **`goref` (no args)**: Runs the TUI by default.
* **`goref list [namespace]`**: Lists all cards grouped dynamically by their namespace prefixes. If a namespace argument is provided, filters to show only cards under that prefix.
* **`goref find <query>`** or **`search`**: Searches for queries inside card names first, falling back to full-text content matches.
* **`goref <space separated arguments>`**: Automatically normalizes space arguments into dotted names (e.g. `goref map delete key` normalizes to `map.delete-key` and opens the card).
* **Topic Section Flags**: Opens specific card sections directly via flags (e.g., `goref show map.delete-key --example` or `goref show setup.project --deep`).
* **Interactive Query Fallback**: If an argument like `goref delete` does not match an exact card name but has matches, GoRef instantly boots the TUI pre-filtered with the query `delete`.

---

## 4. TUI Specifications & Visual Math
The TUI is implemented in `tui.go` using Bubble Tea, Lipgloss, and Glamour. To prevent vertical top-edge clipping and uneven boxes, we use precise mathematical layout coordinates:

### Outer Height Safety Margin
* **Rule:** Bubble Tea's terminal renderer appends a trailing newline (`\n`) to advance the cursor. If box heights match the reported terminal size (`msg.Height`), the viewport scrolls down and cuts off the top borders.
* **Fix:** Set TUI height to `msg.Height - 3` during `WindowSizeMsg` to create a safe layout cushion.

### Pixel-Perfect Bottom Alignment
* **Rule:** Lipgloss's `.Height()` property sets content height. Borders add **2 extra lines** (1 top, 1 bottom) outside this boundary, making boxes taller. Additionally, Glamour's invisible ANSI styling escape codes confuse Lipgloss, causing it to truncate box content and raise the right pane from the bottom.
* **Fix:** 
  1. Remove `.Height()` and `.AlignVertical()` constraints entirely from both panel style builders.
  2. Manually pad the input text strings for both the left list and right content panes to exactly `m.height - 2` lines (with the right pane's markdown content padded to `availableHeight = m.height - 4`).
  3. When Lipgloss renders borders around these identically sized strings, **both panels naturally align pixel-perfectly flush at the bottom** with no truncation!
* **Horizontal wrapping safety:** Pass `rightWidth - 4` to Glamour's text wrap function. This leaves a safe horizontal margin for borders and padding, completely preventing double-wrapping bugs.

### TUI Hotkeys & Navigation
* `j` / `k` (or Arrow keys): Flat list scroll (with compact Visual grouping).
* `/`: Real-time search (matches card names first, then full content). Press `Esc` or `Enter` to lock/exit search edit mode.
* `Tab`: Cycle through namespaces dynamically (filters the list to `cmd.`, `http.`, `map.`, etc.).
* `c` / `d` / `e`: Toggle right-side section between **C**ard, **D**eep, and **E**xample.
* `Enter` or `r`: Opens the **Related Cards** menu. Highlight a related card and press `Enter` to jump straight to it.
* `Backspace`: Navigation history pop (hops back to the previously viewed card).
* `q` / `Ctrl+C`: Exit.

---

## 5. Development & Testing
Run unit tests checking heading parsing, normalization combinations, and related-card extractions before compiling:
```bash
# Run unit test suite
go test -v ./...

# Compile local binary
go build -o goref .

# Test local build
./goref list
```
