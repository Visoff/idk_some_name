package main

import (
	"app/lib/env"
	"backend/auth/api"
	"fmt"
	"net/http"
)

func main() {
	PORT := env.Env("PORT", "8080")

	mux := http.NewServeMux()

	api.Init(mux)

	go fmt.Printf("Server is listening on port %v\n", PORT)
	err := http.ListenAndServe(":"+fmt.Sprint(PORT), mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
