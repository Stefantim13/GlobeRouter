package main

import (
	"database/sql"
	"log"
)

func getRoutes() []Route {
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
	return routes
}

func getTravels(routes []Route, from string, to string) []Travel {
	travels := []Travel{}
	return travels
}
