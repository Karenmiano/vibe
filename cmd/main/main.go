package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Karenmiano/vibe/internal/database/postgres"
	"github.com/Karenmiano/vibe/internal/hub"
	"github.com/Karenmiano/vibe/internal/middleware"
	"github.com/Karenmiano/vibe/internal/room"
	"github.com/Karenmiano/vibe/internal/user"
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

	gob.Register(uuid.UUID{})

	sessionStore := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.SameSite = http.SameSiteLaxMode

	validator, trans := utilities.NewValidator()

	authMiddleware := middleware.NewAuthMiddleware(sessionStore)

	mux := http.NewServeMux()

	mux.Handle("/", authMiddleware.Authenticate(http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This will be the main page with rooms and chats"))
	})))

	hub := hub.NewHub()
	mux.Handle("/hub", hub)
	go hub.Run()

	roomRepo := postgres.NewRoomRepository(dbpool)
	roomService := room.NewRoomService(roomRepo)
	roomHandler := room.NewRoomHandler(roomService, validator, trans)
	mux.HandleFunc("POST /rooms", roomHandler.CreateRoom)

	userRepo := postgres.NewUserRepository(dbpool)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService, sessionStore)
	mux.HandleFunc("GET /register", userHandler.RegistrationForm)
	mux.HandleFunc("POST /register", userHandler.RegisterUser)
	mux.HandleFunc("GET /login", userHandler.LoginForm)
	mux.HandleFunc("POST /login", userHandler.LoginUser)

	port := ":8080"
	fmt.Printf("Server listening on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
