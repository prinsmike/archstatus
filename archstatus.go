package main

import (
"bytes"
"fmt"
"os/exec"
"github.com/mgutz/ansi"
"strings"
)

func main() {

	services := []string{
		"cronie.service",
		"httpd.service",
		"mysqld.service",
		"ntpd.service",
		"postfix.service",
		"sshd.service",
		"home.mount",
		"mnt-extra.mount",
		"tmp.mount",
		"var-lib-mysqltmp.mount",
		"var.mount",
	}

	for _, service := range services {

	// The systemctl command.
		syscommand := exec.Command("systemctl", "status", service)

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

		s := fmt.Sprintf("%s", &b)

		if strings.Contains(s, "active (running)") {
			color := ansi.ColorCode("green+h:black")
			reset := ansi.ColorCode("reset")
			fmt.Printf("%s%s%s", color, s, reset)
		} else {
			color := ansi.ColorCode("red+h:black")
			reset := ansi.ColorCode("reset")
			fmt.Printf("%s%s%s", color, s, reset)
		}
	}
}