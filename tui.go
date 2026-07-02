package main

import (
	"fmt"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type tuiModel struct {
	names           []string
	filtered        []string
	cursor          int
	activeNamespace string

	mode      string
	searching bool
	query     string

	// Advanced TUI navigation
	viewingRelated bool
	relatedCards   []string
	relatedCursor  int
	history        []string // stack of previously visited cards

	width  int
	height int
	scroll int

	content string
	err     string
}

func runTUI(initialQuery string) error {
	names, err := getCardNames()
	if err != nil {
		return err
	}

	m := tuiModel{
		names:    names,
		filtered: names,
		mode:     "Card",
		width:    100,
		height:   27,
	}

	initialQuery = strings.TrimSpace(initialQuery)
	if initialQuery != "" {
		namespaces := m.getNamespaces()
		isNamespace := false
		for _, ns := range namespaces {
			if strings.EqualFold(ns, initialQuery) {
				m.activeNamespace = ns
				isNamespace = true
				break
			}
		}
		if !isNamespace {
			m.query = initialQuery
		}
	}

	m.applyFilter()
	m.refreshContent()

	_, err = tea.NewProgram(m, tea.WithAltScreen()).Run()
	return err
}

func (m tuiModel) Init() tea.Cmd {
	return nil
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height - 3
		m.refreshContent()
		return m, nil

	case tea.KeyMsg:
		key := msg.String()

		if m.searching {
			switch key {
			case "esc":
				m.searching = false
				return m, nil
			case "enter":
				m.searching = false
				return m, nil
			case "backspace", "ctrl+h":
				if len(m.query) > 0 {
					m.query = m.query[:len(m.query)-1]
					m.applyFilter()
					m.refreshContent()
				}
				return m, nil
			default:
				if len(key) == 1 {
					m.query += key
					m.applyFilter()
					m.refreshContent()
				}
				return m, nil
			}
		}

		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "/":
			m.searching = true
			m.viewingRelated = false
			return m, nil

		case "tab":
			namespaces := m.getNamespaces()
			if len(namespaces) > 0 {
				idx := -1
				for i, ns := range namespaces {
					if ns == m.activeNamespace {
						idx = i
						break
					}
				}
				if idx == len(namespaces)-1 {
					m.activeNamespace = ""
				} else {
					m.activeNamespace = namespaces[idx+1]
				}
				m.applyFilter()
				m.refreshContent()
			}
			return m, nil

		case "esc":
			if m.viewingRelated {
				m.viewingRelated = false
				return m, nil
			}
			m.query = ""
			m.activeNamespace = ""
			m.applyFilter()
			m.refreshContent()
			return m, nil

		case "backspace", "ctrl+h":
			if m.viewingRelated {
				m.viewingRelated = false
				return m, nil
			}
			if len(m.history) > 0 {
				prevCard := m.history[len(m.history)-1]
				m.history = m.history[:len(m.history)-1]
				m.navigateToCard(prevCard)
			}
			return m, nil

		case "j", "down":
			if m.viewingRelated {
				if m.relatedCursor < len(m.relatedCards)-1 {
					m.relatedCursor++
				}
			} else {
				if m.cursor < len(m.filtered)-1 {
					m.cursor++
					m.scroll = 0
					m.refreshContent()
				}
			}
			return m, nil

		case "k", "up":
			if m.viewingRelated {
				if m.relatedCursor > 0 {
					m.relatedCursor--
				}
			} else {
				if m.cursor > 0 {
					m.cursor--
					m.scroll = 0
					m.refreshContent()
				}
			}
			return m, nil

		case "enter", "r":
			if m.viewingRelated {
				if len(m.relatedCards) > 0 && m.relatedCursor < len(m.relatedCards) {
					targetCard := m.relatedCards[m.relatedCursor]
					currentCard := m.filtered[m.cursor]
					m.history = append(m.history, currentCard)
					m.viewingRelated = false
					m.navigateToCard(targetCard)
				}
			} else {
				related := extractRelatedCards(m.getCurrentCardContent())
				if len(related) > 0 {
					m.viewingRelated = true
					m.relatedCards = related
					m.relatedCursor = 0
				}
			}
			return m, nil

		case "c":
			m.mode = "Card"
			m.viewingRelated = false
			m.scroll = 0
			m.refreshContent()
			return m, nil

		case "d":
			m.mode = "Deep"
			m.viewingRelated = false
			m.scroll = 0
			m.refreshContent()
			return m, nil

		case "e":
			m.mode = "Example"
			m.viewingRelated = false
			m.scroll = 0
			m.refreshContent()
			return m, nil

		case "pgdown", "ctrl+d", "right":
			m.scroll += 8
			return m, nil

		case "pgup", "ctrl+u", "left":
			m.scroll -= 8
			if m.scroll < 0 {
				m.scroll = 0
			}
			return m, nil
		}
	}

	return m, nil
}

