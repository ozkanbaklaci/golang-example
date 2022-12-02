package main

import (
	"fmt"
	"os"
   "net/http"

	"github.com/go-martini/martini"
)

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello Ziraat Team from Özkan"
	}

	m := martini.Classic()
       
	m.Get("/", func() string {
		fmt.Println(message)
		return message
	})
	 if err := http.ListenAndServe(":11130", m); err != nil {
		fmt.Println(err)
	}
}
