package tui

import (
	"github.com/charmbracelet/lipgloss"
	styles "github.com/gregriff/gpt-cli-go/tui/styles"
)

// ChatEntry contains unformatted chat text for storage (but still valid Markdown)
type ChatEntry struct {
	prompt,
	reasoning,
	response,
	error string
}

// createFormattedPrompt creates a prompt string formatted with margin and padding.
// It tries to emulate iMessage, right-justifying the prompt and adding a left margin
func (c *ChatEntry) createFormattedPrompt(marginStyle, promptStyle lipgloss.Style, maxWidth int) string {
	fullPromptStyle := promptStyle.
		PaddingLeft(maxWidth - lipgloss.Width(c.prompt) - styles.H_PADDING*2). // replace lg.W with textWidth if bugged
		PaddingTop(styles.PROMPT_V_PADDING).
		PaddingBottom(styles.PROMPT_V_PADDING)
	return lipgloss.JoinHorizontal(lipgloss.Top, marginStyle.Render(""), fullPromptStyle.Render(c.prompt))
}
