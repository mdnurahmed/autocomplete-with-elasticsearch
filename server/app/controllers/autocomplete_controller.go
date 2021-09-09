package controllers

import (
	"autocomplete/app/DTO"
	"autocomplete/app/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type IAutocompleteController interface {
	Search(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
}

type AutocompleteController struct {
	autocompleteService services.IAutocompleteService
}

func NewInstanceOfAutocompleteController(
	autocompleteService services.IAutocompleteService) AutocompleteController {
	return AutocompleteController{autocompleteService: autocompleteService}
}

func (a *AutocompleteController) Search(c *gin.Context) {
	m := DTO.Request{}
	m.SearchString = c.DefaultQuery("Word", "")
	if m.SearchString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("NoWordProvided", "searchStringhas to be provided"),
		})
		return
	}
	result, err := a.autocompleteService.Search(m.SearchString)
	if err != nil {
		log.WithFields(log.Fields{
			"error_message": err.Error(),
		}).Error("error in controller ")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("ServiceError", "Something Went Wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})

}

func (a *AutocompleteController) Insert(c *gin.Context) {
	m := DTO.Request{}
	err := c.Bind(&m)
	if err != nil || m.SearchString == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("NoWordProvided", "searchStringhas to be provided"),
		})
		return
	}
	err = a.autocompleteService.Insert(m.SearchString)
	if err != nil {
		log.WithFields(log.Fields{
			"error_message": err.Error(),
		}).Error("error in controller ")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("ServiceError", "Something Went Wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "searchStringinserted successfully",
	})
}

func (a *AutocompleteController) Delete(c *gin.Context) {
	fmt.Println("~~~~~~~~~~~~~~")
	err := a.autocompleteService.Delete()
	fmt.Println(err)
	if err != nil {
		log.WithFields(log.Fields{
			"erro_message": err.Error(),
		}).Error("error in controller ")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": DTO.NewErrorResponse("ServiceError", "Something Went Wrong"),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Deleted successfully",
	})
}
