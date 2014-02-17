package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	// The systemctl command.
	syscommand := exec.Command("systemctl", "status", "cronie.service")

	// The grep command.
	grepcommand := exec.Command("grep", "Active:")

	// Pipe the stdout of syscommand to the stdin of grepcommand.
	grepcommand.Stdin, _ = syscommand.StdoutPipe()

	// Create a buffer of bytes.
	var b bytes.Buffer

	// Assign the address of our buffer to grepcommand.Stdout.
	grepcommand.Stdout = &b

	// Start grepcommand.
	_ = grepcommand.Start()

	// Run syscommand
	_ = syscommand.Run()

	// Wait for grepcommand to exit.
	_ = grepcommand.Wait()

	fmt.Printf("%s", &b)
}