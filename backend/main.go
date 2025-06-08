package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

var config Config

func main() {
	config.db = initDB()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(config.db)

	r := gin.New()
	r.Use(gin.Logger())

	r.POST("/login", loginEndpointPost)
	r.POST("/signup", signupEndpointPost)
	r.GET("/routes/:from/:to", routesEndpoint)
	r.POST("/routes/:from/:to", saveTravelEndpoint)
	r.GET("/:id", savedTravelsEndpoint)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
