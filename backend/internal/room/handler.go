package room

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Karenmiano/vibe/internal/models"
	"github.com/Karenmiano/vibe/pkg/utilities"
)

type RoomHandler struct {
	roomService *RoomService
	validator *validator.Validate
	trans ut.Translator
}

func NewRoomHandler(roomService *RoomService, validator *validator.Validate, trans ut.Translator) *RoomHandler {
	return &RoomHandler{
		roomService:  roomService,
		validator: validator,
		trans:    trans,
	}
}

func (h *RoomHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
		var newRoomData models.CreateRoomData

		err := json.NewDecoder(r.Body).Decode(&newRoomData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = h.validator.Struct(newRoomData)
		if err != nil {
			var validateErrors validator.ValidationErrors
			if errors.As(err, &validateErrors) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(utilities.TransformErrors(validateErrors, h.trans))
				return
			} else {
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}
		}

		err = h.roomService.CreateRoom(r.Context(), newRoomData)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"name": "This name is already taken"})
				return
			}

			log.Println(err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("New room created successfully"))
}

