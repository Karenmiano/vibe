package routes

import (
	"net/http"
	"github.com/Karenmiano/vibe/internal/handlers"
)


func NewRouter() http.Handler  {
	mux := http.NewServeMux()

	mux.Handle("/", &handlers.TemplateHandler{Filename: "chat.html"})

	r := handlers.NewRoom()
	mux.Handle("/room", r)

	go r.Run()

	return mux
}
