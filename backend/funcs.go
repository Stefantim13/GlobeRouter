package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello from Gin!",
	})
}

func loginEndpointPost(c *gin.Context) {
	var req UserConfirmation
	fmt.Println("ok")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("ok")
	email := req.Email
	password := req.Password
	loginError := ""
	fmt.Println("ok")

	if !strings.Contains(email, "@") {
		loginError = "Invalid email address!"
	}

	fmt.Println("ok")
	row, err := config.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok")
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	fmt.Println("ok")
	for row.Next() {
		user := User{}
		err := row.Scan(&user.Email, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ok")

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			log.Fatal(err)
			loginError = "Invalid password!"
		}
		fmt.Println("ok")
	}
	fmt.Println("ok")
	if loginError != "" {
		c.JSON(http.StatusOK, gin.H{
			"error": loginError,
		})
	}
	fmt.Println("ok")
	c.JSON(http.StatusOK, gin.H{})
}

func signupEndpointPost(c *gin.Context) {
	var req UserConfirmation
	fmt.Println("ok")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("ok")
	email := req.Email
	password := req.Password
	confirmationPassword := req.ConfirmationPassword
	signUpError := ""
	fmt.Println("ok")

	if !strings.Contains(email, "@") {
		signUpError = "Invalid email address!"
	} else if password != confirmationPassword {
		signUpError = "Passwords do not match!"
	}
	fmt.Println("ok")

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(email)
	fmt.Println(password)

	_, err = config.db.Exec("INSERT INTO users(email, password) VALUES(?, ?)", email, string(encryptedPassword))
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	}
	fmt.Println("ok")
	if signUpError != "" {
		c.JSON(http.StatusOK, gin.H{
			"error": signUpError,
		})
	}
	fmt.Println("ok")
	c.JSON(http.StatusOK, gin.H{})
}

func routesEndpoint(c *gin.Context) {
	var route Route
	if err := c.ShouldBindJSON(&route); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	//from, to := c.Query("from"), c.Query("to")
	//routes := getRoutes()
	//travels := getTravels(routes, from, to)
	c.JSON(http.StatusOK, gin.H{"message": "Ruta a fost salvată!"})
}

func saveRouteEndpoint(c *gin.Context) {
	var req RouteRequest
	fmt.Println("ok")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Println("ok")
	route := req.Route
	email := req.Email
	fmt.Println("ok")

	var tid int64
	row := config.db.QueryRow(
		"SELECT id FROM travels WHERE cities = ? AND totalCost = ? AND duration = ? AND departure = ? AND arrival = ?",
		route.Path, route.Price, route.Duration, route.Departure, route.Arrival,
	)
	err := row.Scan(&tid)
	if err == sql.ErrNoRows {
		res, err := config.db.Exec(
			"INSERT INTO travels(cities, totalCost, duration, departure, arrival) VALUES(?, ?, ?, ?, ?)",
			route.Path, route.Price, route.Duration, route.Departure, route.Arrival,
		)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Eroare la salvarea rutei"})
			return
		}
		tid, err = res.LastInsertId()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": "Eroare la salvarea rutei"})
			return
		}
	} else if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Eroare la verificarea rutei"})
		return
	}
	_, err = config.db.Exec("INSERT OR IGNORE INTO savedTravels(tid, uid) VALUES(?, ?)", tid, email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Eroare la salvarea rutei"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func deleteRouteEndpoint(c *gin.Context) {
	var req RouteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	route := req.Route
	email := req.Email

	// Găsește id-ul rutei
	row := config.db.QueryRow(
		"SELECT id FROM travels WHERE cities = ? AND totalCost = ? AND duration = ? AND departure = ? AND arrival = ?",
		route.Path, route.Price, route.Duration, route.Departure, route.Arrival,
	)
	var tid int64
	err := row.Scan(&tid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Ruta nu a fost găsită"})
		return
	}
	// Șterge legătura dintre utilizator și rută
	_, err = config.db.Exec("DELETE FROM savedTravels WHERE tid = ? AND uid = ?", tid, email)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Eroare la ștergerea rutei"})
		return
	}
	// Șterge ruta dacă nu mai există nicio legătură
	row2 := config.db.QueryRow("SELECT COUNT(*) FROM savedTravels WHERE tid = ?", tid)
	var count int
	row2.Scan(&count)
	if count == 0 {
		config.db.Exec("DELETE FROM travels WHERE id = ?", tid)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func savedTravelsEndpoint(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	row, err := config.db.Query("SELECT t.cities, t.totalCost, t.duration, t.departure, t.arrival FROM travels t, savedTravels ts WHERE t.id = ts.tid AND ts.uid = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	var travels []Travel
	for row.Next() {
		travel := Travel{}
		err := row.Scan(&travel.Cities, &travel.TotalCost, &travel.Duration, &travel.Departure, &travel.Arrival)
		if err != nil {
			log.Fatal(err)
		}
		travels = append(travels, travel)
	}
	for _, travel := range travels {
		fmt.Println(travel.Cities, travel.TotalCost, travel.Duration, travel.Departure, travel.Arrival)
	}
	c.JSON(http.StatusOK, gin.H{
		"travels": travels,
	})
}

func bestRouteEndpoint(c *gin.Context) {
	var req struct {
		From string `json:"from"`
		To   string `json:"to"`
		By   string `json:"by"` // "price" or "duration"
		N    int    `json:"n"`  // numărul de rute dorite
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	n := req.N
	if n <= 0 {
		n = 5 // fallback la 5 rute dacă nu e specificat sau e invalid
	}
	paths, costs := getFirstKBestRoutesAsRoutes(req.From, req.To, req.By, n)
	if len(paths) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "No route found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"routes": paths,
		"total":  costs,
	})
}

// Serve worldcities.csv for React frontend
func serveWorldCitiesCSV(c *gin.Context) {
	c.File("./worldcities.csv")
}
