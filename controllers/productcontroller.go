package productcontroller

import (
	"net/http"

	"github.com/aidilfitra08/go-rest-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(context *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	context.JSON(http.StatusOK, gin.H{"products": products})

	
}

func Show(context *gin.Context) {
 	var product models.Product
	id := context.Param("id")
	if error := models.DB.First(&product, id).Error; error != nil {
		switch error {
		case gorm.ErrRecordNotFound:
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data not found"})
			return
		default:
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		}
	}
	context.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(context *gin.Context) {

}

func Update(context *gin.Context) {

}

func Delete(context *gin.Context) {

}