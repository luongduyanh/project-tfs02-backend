package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project-tfs02/api/models"
	"project-tfs02/api/utils"
	"project-tfs02/api/utils/format_error"
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

func (server *Server) GetOrders(w http.ResponseWriter, r *http.Request) {}

func (server *Server) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {}

func (server *Server) GetOrderLinesByOrderID(w http.ResponseWriter, r *http.Request) {}
