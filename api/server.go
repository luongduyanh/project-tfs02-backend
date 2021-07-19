package api

import (
	"fmt"
	"log"
	"os"

	"project-tfs02/api/controllers"
	"project-tfs02/api/models"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var server = controllers.Server{}

//tao cac table
func Load(db *gorm.DB) {

	//bang nao co truoc thi tao dau tien
	db.Debug().AutoMigrate(&models.User{}, &models.Brand{}, &models.Category{})
	db.Debug().AutoMigrate(&models.Post{}, &models.Order{}, &models.Ship{}, &models.SubCategory{}, &models.Product{})
	db.Debug().AutoMigrate(&models.Option{}, &models.Variant{}, &models.OrderLine{})
	db.Debug().AutoMigrate(&models.OptionValue{}, &models.Image{})
	db.Debug().AutoMigrate(&models.ProductOption{})
	//---------------------------------------------------------------------------------------------//
}

func Run() {
	var err = godotenv.Load(os.ExpandEnv("./.env"))
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("Getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Load(server.DB)

	server.Run(":8000")

}
