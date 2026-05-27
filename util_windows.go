//go:build windows

package main

import (
	"io"
	"os/exec"
)

// stop stops the command and all its child processes.
func stop(cmd *exec.Cmd) {
	_ = "STUB: not implemented"
	// https://stackoverflow.com/a/44551450
	return
}

// setpgid is a no-op on windows.
func setpgid(cmd *exec.Cmd) {
	_ = "STUB: not implemented"

	// flushStdin tells the OS to flush the text currently buffered in stdin.
	return
}

func flushStdin(r io.Reader) { _ = "STUB: not implemented"; return }

// FlushConsoleInputBuffer(handle)

// joinArgs joins the arguments of the command into a string which can then be
// passed to `exec.Command("pwsh.exe", "-command", $STRING)`. Examples:
//
// ["echo", "foo"] => echo foo
//
// ["echo", "hello goodbye"] => echo 'hello goodbye'
func joinArgs(args []string) string {
	_ = "STUB: not implemented"
	// references:
	// https://www.rlmueller.net/PowerShellEscape.htm
	// https://stackoverflow.com/a/11231504
	return ""
}
