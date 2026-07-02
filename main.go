package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

//go:embed refs/*.md
var refs embed.FS

// getCardContent loads the markdown content for a given card name.
// It checks the local refs/ folder first, then falls back to the embedded FS.
func getCardContent(name string) ([]byte, error) {
	localPath := filepath.Join("refs", name+".md")
	if data, err := os.ReadFile(localPath); err == nil {
		return data, nil
	}
	return refs.ReadFile("refs/" + name + ".md")
}

// getCardNames returns a sorted slice of all unique card names
// from both the local refs/ folder and the embedded FS.
func getCardNames() ([]string, error) {
	nameMap := make(map[string]bool)

	// Read embedded files
	if entries, err := refs.ReadDir("refs"); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
				nameMap[strings.TrimSuffix(entry.Name(), ".md")] = true
			}
		}
	}

	// Read local refs directory
	if entries, err := os.ReadDir("refs"); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
				nameMap[strings.TrimSuffix(entry.Name(), ".md")] = true
			}
		}
	}

	names := make([]string, 0, len(nameMap))
	for name := range nameMap {
		names = append(names, name)
	}
	sort.Strings(names)
	return names, nil
}

// getCardTitle extracts the top-level "# <Title>" from the markdown content.
func getCardTitle(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "# ") {
			return trimmed
		}
	}
	return ""
}

// extractSection parses a markdown string and extracts the content under a specific ## heading.
func extractSection(content, section string) (string, error) {
	lines := strings.Split(content, "\n")
	target := strings.ToLower(strings.TrimSpace(section))

	inSection := false
	startLevel := 0
	var out []string

	for _, line := range lines {
		level, title, ok := parseHeading(line)

		if ok {
			normalizedTitle := strings.ToLower(strings.TrimSpace(title))

			if !inSection && normalizedTitle == target {
				inSection = true
				startLevel = level
				continue
			}

			if inSection && level <= startLevel {
				break
			}
		}

		if inSection {
			out = append(out, line)
		}
	}

	result := strings.TrimSpace(strings.Join(out, "\n"))
	if result == "" {
		return "", fmt.Errorf("section not found: %s", section)
	}

	return result, nil
}

// parseHeading parses a line to see if it is a markdown heading.
func parseHeading(line string) (level int, title string, ok bool) {
	trimmed := strings.TrimSpace(line)

	if !strings.HasPrefix(trimmed, "#") {
		return 0, "", false
	}

	for level < len(trimmed) && trimmed[level] == '#' {
		level++
	}

	if level == 0 || level > 6 {
		return 0, "", false
	}

	if len(trimmed) <= level || trimmed[level] != ' ' {
		return 0, "", false
	}

	title = strings.TrimSpace(trimmed[level:])
	return level, title, true
}

// normalizeTopic normalizes separate command arguments into a card name.
// E.g., ["setup", "project"] -> "setup.project"
// E.g., ["map", "delete", "key"] -> "map.delete-key"
func normalizeTopic(parts []string) string {
	if len(parts) == 0 {
		return ""
	}

	// If there is only one part and it contains a dot, it's already in namespace.action format
	if len(parts) == 1 && strings.Contains(parts[0], ".") {
		return parts[0]
	}

	// Check if we can find an exact match in our existing card names by trying combinations
	names, err := getCardNames()
	if err == nil {
		joinedSpace := strings.Join(parts, " ")
		joinedDot := strings.Join(parts, ".")
		joinedDash := strings.Join(parts, "-")

		// Try namespace.action style: first part as namespace, rest joined with dash
		var joinedNamespaceAction string
		if len(parts) > 1 {
			joinedNamespaceAction = parts[0] + "." + strings.Join(parts[1:], "-")
		}

		for _, name := range names {
			lname := strings.ToLower(name)
			if lname == strings.ToLower(joinedDot) ||
				lname == strings.ToLower(joinedNamespaceAction) ||
				lname == strings.ToLower(joinedSpace) ||
				lname == strings.ToLower(joinedDash) {
				return name
			}
		}
	}

	// Fallback: join first part with dot, subsequent parts with dashes
	if len(parts) > 1 {
		return parts[0] + "." + strings.Join(parts[1:], "-")
	}
	return parts[0]
}

func cardExists(name string) bool {
	names, err := getCardNames()
	if err != nil {
		return false
	}
	for _, n := range names {
		if strings.EqualFold(n, name) {
			return true
		}
	}
	return false
}