func (m tuiModel) View() string {
	if m.width <= 0 {
		return ""
	}

	leftWidth := 30
	if m.width < 90 {
		leftWidth = 24
	}

	rightWidth := m.width - leftWidth - 5
	if rightWidth < 30 {
		rightWidth = 30
	}

	left := m.renderList(leftWidth)
	right := m.renderContent(rightWidth)

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func (m tuiModel) renderList(width int) string {
	title := " GoRef "
	if m.query != "" {
		title = " /" + m.query + " "
	} else if m.activeNamespace != "" {
		title = " [" + m.activeNamespace + "] "
	}
	if m.searching {
		title += "█"
	}

	var lines []string
	lines = append(lines, title)
	lines = append(lines, "")

	if m.viewingRelated {
		lines[0] = " Related Cards "
		if len(m.relatedCards) == 0 {
			lines = append(lines, "no related cards")
		} else {
			L_rel := m.height - 7
			end := len(m.relatedCards)
			if end > L_rel {
				end = L_rel
			}
			for i := 0; i < end; i++ {
				name := m.relatedCards[i]
				prefix := "  "
				if i == m.relatedCursor {
					prefix = "> "
				}
				lines = append(lines, prefix+tuiTruncate(name, width-4))
			}
		}
		lines = append(lines, "")
		lines = append(lines, "Enter open")
		lines = append(lines, "esc/backspace back")
	} else {
		if len(m.filtered) == 0 {
			lines = append(lines, "no matches")
		} else {
			maxLines := m.height - 9
			if maxLines < 5 {
				maxLines = 5
			}

			start := m.cursor - 10
			if start < 0 {
				start = 0
			}

			end := start + maxLines
			if end > len(m.filtered) {
				end = len(m.filtered)
			}

			for i := start; i < end; i++ {
				prefix := "  "
				if i == m.cursor {
					prefix = "> "
				}
				lines = append(lines, prefix+tuiTruncate(m.filtered[i], width-4))
			}
		}

		targetInnerHeight := m.height - 2
		currentLinesCount := len(lines) + 5
		for currentLinesCount < targetInnerHeight {
			lines = append(lines, "")
			currentLinesCount++
		}

		lines = append(lines, "")
		lines = append(lines, "j/k move | / search | tab cycle")
		lines = append(lines, "c card | d deep | e example")
		lines = append(lines, "Enter/r related | backspace back")
		lines = append(lines, "q quit")
	}

	return lipgloss.NewStyle().
		Width(width).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Render(strings.Join(lines, "\n"))
}

func (m tuiModel) renderContent(width int) string {
	name := "no card"
	if len(m.filtered) > 0 {
		name = m.filtered[m.cursor]
	}

	header := fmt.Sprintf(" %s [%s] ", name, strings.ToLower(m.mode))

	body := m.content
	if m.err != "" {
		body = m.err
	}

	lines := strings.Split(body, "\n")

	availableHeight := m.height - 4
	if availableHeight < 5 {
		availableHeight = 5
	}

	if m.scroll > len(lines)-1 {
		m.scroll = max(0, len(lines)-1)
	}

	end := m.scroll + availableHeight
	if end > len(lines) {
		end = len(lines)
	}

	visible := strings.Join(lines[m.scroll:end], "\n")

	// Manually pad visible lines to availableHeight to ensure both panes are perfectly even at the bottom
	visibleLines := strings.Split(visible, "\n")
	for len(visibleLines) < availableHeight {
		visibleLines = append(visibleLines, "")
	}
	visible = strings.Join(visibleLines, "\n")

	return lipgloss.NewStyle().
		Width(width).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Render(header + "\n\n" + visible)
}

func (m *tuiModel) applyFilter() {
	var source []string
	if m.activeNamespace != "" {
		prefix := m.activeNamespace + "."
		for _, name := range m.names {
			if strings.HasPrefix(strings.ToLower(name), prefix) {
				source = append(source, name)
			}
		}
	} else {
		source = m.names
	}

	query := strings.ToLower(strings.TrimSpace(m.query))
	if query == "" {
		m.filtered = source
		m.cursor = 0
		return
	}

	var nameMatches []string
	var contentMatches []string

	for _, name := range source {
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

	m.filtered = append(nameMatches, contentMatches...)
	m.cursor = 0
	m.scroll = 0
}

func (m *tuiModel) refreshContent() {
	m.err = ""
	m.content = ""

	if len(m.filtered) == 0 {
		m.content = "No cards found."
		return
	}

	name := m.filtered[m.cursor]

	body, err := getCardContent(name)
	if err != nil {
		m.err = err.Error()
		return
	}

	section, err := extractSection(string(body), m.mode)
	if err != nil {
		m.err = err.Error()
		return
	}

	leftWidth := 30
	if m.width < 90 {
		leftWidth = 24
	}
	rightWidth := m.width - leftWidth - 5
	if rightWidth < 30 {
		rightWidth = 30
	}

	rendered, err := tuiRenderMarkdown(cleanMarkdownForTUI(section), rightWidth-4)
	if err != nil {
		m.err = err.Error()
		return
	}

	m.content = strings.TrimSpace(rendered)
}

func (m tuiModel) getNamespaces() []string {
	set := make(map[string]bool)
	for _, name := range m.names {
		parts := strings.SplitN(name, ".", 2)
		if len(parts) > 1 {
			set[parts[0]] = true
		}
	}

	var list []string
	for ns := range set {
		list = append(list, ns)
	}
	sort.Strings(list)
	return list
}

func (m *tuiModel) navigateToCard(name string) {
	foundIdx := -1
	for i, fName := range m.filtered {
		if fName == name {
			foundIdx = i
			break
		}
	}

	if foundIdx != -1 {
		m.cursor = foundIdx
	} else {
		m.query = ""
		m.activeNamespace = ""
		m.applyFilter()

		for i, fName := range m.filtered {
			if fName == name {
				m.cursor = i
				break
			}
		}
	}
	m.scroll = 0
	m.refreshContent()
}

func (m tuiModel) getCurrentCardContent() string {
	if len(m.filtered) == 0 {
		return ""
	}
	name := m.filtered[m.cursor]
	bytes, err := getCardContent(name)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func extractRelatedCards(content string) []string {
	section, err := extractSection(content, "Related")
	if err != nil {
		return nil
	}
	var list []string
	lines := strings.Split(section, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "-") {
			card := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
			if card != "" {
				list = append(list, card)
			}
		} else if strings.HasPrefix(trimmed, "*") {
			card := strings.TrimSpace(strings.TrimPrefix(trimmed, "*"))
			if card != "" {
				list = append(list, card)
			}
		}
	}
	return list
}

func tuiRenderMarkdown(md string, width int) (string, error) {
	if width < 40 {
		width = 40
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return "", err
	}

	return renderer.Render(md)
}

func tuiTruncate(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}

	if len(s) <= maxLen {
		return s
	}

	if maxLen <= 1 {
		return s[:maxLen]
	}

	return s[:maxLen-1] + "…"
}

func cleanMarkdownForTUI(md string) string {
	lines := strings.Split(md, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "#") {
			level := 0
			for level < len(trimmed) && trimmed[level] == '#' {
				level++
			}
			if level > 0 && level < len(trimmed) && trimmed[level] == ' ' {
				title := strings.TrimSpace(trimmed[level:])
				if title != "" {
					lines[i] = "**" + title + "**"
				}
			}
		}
	}
	return strings.Join(lines, "\n")
}
