package app

import (
	"fmt"
	"net/http"
	"os"
)

type App struct{}

func (*App) Run() {
	initApp()

	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello from app")
	})

	fmt.Println(http.ListenAndServe(":8000", server))
	os.Exit(0)
}

func initApp() {
}
