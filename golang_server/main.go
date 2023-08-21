package main

import (
	"fmt"
	"idk/main/api"
	"idk/main/db"
	"idk/main/env"
)

func main() {
	err := db.Connect(db.UrlFromEnv(env.Env))
	if err != nil {
		fmt.Println("Db connection error")
		return
	}
	fmt.Println("Connected")

	go api.Init()

	for {
		fmt.Scanln()
	}
}
