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
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data not found"})
			return
		default:
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	context.JSON(http.StatusOK, gin.H{"product": product})
}

func Create(context *gin.Context) {
	var product models.Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&product)
	context.JSON(http.StatusOK, gin.H{"product": product})
}

func Update(context *gin.Context) {
	var product models.Product
	id := context.Param("id")

	if err := context.ShouldBindJSON(&product); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

func Delete(context *gin.Context) {
	var product models.Product
	// id := context.Param("id")
	id := context.Query("id")

	if err := context.ShouldBindJSON(context.Request.URL); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Delete(&product, id).RowsAffected == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}