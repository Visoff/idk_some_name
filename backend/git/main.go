package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	githttp "github.com/AaronO/go-git-http"
)

var root_path string = "/srv/git"

func initNewRepo(repo_path string) error {
	repo_path = path.Join(root_path, repo_path)
	err := os.MkdirAll(repo_path, os.ModePerm)
	if err != nil {
		return err
	}
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "git", "init", "--bare", "--shared")
	} else {
		cmd = exec.Command("git", "init", "--bare", "--shared")
	}
	cmd.Dir = repo_path
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func lsRepo(repo_path string) ([]byte, error) {
	repo_path = path.Join(root_path, repo_path)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "git", "ls-tree", "--name-only", "-r", "HEAD")
	} else {
		cmd = exec.Command("git", "ls-tree", "--name-only", "-r", "HEAD")
	}
	cmd.Dir = repo_path
	list, err := cmd.Output()
	if err != nil {
		return []byte{}, err
	}
	return list, nil
}

func main() {
	git := githttp.New(root_path)
	git, err := git.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mux := http.NewServeMux()

	mux.Handle("/repo/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, "/repo/")
		git.ServeHTTP(w, r)
	}))

	mux.HandleFunc("/init/", func(w http.ResponseWriter, r *http.Request) {
		err := initNewRepo(strings.TrimPrefix(r.URL.Path, "/init/"))
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
	})

	mux.HandleFunc("/ls/", func(w http.ResponseWriter, r *http.Request) {
		out, err := lsRepo(strings.TrimPrefix(r.URL.Path, "/ls/"))
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(out)
	})

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
