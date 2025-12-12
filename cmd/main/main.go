package main


import (
	"fmt"
	"log"
	"net/http"

	"github.com/Karenmiano/vibe/internal/routes"
)

func main() {
	router := routes.NewRouter()
	port := ":8080"
	fmt.Printf("Server listening on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


