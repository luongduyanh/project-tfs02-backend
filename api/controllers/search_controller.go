package controllers

import (
	"net/http"
	"project-tfs02/api/models"
	"project-tfs02/api/utils"
	"strings"

	"github.com/gorilla/mux"
)

func (server *Server) SearchProductsByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := strings.ReplaceAll(vars["name"], "-", " ")
	product := models.Product{}
	productGotten, err := product.FindProductsByName(server.DB, name)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, productGotten)

}
