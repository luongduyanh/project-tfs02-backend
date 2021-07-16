package controllers

import (
	"net/http"
	"project-tfs02/api/utils"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
