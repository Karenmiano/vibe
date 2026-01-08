package user

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"unicode/utf8"

	"github.com/gorilla/sessions"

	"github.com/Karenmiano/vibe/pkg/utilities"
)


var rxUserName = regexp.MustCompile(`^[a-zA-Z0-9._]+$`)

type CreateUserData struct {
	Username string
	Password string
	Errors map[string][]string
}

func (data *CreateUserData) Validate() bool {
	data.Errors = make(map[string][]string)

	if !rxUserName.MatchString(data.Username) {
		data.Errors["username"] = append(data.Errors["username"], "Username can only contain letters, numbers, dots (.), and underscores (_)")
	} else {
		uLen := utf8.RuneCountInString(data.Username)
		if uLen < 3 {
			data.Errors["username"] = append(data.Errors["username"], "Username must be at least 3 characters long")
		} else if uLen > 16 {
			data.Errors["username"] = append(data.Errors["username"], "Username cannot exceed 16 characters")
		}
	}

	pLen := utf8.RuneCountInString(data.Password)
	if pLen < 6 {
		data.Errors["password"] = append(data.Errors["password"], "Password must be at least 6 characters long")
	} else if pLen > 72 {
		data.Errors["password"] = append(data.Errors["password"], "Password is too long")
	}
	
	return len(data.Errors) == 0
}


type UserHandler struct {
	userService *UserService
	sessionStore sessions.Store
}

func NewUserHandler(userService *UserService, sessionStore sessions.Store) *UserHandler {
	return &UserHandler{
		userService: userService,
		sessionStore: sessionStore,
	}
}

func (h *UserHandler) RegistrationForm(w http.ResponseWriter, r *http.Request) {
	utilities.Render(w, "web/templates/register.html", nil, http.StatusOK)
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	newUserData := &CreateUserData{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}

	if newUserData.Validate() == false {
		utilities.Render(w, "web/templates/register.html", newUserData, http.StatusBadRequest)
		return
	}

	userId, err := h.userService.RegisterUser(r.Context(), newUserData.Username, newUserData.Password)
	if err != nil {
		// if username is taken return error message on field username
		if errors.Is(err, ErrUserExists) {
			newUserData.Errors["username"] = append(newUserData.Errors["username"], ErrUserExists.Error())
			utilities.Render(w, "web/templates/register.html", newUserData, http.StatusBadRequest)
			return
		}

		log.Println(err)	
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	// logging user in by giving a session cookie
	session, _ := h.sessionStore.Get(r, "vibe")
	session.Values["userId"] = userId
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registration success!"))
}


func (h *UserHandler) LoginForm(w http.ResponseWriter, r *http.Request) {
	utilities.Render(w, "web/templates/login.html", nil, http.StatusOK)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	userId, err := h.userService.LoginUser(r.Context(), r.PostFormValue("username"), r.PostFormValue("password"))
	if err != nil {
		if errors.Is(err, ErrInvalidCredentials) {
			utilities.Render(w, "web/templates/login.html", ErrInvalidCredentials.Error(), http.StatusBadRequest)
			return
		}

		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}


	session, _ := h.sessionStore.Get(r, "vibe")
	session.Values["userId"] = userId
	err = session.Save(r, w)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
