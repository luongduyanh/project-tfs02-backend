package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"project-tfs02/api/auth"
	"project-tfs02/api/models"
	"project-tfs02/api/utils"
	"project-tfs02/api/utils/format_error"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := format_error.FormatError(err.Error())
		utils.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	utils.JSON(w, http.StatusOK, token)
}


func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	//tao ra 1 token co chua user_id
	return auth.CreateToken(uint32(user.ID))
}

func LogOut() {

}
