package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"project-tfs02/api/auth"
	"project-tfs02/api/models"
	"project-tfs02/api/rabbitMQ/producer"
	"project-tfs02/api/rabbitMQ/rabbitmq"
	"project-tfs02/api/utils"
	"project-tfs02/api/utils/format_error"

	"github.com/gorilla/mux"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Đọc request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	// Khởi tạo user
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Chuẩn hóa user: xóa bỏ khoảng trắng, encode các kí tự đặc biệt
	user.Prepare()

	// Validate đã có ở Front-end, có thể bổ sung hoặc thay đổi vị trí validate trong tương lai
	err = user.Validate("")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Lưu user mới vào database
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		formattedError := format_error.FormatError(err.Error())

		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	// Chỉnh sửa Header (có thể có hoặc không)
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

	// Trả về response
	utils.JSON(w, http.StatusCreated, userCreated)

	// Khởi tạo rabbitMQ
	rmq := rabbitmq.CreateNewRMQ("amqp://tfs:tfs-ocg@174.138.40.239:5672/#/")

	// Khởi tạo Channel
	pCh, err := rmq.GetChannel()
	if err != nil {
		fmt.Println("Cannot get channel")
		return
	}

	// Khởi tạo producer
	producer := producer.CreateNewProducer("emailRegister", "direct", "abc", pCh)

	// Gửi email lên rabbitMQ
	producer.Send(user.Email)
	// mail.SendNoticeImportSuccessful(userCreated.Name, userCreated.Email)
	producer.Close()
}

func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, users)
}

func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := models.User{}
	userGotten, err := user.FindUserByEmail(server.DB, email)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	utils.JSON(w, http.StatusOK, userGotten)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the user id is valid
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Read the data users
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != uint32(uid) {
		utils.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := format_error.FormatError(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.JSON(w, http.StatusOK, updatedUser)
}

func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := models.User{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		utils.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	utils.JSON(w, http.StatusNoContent, "")
}
