# GoRef

GoRef is a lightning-fast, offline-first Go reference CLI and interactive TUI. Designed for practical, real-world development tasks, GoRef provides screen-dense, actionable answer cards that keep you in the terminal and out of verbose browser tabs.

---

## 🚀 Key Features

*   **Interactive TUI:** A keyboard-driven, dual-pane Terminal User Interface (TUI) powered by Bubble Tea and Glamour for beautiful markdown rendering.
*   **Flexible CLI:** Normalizes space-separated arguments automatically (`goref map delete key` -> matches `map.delete-key`).
*   **Topic-Specific Viewport Flags:** Jump straight to specific sections of a reference card directly from your terminal (`--card`, `--deep`, `--example`, `--related`).
*   **Dynamic Query Fallback:** If a CLI query doesn't match an exact card name, GoRef boots up the TUI instantly with your query pre-filtered!
*   **Hybrid Content Loader:** Merges a high-speed embedded card catalog with your local `refs/` directory, letting you override or extend references on the fly.

---

## 🛠️ TUI Hotkey & Navigation Guide

When running the interactive TUI (using `goref` or `goref tui`), you can navigate with the following keybindings:

| Hotkey | Action |
| :--- | :--- |
| **`j` / `k`** (or `↑` / `↓`) | Scroll through the flat reference card list. |
| **`/`** | Enter real-time search mode (matches card names first, then falls back to content). |
| **`Esc` / `Enter`** | Exit search/edit mode to return to card navigation. |
| **`Tab`** | Cycle through card namespaces (e.g., `cmd.`, `http.`, `map.`, `ptr.`, etc.). |
| **`c`** | Toggle the right pane to the **C**ard section (quick answers and code snippets). |
| **`d`** | Toggle the right pane to the **D**eep section (in-depth structural/runtime mechanics). |
| **`e`** | Toggle the right pane to the **E**xample section (complete runnable code/scenarios). |
| **`Enter` / `r`** | Open the **Related Cards** menu. Highlight a card and press `Enter` to jump to it. |
| **`Backspace`** | Go back to the previously viewed card (navigation history stack). |
| **`q` / `Ctrl+C`** | Exit GoRef. |

---

## 💻 CLI Usage & Examples

GoRef shines as a quick query tool directly on the command line.

### 1. View Card (Auto-Normalizing)
Separate words with spaces; GoRef automatically normalizes them into dot-and-dash notation:
```bash
# Both of these load the map.delete-key.md reference
goref map delete key
goref show map.delete-key
```

### 2. View Specific Sections
Use viewport flags to output only the section you need to your stdout:
```bash
# Show only the runnable example for setting up a project
goref setup project --example

# Show deep architectural details on slices
goref slice basics --deep
```

### 3. List Reference Cards
List all available reference cards grouped dynamically by namespace:
```bash
# List all cards
goref list

# Filter list to only show the "http" namespace
goref list http
```

### 4. Direct Searching
Search across card names and full-text content:
```bash
goref find decode
```

---

## 🗂️ Reference Card Anatomy

Reference cards live in the `refs/` directory.

### Naming Convention
*   Must be named as `<namespace>.<action>.md`.
*   *Examples:* `map.delete-key.md`, `json.decode-request.md`, `cmd.build.md`, `setup.project.md`.

### Standard Markdown Structure
Every card must adhere strictly to a 5-tier header structure to display correctly in the CLI and TUI:

```markdown
# namespace.action

## Card
Quick, screen-dense answer with immediate code snippets.

### Subcard: Small Action
Optional minor related sub-actions.

## Example
A complete, step-by-step, runnable shell walkthrough or Go file example showing expected inputs and outputs.

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

## ⚙️ How to Setup GoRef Globally

To run `goref` from anywhere in your terminal, you can choose one of the following methods:

### Method A: The Go Way (Recommended)
Go can build and install the binary directly into your `$GOPATH/bin` directory.

1.  Ensure your Go binary path (`$GOPATH/bin` or `~/go/bin`) is included in your system's `$PATH` environment variable.
2.  Run the install command inside the GoRef project root directory:
    ```bash
    go install .
    ```

### Method B: Manual Build & Symlink (or Copy)
You can compile a local binary and link/copy it to a system-wide binary folder (like `/usr/local/bin`):

1.  Compile the binary locally:
    ```bash
    go build -o goref .
    ```
2.  Create a symlink to a directory in your `$PATH` (highly recommended, as updates are automatic when you rebuild locally):
    ```bash
    sudo ln -sf "$(pwd)/goref" /usr/local/bin/goref
    ```
    *Alternatively, copy the binary directly:*
    ```bash
    sudo cp goref /usr/local/bin/goref
    ```

---

## 🔄 Rebuilding & Updating GoRef

Because GoRef embeds reference cards (`refs/*.md`) directly into the executable binary during compilation, **any updates to core Go code or local reference cards require a rebuild to be permanently compiled into the global binary.**

### 1. Instant Runtime Updates (No Rebuild Required!)
Due to GoRef's **hybrid content loader**, if you are working within the GoRef project directory, local cards in the `refs/` directory will **automatically override** the embedded versions. This allows you to edit and test your custom reference cards instantly in the TUI or CLI without rebuilding!

### 2. Updating the Global Binary
When you are ready to persist your changes globally, run the appropriate update command from the GoRef project root:

#### If you installed via `go install`:
Simply run:
```bash
go install .
```
This re-compiles the updated cards/code and instantly replaces the existing global binary in your `$GOPATH/bin`.

#### If you installed via Symlink:
If you used `sudo ln -sf "$(pwd)/goref" /usr/local/bin/goref`, you only need to rebuild the local file:
```bash
go build -o goref .
```
The symlink in `/usr/local/bin/goref` automatically points to this updated binary!

#### If you installed via Manual Copy:
Rebuild and copy the fresh binary over:
```bash
go build -o goref . && sudo cp goref /usr/local/bin/goref
```
