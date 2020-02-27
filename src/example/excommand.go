package main

import (
	"fmt"
	"os/exec"
	"syscall"
)

// exec.Command Example
func main() {
	c := exec.Command("cmd", "/C", "del", "del.txt")

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	binary, lookErr := exec.LookPath("cmd.exe")
	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Println(binary)

	cmd_path := binary
	cmd_instance := exec.Command(cmd_path, "/c", "notepad")
	cmd_instance.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd_instance.Output()
}
