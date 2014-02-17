package main

import (
	"fmt"
	"os/exec"
	"log"
)

func main() {
	out, e := exec.Command("date").Output()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("The date is %s\n", out)
}