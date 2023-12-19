package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type Alias struct {
	WorkingDir string
	Cmd        string
}

func main() {
	aliases := map[string]Alias{
		"format": {
			WorkingDir: "~/dev/crypto/bsenceng/",
			Cmd:        "find bsenceng/ -iname \"*.h\" -o -iname \"*.cpp\" | xargs clang-format -i --style=\"file:../lintrc/.clang-format-03\"",
		},
	}

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	alias, found := aliases[os.Args[1]]
	if !found {
		fmt.Fprintf(os.Stderr, "%s not found\n", os.Args[1])
		os.Exit(1)
	}

	cmd := exec.Command("sh", alias.Cmd)
	cmd.Dir = alias.WorkingDir

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	fmt.Fprint(os.Stdout, stderr.String())
	fmt.Fprint(os.Stderr, stderr.String())

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		os.Exit(exitErr.ExitCode())
	}
}
