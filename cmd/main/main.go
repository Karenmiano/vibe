package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Karenmiano/vibe/internal/database/postgres"
	"github.com/Karenmiano/vibe/internal/hub"
	"github.com/Karenmiano/vibe/internal/room"
	"github.com/Karenmiano/vibe/pkg/utilities"
)



func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	validator, trans := utilities.NewValidator()

	mux := http.NewServeMux()

	hub := hub.NewHub()
	mux.Handle("/hub", hub)
	go hub.Run()

	roomRepo := postgres.NewRoomRepository(dbpool)
	roomService := room.NewRoomService(roomRepo)
	roomHandler := room.NewRoomHandler(roomService, validator, trans)
	mux.HandleFunc("POST /rooms", roomHandler.CreateRoom)

	port := ":8080"
	fmt.Printf("Server listening on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
