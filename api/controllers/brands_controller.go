package controllers

import (
	"net/http"
	"project-tfs02/api/models"
	"project-tfs02/api/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) GetBrands(w http.ResponseWriter, r *http.Request) {

	brand := models.Brand{}

	brands, err := brand.FindAllBrands(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, brands)
}

func (server *Server) GetBrand(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	brand := models.Brand{}
	brandGotten, err := brand.FindBrandByID(server.DB, uint(uid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, brandGotten)
}

func (server *Server) GetProductsByBrandID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	subCate := models.Product{}
	SubCatesGotten, err := subCate.FindProductsByBrandID(server.DB, uint(uid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, SubCatesGotten)
}
