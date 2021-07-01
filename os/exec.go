package main

import (
	"os"
	"os/exec"
	"syscall"
)

// completely replace the current Go process with another (perhaps non-Go) one

func main() {
	// Go requires an absolute path to the binary we want to execute
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	// Exec also needs a set of environment variables to use
	// Here we just provide our current environmen
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
