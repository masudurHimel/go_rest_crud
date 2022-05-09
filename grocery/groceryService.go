package grocery

import "C"
import (
	"github.com/gin-gonic/gin"
	"go_rest/model"
	"log"
	"net/http"
)

type NewGrocery struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type UpdateGroceries struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

func GetGroceries(c *gin.Context) {
	var groceries []model.Grocery

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&groceries).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, groceries)
}

func GetGrocery(c *gin.Context) {

	var grocery model.Grocery

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found"})
		return
	}

	c.JSON(http.StatusOK, grocery)

}

func PostGrocery(c *gin.Context) {
	var grocery NewGrocery
	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGrocery := model.Grocery{Name: grocery.Name, Quantity: grocery.Quantity}

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newGrocery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newGrocery)
}

func UpdateGrocery(c *gin.Context) {
	var grocery model.Grocery

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found!"})
		return
	}

	var updateGrocery UpdateGroceries

	if err := c.ShouldBindJSON(&updateGrocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&grocery).Updates(model.Grocery{Name: updateGrocery.Name, Quantity: updateGrocery.Quantity}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grocery)
}

func DeleteGrocery(c *gin.Context) {
	var grocery model.Grocery

	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&grocery).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found!"})
		return
	}
	if err := db.Delete(&grocery).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Grocery deleted"})
}