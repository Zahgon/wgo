//go:build !windows

package main

import (
	"bytes"
	"io"
	"os/exec"
)

// NOTE: We shouldn't encounter the macOS file limit of 256 anymore now that
// importing "os" automatically raises it for us in an init function.
// https://github.com/golang/go/issues/46279
// https://go-review.googlesource.com/c/go/+/393354/4/src/os/rlimit.go

const (
	specialChars      = "\\'\"`${[|&;<>()*?!"
	extraSpecialChars = " \t\n"
	prefixChars       = "~"
)

// stop stops the command and all its child processes.
func stop(cmd *exec.Cmd) {
	_ = "STUB: not implemented"
	// https://stackoverflow.com/questions/22470193/why-wont-go-kill-a-child-process-correctly
	// https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773
	return
}

// https://stackoverflow.com/questions/22470193/why-wont-go-kill-a-child-process-correctly
// https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773
func setpgid(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }

// flushStdin tells the OS to flush the text currently buffered in stdin.
//
// Using raw numbers here instead of importing package constants to avoid
// creating a bunch more files to use conditional compilation to select the
// right package constants for the os and arch.
func flushStdin(r io.Reader) { _ = "STUB: not implemented"; return }

// ioctl(fd, TCFLSH, TCIFLUSH)

/* golang.org/x/sys/unix.TCIFLUSH on linux */

// ioctl(fd, TIOCFLUSH, &TCIFLUSH)
/* golang.org/x/sys/unix.TCIFLUSH */

/* syscall.SYS_IOCTL on darwin, dragonfly, freebsd, netbsd, openbsd */

/* golang.org/x/sys/unix.TIOCFLUSH on darwin, dragonfly, freebsd, netbsd, openbsd */

// linuxSYSIOCTL returns syscall.SYS_IOCTL based on runtime.GOARCH.
func linuxSYSIOCTL() uintptr { _ = "STUB: not implemented"; return 0 }

/* syscall.SYS_IOCTL on linux/amd64 */

/* syscall.SYS_IOCTL on linux/arm64, linux/loong64, linux/riscv64 */

/* syscall.SYS_IOCTL on linux/mips, linux/mipsle */

/* syscall.SYS_IOCTL on linux/mips64, linux/mips64le */

/* syscall.SYS_IOCTL on other linux architectures */

// linuxSYSIOCTL returns golang.org/x/sys/unix.TCFLSH based on runtime.GOARCH.
func linuxTCFLSH() uintptr { _ = "STUB: not implemented"; return 0 }

/* golang.org/x/sys/unix.TCFLSH on linux/mips, linux/mips64, linux/mips64le, linux/mipsle */

/* golang.org/x/sys/unix.TCFLSH on linux/ppc, linux/ppc64, linux/ppc64le */

/* golang.org/x/sys/unix.TCFLSH on linux/sparc64 */

/* golang.org/x/sys/unix.TCFLSH on other linux architectures */

// joinArgs joins the arguments of the command into a string which can then be
// passed to `exec.Command("sh", "-c", $STRING)`. Examples:
//
// ["echo", "foo"] => echo foo
//
// ["echo", "hello goodbye"] => echo 'hello goodbye'
func joinArgs(args []string) string {
	_ = "STUB: not implemented"
	// https://github.com/kballard/go-shellquote/blob/master/quote.go
	//
	// # Copyright (C) 2014 Kevin Ballard
	//
	// Permission is hereby granted, free of charge, to any person obtaining
	// a copy of this software and associated documentation files (the "Software"),
	// to deal in the Software without restriction, including without limitation
	// the rights to use, copy, modify, merge, publish, distribute, sublicense,
	// and/or sell copies of the Software, and to permit persons to whom the
	// Software is furnished to do so, subject to the following conditions:
	//
	// The above copyright notice and this permission notice shall be included
	// in all copies or substantial portions of the Software.
	//
	// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
	// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
	// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
	// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
	// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
	return ""
}

func quote(word string, buf *bytes.Buffer) {
	_ = "STUB: not implemented"
	// https://github.com/kballard/go-shellquote/blob/master/quote.go
	//
	// # Copyright (C) 2014 Kevin Ballard
	//
	// Permission is hereby granted, free of charge, to any person obtaining
	// a copy of this software and associated documentation files (the "Software"),
	// to deal in the Software without restriction, including without limitation
	// the rights to use, copy, modify, merge, publish, distribute, sublicense,
	// and/or sell copies of the Software, and to permit persons to whom the
	// Software is furnished to do so, subject to the following conditions:
	//
	// The above copyright notice and this permission notice shall be included
	// in all copies or substantial portions of the Software.
	//
	// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
	// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
	// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
	// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
	// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
	// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
	// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
	return
}

// We want to try to produce a "nice" output. As such, we will
// backslash-escape most characters, but if we encounter a space, or if we
// encounter an extra-special char (which doesn't work with
// backslash-escaping) we switch over to quoting the whole word. We do this
// with a space because it's typically easier for people to read multi-word
// arguments when quoted with a space rather than with ugly backslashes
// everywhere.

// oops, no content

// copy the non-special chars up to this point

// start over in quote mode

// quote mode
// Use single-quotes, but if we find a single-quote in the word, we need
// to terminate the string, emit an escaped quote, and start the string up
// again
