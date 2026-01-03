package user

import (
	"net/http"
	"regexp"
	"strings"
	"unicode/utf8"

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

	data.Username = strings.TrimSpace(data.Username)

	uLen := utf8.RuneCountInString(data.Username)
	if uLen < 3 {
		data.Errors["username"] = append(data.Errors["username"], "Username must be at least 3 characters long")
	} else if uLen > 16 {
		data.Errors["username"] = append(data.Errors["username"], "Username cannot exceed 16 characters")
	} else {
		if !rxUserName.MatchString(data.Username) {
			data.Errors["username"] = append(data.Errors["username"], "Username can only contain letters, numbers, dots (.), and underscores (_)")
		}
		if strings.HasPrefix(data.Username, ".") || strings.HasPrefix(data.Username, "_") {
			data.Errors["username"] = append(data.Errors["username"], "Username cannot start with a dot or underscore")
		}
		if strings.HasSuffix(data.Username, ".") || strings.HasSuffix(data.Username, "_") {
			data.Errors["username"] = append(data.Errors["username"], "Username cannot end with a dot or underscore")
		}
	}

	if utf8.RuneCountInString(data.Password) < 6 {
		data.Errors["password"] = append(data.Errors["password"], "Password must be at least 6 character long")
	}
	
	return len(data.Errors) == 0
}


type UserHandler struct {

}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	newUserData := &CreateUserData{
		Username: r.PostFormValue("username"),
		Password: r.PostFormValue("password"),
	}

	if newUserData.Validate() == false {
		utilities.Render(w, "web/templates/register.html", newUserData)
		return
	}

	w.Write([]byte("User registration success!"))
}

func (h *UserHandler) RegisterUserForm(w http.ResponseWriter, r *http.Request) {
	utilities.Render(w, "web/templates/register.html", nil)
}
