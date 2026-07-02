package main

import (
	"reflect"
	"testing"
)

func TestParseHeading(t *testing.T) {
	tests := []struct {
		line  string
		level int
		title string
		ok    bool
	}{
		{"# setup.project", 1, "setup.project", true},
		{"## Card", 2, "Card", true},
		{"### Subcard: Simple", 3, "Subcard: Simple", true},
		{"#", 0, "", false},
		{"####### Too Deep", 0, "", false},
		{"No Heading", 0, "", false},
		{"##NoSpace", 0, "", false},
	}

	for _, tt := range tests {
		level, title, ok := parseHeading(tt.line)
		if level != tt.level || title != tt.title || ok != tt.ok {
			t.Errorf("parseHeading(%q) = (%d, %q, %t); want (%d, %q, %t)",
				tt.line, level, title, ok, tt.level, tt.title, tt.ok)
		}
	}
}

func TestExtractSection(t *testing.T) {
	content := `# setup.project

## Card

Set up a new Go project from zero.

### Subcard: Create Folder

mkdir myapp

## Example

Some working example.

## Deep

Longer explanation.`

	tests := []struct {
		section string
		want    string
		wantErr bool
	}{
		{"Card", "Set up a new Go project from zero.\n\n### Subcard: Create Folder\n\nmkdir myapp", false},
		{"Example", "Some working example.", false},
		{"Deep", "Longer explanation.", false},
		{"Gotchas", "", true},
	}

	for _, tt := range tests {
		got, err := extractSection(content, tt.section)
		if (err != nil) != tt.wantErr {
			t.Errorf("extractSection(%q) error = %v, wantErr %v", tt.section, err, tt.wantErr)
		}
		if got != tt.want {
			t.Errorf("extractSection(%q) = %q; want %q", tt.section, got, tt.want)
		}
	}
}

func TestGetCardTitle(t *testing.T) {
	content := `Some preamble
# map.delete-key

## Card
delete(m, key)`

	title := getCardTitle(content)
	want := "# map.delete-key"
	if title != want {
		t.Errorf("getCardTitle() = %q; want %q", title, want)
	}
}

func TestNormalizeTopic(t *testing.T) {
	// Seeded/existing card names mock is setup in code if possible.
	// Since normalizeTopic dynamically reads using getCardNames (which reads files in refs/ or embed),
	// it will see the real refs we seeded earlier. Let's test with those!
	tests := []struct {
		parts []string
		want  string
	}{
		{[]string{"setup.project"}, "setup.project"},
		{[]string{"setup", "project"}, "setup.project"},
		{[]string{"map", "delete", "key"}, "map.delete-key"},
		{[]string{"map", "delete-key"}, "map.delete-key"},
		{[]string{"slice", "remove", "item"}, "slice.remove-item"},
		{[]string{"cmd", "tidy"}, "cmd.tidy"},
		{[]string{"unknown", "namespace"}, "unknown.namespace"}, // Fallback behavior
	}

	for _, tt := range tests {
		got := normalizeTopic(tt.parts)
		if got != tt.want {
			t.Errorf("normalizeTopic(%v) = %q; want %q", tt.parts, got, tt.want)
		}
	}
}

func TestExtractRelatedCards(t *testing.T) {
	content := `# test.card

## Related

- map.delete-key
- cmd.tidy
* setup.project`

	got := extractRelatedCards(content)
	want := []string{"map.delete-key", "cmd.tidy", "setup.project"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("extractRelatedCards() = %v; want %v", got, want)
	}
}
