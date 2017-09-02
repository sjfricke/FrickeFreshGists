package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func run(args ...string) {

	cmd := exec.Command("adb", args...)

	fmt.Printf("adb %s", strings.Join(args, " "))

	err := cmd.Run()
	if err != nil {
		fmt.Printf("FAILED: adb %s: %v", strings.Join(args, " "), err)
	}
}

func main() {
	fmt.Println("STARTING...")

	run("shell", "setprop", "debug.vulkan.layers", "")

	fmt.Println("") // new line format
}
