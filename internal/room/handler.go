package room

import (
	"encoding/json"
	"errors"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/Karenmiano/vibe/internal/models"
	"github.com/Karenmiano/vibe/pkg/utilities"
)

type RoomHandler struct {
	service *RoomService
	validator *validator.Validate
	trans ut.Translator
}

func NewRoomHandler(service *RoomService, validator *validator.Validate, trans ut.Translator) *RoomHandler {
	return &RoomHandler{
		service:  service,
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
				http.Error(w, "An unexpected error occurred", http.StatusBadRequest)
				return
			}
		}

		err = h.service.CreateRoom(r.Context(), newRoomData)
		if err != nil {
			http.Error(w, "Could not create room", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("New room created successfully"))
	}

