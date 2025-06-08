package main

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func loginEndpointPost(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	loginError := ""

	if !strings.Contains(email, "@") {
		loginError = "Invalid email address!"
	}

	row, err := config.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	for row.Next() {
		user := User{}
		err := row.Scan(&user.email, &user.password)
		if err != nil {
			log.Fatal(err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(user.password))
		if err != nil {
			log.Fatal(err)
			loginError = "Invalid password!"
		}
	}
	if loginError != "" {
		c.JSON(http.StatusOK, gin.H{
			"error": loginError,
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func signupEndpointPost(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	confirmationPassword := c.Query("confirmationPassword")
	signUpError := ""

	if !strings.Contains(email, "@") {
		signUpError = "Invalid email address!"
	} else if password != confirmationPassword {
		signUpError = "Passwords do not match!"
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}

	_, err = config.db.Exec("INSERT INTO users(email, password) VALUES(?, ?)", email, string(encryptedPassword))
	if err != nil {
		log.Fatal(err)
	}
	if signUpError != "" {
		c.JSON(http.StatusOK, gin.H{
			"error": signUpError,
		})
	}
	c.JSON(http.StatusOK, gin.H{})
}

func routesEndpoint(c *gin.Context) {
	row, err := config.db.Query("SELECT * FROM routes")
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	var routes []Route
	for row.Next() {
		route := Route{}
		err := row.Scan(&route.cost, &route.departure, &route.arrival)
		if err != nil {
			log.Fatal(err)
		}
		routes = append(routes, route)
	}
	c.JSON(http.StatusOK, gin.H{
		"routes": routes,
	})
}

func savedRoutesEndpoint(c *gin.Context) {
	id := c.Query("id")
	row, err := config.db.Query("SELECT * FROM routes r, savedRoutes sr WHERE r.id == sr.rid AND sr.uid == ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	var routes []Route
	for row.Next() {
		route := Route{}
		err := row.Scan(&route.cost, &route.departure, &route.arrival)
		if err != nil {
			log.Fatal(err)
		}
		routes = append(routes, route)
	}
	c.JSON(http.StatusOK, gin.H{
		"routes": routes,
	})
}
