package api

import (
	"encoding/json"
	"fmt"
	"idk/main/db"
	"io"
	"net/http"
)

func ApplyConference() {
	Mux.HandleFunc("/api/conference/call", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Provide body"))
				return
			}
			call, err := db.Query(`insert into "Call"("offer") values('%v') returning "id"`, string(body))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprint(call[0]["id"])))
			return
		case "GET":
			call := r.Header.Get("Call")
			if call == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Call header"))
				return
			}
			offer, err := db.Query(`select "offer" from "Call" where "id" = '%v'`, call)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprint(offer[0]["offer"])))
			return
		}
	})

	Mux.HandleFunc("/api/conference/answer", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "POST":
			call := r.Header.Get("Call")
			if call == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Call header"))
				return
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Provide body = answer description"))
				return
			}
			_, err = db.Query(`update "Call" set "answer" = '%v'`, string(body))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte("ok"))
			return
		case "GET":
			call := r.Header.Get("Call")
			if call == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Call header"))
				return
			}
			answer, err := db.Query(`select "answer" from "Call" where "id" = '%v' and "answer" is not null`, call)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			if len(answer) == 0 {
				w.WriteHeader(200)
				w.Write([]byte("null"))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte(fmt.Sprint(answer[0]["answer"])))
			return
		}
	})

	Mux.HandleFunc("/api/conference/ice", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "POST":
			call := r.Header.Get("Call")
			if call == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Call header"))
				return
			}
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Provide body = answer description"))
				return
			}
			_, err = db.Query(`insert into "IceCandidates"("Call_id", "Candidate") values('%v', '%v')`, call, string(body))
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write([]byte("ok"))
			return
		case "GET":
			call := r.Header.Get("Call")
			if call == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Call header"))
				return
			}
			candidates, err := db.Query(`select "Candidate" from "IceCandidates" where "Call_id"= '%v'`, call)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			candidates_json, err := json.Marshal(candidates)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write(candidates_json)
			return
		}
	})

	Mux.HandleFunc("/api/conference/icepoll", func(w http.ResponseWriter, r *http.Request) {
		AllowCors(&w)
		switch r.Method {
		case "GET":
			call := r.Header.Get("Call")
			if call == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide Call header"))
				return
			}
			date := r.Header.Get("LongPoll")
			if date == "" {
				w.WriteHeader(400)
				w.Write([]byte("Provide LongPoll header"))
				return
			}
			ice, err := db.Query(`select "Candidate" from "IceCandidates" where "Call_id" = '%v' and "added_on" > '%v'`, call, date)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			ice_json, err := json.Marshal(ice)
			if err != nil {
				w.WriteHeader(500)
				w.Write([]byte(err.Error()))
				return
			}
			w.WriteHeader(200)
			w.Write(ice_json)
			return
		}
	})
}
