package main

import (
	"github.com/gin-gonic/gin"
	"go_rest/grocery"
	"go_rest/model"
	"log"
)

func main() {
	db, err := model.Database()
	if err != nil {
		log.Println(err)
	}

	db.DB()
	router := gin.Default()

	router.GET("/groceries", grocery.GetGroceries)
	router.GET("/grocery/:id", grocery.GetGrocery)
	router.POST("/grocery", grocery.PostGrocery)
	router.PUT("/grocery/:id", grocery.UpdateGrocery)
	router.DELETE("/grocery/:id", grocery.DeleteGrocery)

	log.Fatal(router.Run(":8000"))
}
