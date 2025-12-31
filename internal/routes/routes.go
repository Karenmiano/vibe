package routes
/*
import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/Karenmiano/vibe/internal/handlers"
	"github.com/Karenmiano/vibe/internal/repository/postgres"
)


func NewRouter(db *pgxpool.Pool) http.Handler  {
	mux := http.NewServeMux()

	// mux.Handle("/", &handlers.TemplateHandler{Filename: "chat.html"})

	hub := handlers.NewHub()
	mux.Handle("/hub", hub)
	go hub.Run()

	roomRepo := postgres.NewRoomRepository(db)
	mux.HandleFunc("POST /rooms", handlers.CreateRoom(roomRepo))

	return mux
}
*/
