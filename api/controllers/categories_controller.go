package controllers

import (
	"net/http"
	"project-tfs02/api/models"
	"project-tfs02/api/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) GetCategories(w http.ResponseWriter, r *http.Request) {

	category := models.Category{}

	categories, err := category.FindAllCategories(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, categories)
}

func (server *Server) GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	category := models.Category{}
	cateGotten, err := category.FindCategoryByID(server.DB, uint(uid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, cateGotten)
}

func (server *Server) GetSubCategories(w http.ResponseWriter, r *http.Request) {

	subCategory := models.SubCategory{}

	subCates, err := subCategory.FindAllSubCates(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, subCates)
}

func (server *Server) GetSubCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	subCate := models.SubCategory{}
	subCateGotten, err := subCate.FindSubCateByID(server.DB, uint(uid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, subCateGotten)
}

func (server *Server) GetSubCatesByCateID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	subCate := models.SubCategory{}
	SubCatesGotten, err := subCate.FindSubCatesByCateID(server.DB, uint(uid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, SubCatesGotten)
}
