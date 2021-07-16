package controllers

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

var err error

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		sqlDB, _ := sql.Open(Dbdriver, DBURL)
		server.DB, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to the %s database", Dbdriver)
		}
	}
	server.Router = mux.NewRouter().StrictSlash(true)

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"*"},
	}).Handler(server.Router)
	fmt.Println("Listening to port: ", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
