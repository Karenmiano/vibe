package user

import (
	"errors"
	"net/http"

	"github.com/alexedwards/scs/v2"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/Karenmiano/vibe/pkg/utilities"
)

type createUserData struct {
	FullName string `json:"fullName" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username,min=3,max=16"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}

type loginUserData struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserHandler struct {
	userService *UserService
	sessionManager *scs.SessionManager
	validator *validator.Validate
	trans ut.Translator
}

func NewUserHandler(userService *UserService, sessionManager *scs.SessionManager, validator *validator.Validate, trans ut.Translator) *UserHandler {
	return &UserHandler{
		userService: userService,
		sessionManager: sessionManager,
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

	err = h.userService.RegisterUser(r.Context(), newUserData.FullName, newUserData.Email, newUserData.Username, newUserData.Password)
	if err != nil {
		if errors.Is(err, ErrUsernameTaken) { // if username is taken return error message on field username
			utilities.WriteJSON(w, http.StatusConflict, map[string]string{"username": ErrUsernameTaken.Error()})
			return
		} else if errors.Is(err, ErrEmailTaken) { // if email is taken return error message on field email
			utilities.WriteJSON(w, http.StatusConflict, map[string]string{"email": ErrEmailTaken.Error()})
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
			utilities.WriteJSON(w, http.StatusUnauthorized, map[string]string{"message": ErrInvalidCredentials.Error()})
			return
		}

		utilities.ServerErrorJSON(w, err)
		return
	}

	// create a session and add userId to it
	// renew the sesion token first to prevent session fixation attacks
	err = h.sessionManager.RenewToken(r.Context())
	if err != nil {
		utilities.ServerErrorJSON(w, err)
		return
	}
	h.sessionManager.Put(r.Context(), "userId", userId)

	utilities.WriteJSON(w, http.StatusOK, map[string]string{"message": "login successful"})
}

func (h *UserHandler) LogoutUser(w http.ResponseWriter, r *http.Request) {
	err := h.sessionManager.Destroy(r.Context())
	if err != nil {
		utilities.ServerErrorJSON(w, err)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}
