package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	var c *exec.Cmd
	if runtime.GOOS == "windows" {
		c = exec.Command("cmd", "/C", "git", "--version")
	} else {
		c = exec.Command("git", "--version")
	}
	// c.Dir = "/abs/path"
	out, err := c.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
}
