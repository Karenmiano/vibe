package user

import (
	"errors"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"

	"github.com/Karenmiano/vibe/pkg/utilities"
)

type createUserData struct {
	Username string `json:"username" validate:"required,username,min=3,max=16"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

type loginUserData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserHandler struct {
	userService *UserService
	sessionStore sessions.Store
	validator *validator.Validate
	trans ut.Translator
}

func NewUserHandler(userService *UserService, sessionStore sessions.Store, validator *validator.Validate, trans ut.Translator) *UserHandler {
	return &UserHandler{
		userService: userService,
		sessionStore: sessionStore,
		validator: validator,
		trans: trans,
	}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUserData createUserData

	err := utilities.DecodeJSONBody(w, r, &newUserData)
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

	err = h.validator.Struct(newUserData)
	if err != nil {
		var validateErrors validator.ValidationErrors
		if errors.As(err, &validateErrors) {
			utilities.WriteJSON(w, http.StatusUnprocessableEntity, utilities.TransformErrors(validateErrors, h.trans))
			return
		}

		utilities.ServerErrorJSON(w, err)
		return
	}

	err = h.userService.RegisterUser(r.Context(), newUserData.Username, newUserData.Password)
	if err != nil {
		// if username is taken return error message on field username
		if errors.Is(err, ErrUserExists) {
			utilities.WriteJSON(w, http.StatusUnprocessableEntity, map[string]string{"username": ErrUserExists.Error()})
			return
		}

		utilities.ServerErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user created successfully"})
}


func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var creds loginUserData

	err := utilities.DecodeJSONBody(w, r, &creds)
	if err != nil {
		var mr *utilities.MalformedRequest
		if errors.As(err, &mr) {
			utilities.WriteJSON(w, mr.Status, map[string]string{"message": mr.Msg})
			return
		}

		utilities.ServerErrorJSON(w, err)
		return
	}

	err = h.validator.Struct(creds)
	if err != nil {
		var validateErrors validator.ValidationErrors
		if errors.As(err, &validateErrors) {
			utilities.WriteJSON(w, http.StatusUnprocessableEntity, utilities.TransformErrors(validateErrors, h.trans))
			return
		}

		utilities.ServerErrorJSON(w, err)
		return
	}

	userId, err := h.userService.LoginUser(r.Context(), creds.Username, creds.Password)
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			utilities.WriteJSON(w, http.StatusBadRequest, map[string]string{"message": ErrInvalidCredentials.Error()})
			return
		}

		utilities.ServerErrorJSON(w, err)
		return
	}


	session, _ := h.sessionStore.Get(r, "vibe")
	session.Values["userId"] = userId
	err = session.Save(r, w)
	if err != nil {
		utilities.ServerErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, map[string]string{"message": "login successful"})
}

func (h *UserHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	session, _ := h.sessionStore.Get(r, "vibe")
	delete(session.Values, "userId")
	err := session.Save(r, w)
	if err != nil {
		utilities.ServerErrorJSON(w, err)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
