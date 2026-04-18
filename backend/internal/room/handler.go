package room

import (
	"errors"
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

		err := utilities.DecodeJSONBody(w, r, &newRoomData)
		if err != nil {
			var mr *utilities.MalformedRequest
			if errors.As(err, &mr) {
				utilities.WriteJSON(w, mr.Status, map[string]string{"message": mr.Msg})
				return
			}
			// if error is not a MalformedRequest, log and send a 500 internal server error.
			utilities.ServerErrorJSON(w, err)
			return
		}

		err = h.validator.Struct(newRoomData)
		if err != nil {
			var validateErrors validator.ValidationErrors
			if errors.As(err, &validateErrors) {
				utilities.WriteJSON(w, http.StatusUnprocessableEntity, utilities.TransformErrors(validateErrors, h.trans))
				return
			}

			utilities.ServerErrorJSON(w, err)
			return
		}

		err = h.roomService.CreateRoom(r.Context(), newRoomData)
		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
				utilities.WriteJSON(w, http.StatusUnprocessableEntity, map[string]string{"name": "This name is already taken"})
				return
			}

			utilities.ServerErrorJSON(w, err)
			return
		}

		utilities.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Room created successfully"})
}

