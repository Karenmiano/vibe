package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/alexedwards/scs/goredisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

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

	gob.Register(uuid.UUID{})

	// Postgres db connection
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	// Redis connection
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opt)
	defer rdb.Close()


	// Initialize a new session manager and configure it to use goredisstore as the session store.
	sessionManager := scs.New()
	sessionManager.Store = goredisstore.New(rdb)

	validator, trans := utilities.NewValidator()

	authMiddleware := middleware.NewAuthMiddleware(sessionManager)

	mux := http.NewServeMux()


	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", fs))

	hub := hub.NewHub()
	mux.Handle("/", authMiddleware.Authenticate(http.HandlerFunc(hub.Hub)))
	mux.HandleFunc("/wsConnect", hub.ServeWebSocket)
	go hub.Run()

	roomRepo := postgres.NewRoomRepository(dbpool)
	roomService := room.NewRoomService(roomRepo)
	roomHandler := room.NewRoomHandler(roomService, validator, trans)
	mux.HandleFunc("POST /rooms", roomHandler.CreateRoom)

	userRepo := postgres.NewUserRepository(dbpool)
	userService := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userService, sessionManager, validator, trans)
	mux.HandleFunc("POST /register", userHandler.RegisterUser)
	mux.HandleFunc("POST /login", userHandler.LoginUser)
	mux.HandleFunc("POST /logout", userHandler.LogoutUser)

	port := ":8080"
	fmt.Printf("Server listening on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, sessionManager.LoadAndSave(mux)); err != nil {
		log.Panic("ListenAndServe: ", err)
	}
}
