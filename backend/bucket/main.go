package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	Api_lib "github.com/Visoff/idk_some_name/golang_library/api/lib"
)

func main() {
	mux := http.NewServeMux()
	Api_lib.Rest(mux, "/").Get(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		file, err := os.ReadFile(path.Join("./static/", r.URL.Path))
		if err != nil {
			res.Status(404).Send("File not found")
			return
		}
		res.Status(200).Send(file)
	}).Post(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		url_path := strings.Split(r.URL.Path, "/")
		err := os.MkdirAll(path.Join("./static/", strings.Join(url_path[:len(url_path)-1], "/")), os.ModePerm)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			data = []byte("")
		}
		err = os.WriteFile(path.Join("./static/", url_path[len(url_path)-1]), data, os.ModePerm)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send(data)
	}).Delete(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		err := os.Remove(path.Join("./static/", r.URL.Path))
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send("ok")
	}).Apply()
	http.ListenAndServe(":8080", mux)
}