func hasMatchingCards(query string) bool {
	names, err := getCardNames()
	if err != nil {
		return false
	}
	query = strings.ToLower(query)
	for _, n := range names {
		if strings.Contains(strings.ToLower(n), query) {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		if err := runTUI(""); err != nil {
			fmt.Println("Error:", err)
		}
		return
	}

	nonFlags, section := parseArgs(args)

	if len(nonFlags) == 0 {
		help()
		return
	}

	cmd := nonFlags[0]

	switch cmd {
	case "tui":
		query := ""
		if len(nonFlags) > 1 {
			query = strings.Join(nonFlags[1:], " ")
		}
		if err := runTUI(query); err != nil {
			fmt.Println("Error:", err)
		}
		return

	case "list":
		filter := ""
		if len(nonFlags) > 1 {
			filter = strings.Join(nonFlags[1:], " ")
		}
		listRefs(filter)
		return

	case "search", "find":
		if len(nonFlags) < 2 {
			fmt.Printf("usage: goref %s <query>\n", cmd)
			return
		}
		searchRefs(strings.Join(nonFlags[1:], " "))
		return

	case "show":
		if len(nonFlags) < 2 {
			fmt.Println("usage: goref show <topic> [flags]")
			return
		}
		topic := normalizeTopic(nonFlags[1:])
		showSection(topic, section)
		return

	case "deep":
		if len(nonFlags) < 2 {
			fmt.Println("usage: goref deep <topic>")
			return
		}
		topic := normalizeTopic(nonFlags[1:])
		showSection(topic, "Deep")
		return

	case "example":
		if len(nonFlags) < 2 {
			fmt.Println("usage: goref example <topic>")
			return
		}
		topic := normalizeTopic(nonFlags[1:])
		showSection(topic, "Example")
		return

	case "related":
		if len(nonFlags) < 2 {
			fmt.Println("usage: goref related <topic>")
			return
		}
		topic := normalizeTopic(nonFlags[1:])
		showSection(topic, "Related")
		return

	default:
		topic := normalizeTopic(nonFlags)

		if cardExists(topic) {
			showSection(topic, section)
			return
		}

		if hasMatchingCards(topic) {
			if err := runTUI(topic); err != nil {
				fmt.Println("Error:", err)
			}
			return
		}

		fmt.Printf("ref not found: %s\n", topic)
		fmt.Println("try: goref list")
	}
}

func help() {
	fmt.Println(`GoRef - Local Go reference CLI and TUI

Usage:
  goref [flags]                        Starts TUI by default
  goref list [namespace]               List cards grouped by namespace
  goref find <query>                   Search card names and content
  goref <topic> [flags]                Show section of a card (defaults to Card)
  goref show <topic> [flags]           Show section of a card
  goref tui [query]                    Launch TUI with optional search query

Flags:
  --card                               Show ## Card section (default)
  --deep                               Show ## Deep section
  --example                            Show ## Example section
  --related                            Show ## Related section

Examples:
  goref setup.project
  goref setup project                  Normalizes to setup.project
  goref show setup.project --deep
  goref map delete key                 Normalizes to map.delete-key
  goref list setup                     Lists only setup namespace cards`)
}

func parseArgs(args []string) (nonFlags []string, section string) {
	section = "Card"
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			switch arg {
			case "--deep", "-deep":
				section = "Deep"
			case "--example", "-example", "--ex", "-ex":
				section = "Example"
			case "--related", "-related":
				section = "Related"
			case "--card", "-card":
				section = "Card"
			}
		} else {
			nonFlags = append(nonFlags, arg)
		}
	}
	return nonFlags, section
}

func listRefs(filter string) {
	names, err := getCardNames()
	if err != nil {
		fmt.Println("Error loading cards:", err)
		return
	}

	filter = strings.ToLower(strings.TrimSpace(filter))

	groups := make(map[string][]string)
	var prefixes []string

	for _, name := range names {
		parts := strings.SplitN(name, ".", 2)
		prefix := "other"
		if len(parts) > 1 {
			prefix = parts[0]
		}

		if filter != "" {
			if !strings.HasPrefix(strings.ToLower(name), filter) && !strings.HasPrefix(strings.ToLower(prefix), filter) {
				continue
			}
		}

		if _, exists := groups[prefix]; !exists {
			prefixes = append(prefixes, prefix)
		}
		groups[prefix] = append(groups[prefix], name)
	}

	sort.Strings(prefixes)

	if len(prefixes) == 0 {
		fmt.Println("No cards found.")
		return
	}

	for _, prefix := range prefixes {
		fmt.Printf("%s:\n", prefix)
		for _, name := range groups[prefix] {
			fmt.Printf("  %s\n", name)
		}
		fmt.Println()
	}
}

func showSection(topic, section string) {
	contentBytes, err := getCardContent(topic)
	if err != nil {
		fmt.Printf("ref not found: %s\n", topic)
		fmt.Println("try: goref list")
		return
	}

	content := string(contentBytes)
	body, err := extractSection(content, section)
	if err != nil {
		fmt.Printf("section not found: %s\n", section)
		return
	}

	title := getCardTitle(content)
	if title != "" {
		fmt.Println(title)
		fmt.Println()
	}

	fmt.Println(body)
}

func searchRefs(query string) {
	query = strings.ToLower(strings.TrimSpace(query))
	names, err := getCardNames()
	if err != nil {
		fmt.Println("Error loading cards:", err)
		return
	}

	var nameMatches []string
	var contentMatches []string

	for _, name := range names {
		if strings.Contains(strings.ToLower(name), query) {
			nameMatches = append(nameMatches, name)
			continue
		}

		contentBytes, err := getCardContent(name)
		if err == nil {
			if strings.Contains(strings.ToLower(string(contentBytes)), query) {
				contentMatches = append(contentMatches, name)
			}
		}
	}

	if len(nameMatches) == 0 && len(contentMatches) == 0 {
		fmt.Println("No refs found.")
		return
	}

	if len(nameMatches) > 0 {
		fmt.Println("Card name matches:")
		for _, name := range nameMatches {
			fmt.Printf("  %s\n", name)
		}
		fmt.Println()
	}

	if len(contentMatches) > 0 {
		fmt.Println("Content matches:")
		for _, name := range contentMatches {
			fmt.Printf("  %s\n", name)
		}
		fmt.Println()
	}
}
