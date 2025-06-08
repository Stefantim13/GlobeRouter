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
	r.GET("/routes", routesEndpoint)
	r.GET("/routes/:id", savedRoutesEndpoint)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
