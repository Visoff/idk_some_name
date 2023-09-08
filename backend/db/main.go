package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	Api_lib "github.com/Visoff/idk_some_name/golang_library/api/lib"
	api_middleware "github.com/Visoff/idk_some_name/golang_library/api/middleware"
	"github.com/Visoff/idk_some_name/golang_library/db"
	"github.com/Visoff/idk_some_name/golang_library/env"
)

func main() {
	err := db.Connect(db.UrlFromEnv(env.Env))
	if err != nil {
		fmt.Println("Db connection error")
		fmt.Println(err.Error())
		return
	}

	mux := http.NewServeMux()

	Api_lib.Rest(mux, "/").Use(api_middleware.AllowCors).Use(api_middleware.BodyParser).Get(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		path := strings.Split(r.URL.Path, "/")
		query := r.URL.Query().Get("query")
		if query != "" {
			query = "where " + query
		}
		data, err := db.Query(`select * from "%v"%v`, path[len(path)-1], query)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send(data)
	}).Post(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		path := strings.Split(r.URL.Path, "/")
		var body map[string]string
		err := json.Unmarshal([]byte(r.URL.Query().Get("Body")), &body)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		var keys []string
		var values []string
		for key, value := range body {
			keys = append(keys, `"`+key+`"`)
			values = append(values, `'`+value+`'`)
		}
		inserted, err := db.Query(`insert into "%v"(%v) values(%v) returning *`, path[len(path)-1], strings.Join(keys, ", "), strings.Join(values, ", "))
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send(inserted)
	}).Patch(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		path := strings.Split(r.URL.Path, "/")
		query := r.URL.Query().Get("query")
		if query != "" {
			query = "where " + query
		}
		var body map[string]string
		err := json.Unmarshal([]byte(r.URL.Query().Get("Body")), &body)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		var values []string
		for key, value := range body {
			values = append(values, `"`+key+`" = '`+value+`'`)
		}
		updated, err := db.Query(`update "%v" set %v%v returning *`, path[len(path)-1], strings.Join(values, ", "), query)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send(updated)
	}).Delete(func(w http.ResponseWriter, r *http.Request) {
		res := Api_lib.NewBetterResponseWriter(w)
		path := strings.Split(r.URL.Path, "/")
		query := r.URL.Query().Get("query")
		if query != "" {
			query = "where " + query
		} else {
			res.Status(200).Send(`Please provide selector in "query" query parameter`)
			return
		}
		_, err := db.Query(`delete from "%v"%v`, path[len(path)-1], query)
		if err != nil {
			res.Status(500).Send(err.Error())
			return
		}
		res.Status(200).Send("ok")
	}).Apply()

	go fmt.Println("Running server on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
