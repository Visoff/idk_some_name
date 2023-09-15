package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	githttp "github.com/AaronO/go-git-http"
)

func main() {
	root_path := "/srv/git"
	git := githttp.New(root_path)
	git, err := git.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "git", "init", "--bare", "--shared")
	} else {
		cmd = exec.Command("git", "init", "--bare", "--shared")
	}
	cmd.Dir = root_path
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = http.ListenAndServe(":8080", git)
	if err != nil {
		fmt.Println(err.Error())
	}
}
