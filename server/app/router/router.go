package router

import (
	"autocomplete/app/controllers"
	mw "autocomplete/app/middlewares"
	"autocomplete/app/repositories"
	"autocomplete/app/services"
	"autocomplete/app/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//InitializeApp initilizes the app like loading configuration,
//environment variables , setting up routes for apis
func InitializeApp() *gin.Engine {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	if os.Getenv("LogToFile") == "True" {
		f, err := os.OpenFile("logs", os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.WithFields(log.Fields{
				"erro_message": err.Error(),
			}).Fatal("Couldn't create a file for logging")
		}
		log.SetOutput(f)
	}

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.WithFields(log.Fields{
			"erro_message": err.Error(),
		}).Fatal("Couldn't load config object")
	}

	utils.Configuration = config
	log.WithFields(log.Fields{
		"Config": config,
	}).Info("Configuration object")
	esrepo := repositories.NewInstanceOfElasticSearchRepository(
		utils.Configuration.Address,
		utils.Configuration.IndexName,
		utils.Configuration.Refresh,
		utils.Configuration.Size)
	esrepo.Bootstrap()
	autocompleteService := services.NewInstanceOfAutocompleteService(&esrepo)
	autocompleteController := controllers.NewInstanceOfAutocompleteController(&autocompleteService)

	r := gin.Default()
	r.Use(mw.CORSMiddleware())
	r.POST(
		"/insert",
		autocompleteController.Insert)
	r.GET(
		"/search",
		autocompleteController.Search)
	r.POST(
		"/delete",
		autocompleteController.Delete)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PONG",
		})
	})
	return r
}
