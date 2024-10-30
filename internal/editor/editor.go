package editor

import (
	"os/exec"

	"github.com/muesli/coral"
)

var defaultEditor = "vim"

var (
	name string
)

// Opens the current slide as a split window in tmux.
func OpenNewWindow(fileName string) error {
	var cmd *exec.Cmd

	switch name {
	case "vim", "nvim":
		cmd = exec.Command("tmux", "split-window", "-h", name, fileName)
	case "code":
		cmd = exec.Command(name, fileName)
	}

	return cmd.Start()
}

func InitEditorFlag(rootCmd *coral.Command) {
	rootCmd.PersistentFlags().StringVarP(&name, "editor", "e", defaultEditor, "Specify the editor to use")
}
