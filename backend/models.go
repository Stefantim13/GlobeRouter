package main

import "database/sql"

/* type Date struct {
	day, month, year string
} */

type Config struct {
	db *sql.DB
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserConfirmation struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	ConfirmationPassword string `json:"confirmationPassword"`
}

type Route struct {
	Path      string `json:"path"`
	Price     int    `json:"price"`
	Duration  int    `json:"duration"`
	Departure int    `json:"departure"`
	Arrival   int    `json:"arrival"`
}

type Travel struct {
	Cities    string `json:"cities"`
	TotalCost int    `json:"totalCost"`
	Duration  int    `json:"duration"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
}

type RouteR struct {
	Path      string `json:"path"`
	Price     int    `json:"price"`
	Duration  int    `json:"duration"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
}

type RouteRequest struct {
	Email string `json:"email"`
	Route RouteR `json:"route"`
}

type Transport struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Departure int    `json:"departure"`
	Arrival   int    `json:"arrival"`
	Duration  int    `json:"duration"`
	Price     int    `json:"price"`
}

type Place struct {
	CurrentCity string `json:"currentCity"`
	Arrival     int    `json:"arrival"`
	Duration    int    `json:"duration"`
	Price       int    `json:"price"`
}
