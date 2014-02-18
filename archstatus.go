package main

import (
"bytes"
"fmt"
"os/exec"
"github.com/mgutz/ansi"
"github.com/stevedomin/termtable"
"strings"
)

func fmtString(color, str, reset string) string {
	return fmt.Sprintf("%s%s%s", color, str, reset)
}

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

	t := termtable.NewTable(nil, nil)
	t.SetHeader([]string{"SERVICE", "STATUS"})

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

			t.AddRow([]string{fmtString(color, service, reset), fmtString(color, s, reset)})		
		} else {
			color := ansi.ColorCode("red+h:black")
			reset := ansi.ColorCode("reset")

			t.AddRow([]string{fmtString(color, service, reset), fmtString(color, s, reset)})
		}
	}
	fmt.Println(t.Render())
}