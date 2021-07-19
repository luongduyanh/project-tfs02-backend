package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project-tfs02/api/models"
	"project-tfs02/api/utils"
	"project-tfs02/api/utils/format_error"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateOrder(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	order := models.Order{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	orderCreated, err := order.SaveOrder(server.DB)
	if err != nil {
		formattedError := format_error.FormatError(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, orderCreated.ID))
	utils.JSON(w, http.StatusCreated, orderCreated)
}

func (server *Server) GetOrders(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}

	orders, err := order.FindAllOrders(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, orders)
}

func (server *Server) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {}

func (server *Server) GetOrderLinesByOrderID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	order := models.Order{}
	ordertGotten, err := order.FindOrderLinesByOrderID(server.DB, uint(uid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, ordertGotten)
}
