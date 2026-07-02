# setup.global-path

## Card

Add your project's compiled binary directory to your system's `$PATH` so it can be executed globally from any directory.

```fish
# In Fish Shell (Instantly persistent, safe absolute path)
fish_add_path $PWD
```

```bash
# In Bash / Zsh (Persistent, append to config)
echo 'export PATH="$PATH:'$PWD'"' >> ~/.zshrc
source ~/.zshrc
```

### Subcard: Verify Active Binary

```bash
# Shows the exact path of the binary being executed
which goref
```

## Example

### Complete Walkthrough for Fish Shell (Active Development)

1. Open your terminal and navigate to your GoRef project directory:
```fish
cd /Users/jn/Documents/GOREF
```

2. Compile your binary locally (overwriting the previous version):
```fish
go build -o goref .
```

3. Register your current working directory persistently into your Fish `$PATH` lookup list using `$PWD`:
```fish
fish_add_path $PWD
```

4. Verify that Fish now correctly identifies your project directory for `goref` lookups:
```fish
which goref
# Expected Output:
# /Users/jn/Documents/GOREF/goref
```

5. You can now navigate to **any directory** (like your home folder) and run your local development binary globally:
```fish
cd ~
goref list
```

---

### Complete Walkthrough for Bash / Zsh

1. Navigate to the project directory and build:
```bash
cd /Users/jn/Documents/GOREF
go build -o goref .
```

2. Append the absolute project path to your `.zshrc` or `.bashrc` profile:
```bash
# Adds the absolute path of the current directory to your shell configuration
echo "export PATH=\"\$PATH:$PWD\"" >> ~/.zshrc
```

3. Reload your shell profile to apply changes instantly:
```bash
source ~/.zshrc
```

4. Run the command globally:
```bash
cd ~
goref list
```

## Deep

When you type a command (like `goref`) and press Enter, your shell doesn't search your entire filesystem. It uses **Command Resolution**:
- The shell iterates through a list of folders defined in the **`$PATH`** environment variable from left to right.
- The moment it finds an executable file matching `goref` inside one of those folders, it stops searching and executes it.
- By binding `$PWD` persistently to your path, your shell points directly to the `/Users/jn/Documents/GOREF` development directory. Because Go's compiler overwrites the binary in-place when you compile, your global command references the updated machine code instantly.

## Gotchas

- **Never use relative paths (e.g. `fish_add_path .`):** Adding the dot `.` tells your shell to search for executables in whatever directory you are *currently* standing in.
  - **Path Errors:** If you `cd` to your home folder, `goref` will stop working because there is no `goref` binary in your home folder.
  - **Severe Security Risk:** If you navigate to an untrusted directory (e.g. an open source project you just cloned) containing a malicious script named `ls` or `cd`, and you type that command, your shell will execute the local malicious script instead of the system binary! **Always bind absolute paths (like `$PWD` or `/Users/jn/Documents/GOREF`), never relative paths (like `.`).**

## Related

- cmd.build
- setup.project
