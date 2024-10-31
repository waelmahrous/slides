package editor

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/coral"
)

var (
	program = "vim"
)

// Opens the current slide as a split window in tmux.
func OpenNewWindow(fileName string) tea.Cmd {
	return tea.ExecProcess(exec.Command(program, fileName), nil)
}

func InitEditorFlag(rootCmd *coral.Command) {
	rootCmd.PersistentFlags().StringVarP(&program, "editor", "e", program, "Specify the editor to use")
}
