package main

import (
	"database/sql"
	"fmt"
	"log"
)

// Fetch all Transport nodes from the DB
func getTransports() []Transport {
	row, err := config.db.Query("SELECT fromCity, toCity, departure, arrival, duration, price FROM routes")
	if err != nil {
		log.Fatal(err)
	}
	defer func(row *sql.Rows) {
		err := row.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(row)

	var transports []Transport
	for row.Next() {
		t := Transport{}
		err := row.Scan(&t.From, &t.To, &t.Departure, &t.Arrival, &t.Duration, &t.Price)
		if err != nil {
			log.Fatal(err)
		}
		transports = append(transports, t)
	}
	return transports
}

// Dijkstra's algorithm for best route by price
func getBestRouteByPrice(transports []Transport, from string, to string) ([]Transport, int) {
	type State struct {
		City string
		Cost int
		Path []Transport
	}
	graph := make(map[string][]Transport)
	for _, t := range transports {
		graph[t.From] = append(graph[t.From], t)
	}
	type PQItem struct {
		State
		Priority int
	}
	pq := []PQItem{{State: State{City: from, Cost: 0, Path: []Transport{}}, Priority: 0}}
	visited := make(map[string]int)

	for len(pq) > 0 {
		curr := pq[0]
		pq = pq[1:]
		if v, ok := visited[curr.City]; ok && curr.Cost > v {
			continue
		}
		visited[curr.City] = curr.Cost
		if curr.City == to {
			return curr.Path, curr.Cost
		}
		for _, t := range graph[curr.City] {
			var edgeCost int
			edgeCost = t.Price
			next := State{
				City: t.To,
				Cost: curr.Cost + edgeCost,
				Path: append(append([]Transport{}, curr.Path...), t),
			}
			// Insert in pq sorted by cost (simple for small graphs)
			inserted := false
			for i, item := range pq {
				if next.Cost < item.Priority {
					pq = append(pq[:i], append([]PQItem{{State: next, Priority: next.Cost}}, pq[i:]...)...)
					inserted = true
					break
				}
			}
			if !inserted {
				pq = append(pq, PQItem{State: next, Priority: next.Cost})
			}
		}
	}
	return nil, -1 // no route found
}

// Dijkstra's algorithm for best route by total duration (final arrival - starting departure)
func getBestRouteByDuration(transports []Transport, from string, to string) ([]Transport, int) {
	type State struct {
		City     string
		Path     []Transport
		FirstDep int
		LastArr  int
	}
	graph := make(map[string][]Transport)
	for _, t := range transports {
		graph[t.From] = append(graph[t.From], t)
	}
	type PQItem struct {
		State
		Priority int // total duration
	}
	pq := []PQItem{{State: State{City: from, Path: []Transport{}, FirstDep: -1, LastArr: -1}, Priority: 0}}
	visited := make(map[string]int)

	for len(pq) > 0 {
		curr := pq[0]
		pq = pq[1:]
		// Use LastArr-FirstDep as cost
		var currDuration int
		if curr.FirstDep != -1 && curr.LastArr != -1 {
			currDuration = curr.LastArr - curr.FirstDep
		} else {
			currDuration = 0
		}
		if v, ok := visited[curr.City]; ok && currDuration > v {
			continue
		}
		visited[curr.City] = currDuration
		if curr.City == to && len(curr.Path) > 0 {
			return curr.Path, currDuration
		}
		for _, t := range graph[curr.City] {
			// Only allow next departure after current arrival (if not first)
			if curr.LastArr != -1 && t.Departure < curr.LastArr {
				continue
			}
			firstDep := curr.FirstDep
			if firstDep == -1 {
				firstDep = t.Departure
			}
			next := State{
				City:     t.To,
				Path:     append(append([]Transport{}, curr.Path...), t),
				FirstDep: firstDep,
				LastArr:  t.Arrival,
			}
			// Insert in pq sorted by duration
			nextDuration := t.Arrival - firstDep
			inserted := false
			for i, item := range pq {
				if nextDuration < item.Priority {
					pq = append(pq[:i], append([]PQItem{{State: next, Priority: nextDuration}}, pq[i:]...)...)
					inserted = true
					break
				}
			}
			if !inserted {
				pq = append(pq, PQItem{State: next, Priority: nextDuration})
			}
		}
	}
	return nil, -1 // no route found
}

// Find multiple best routes (not just one) from 'from' to 'to'.
// Returns up to maxRoutes best routes by price or duration (no cycles).
func getMultipleBestRoutes(transports []Transport, from string, to string, by string, maxRoutes int) ([][]Transport, []int) {
	type State struct {
		City     string
		Cost     int
		Path     []Transport
		Visited  map[string]bool
		FirstDep int
		LastArr  int
	}
	graph := make(map[string][]Transport)
	for _, t := range transports {
		graph[t.From] = append(graph[t.From], t)
	}
	type PQItem struct {
		State
		Priority int
	}
	pq := []PQItem{{State: State{
		City: from, Cost: 0, Path: []Transport{}, Visited: map[string]bool{from: true}, FirstDep: -1, LastArr: -1,
	}, Priority: 0}}
	var foundRoutes [][]Transport
	var foundCosts []int

	for len(pq) > 0 && len(foundRoutes) < maxRoutes {
		curr := pq[0]
		pq = pq[1:]
		// For duration, cost is total trip time; for price, cost is sum of prices
		var currCost int
		if by == "duration" {
			if curr.FirstDep != -1 && curr.LastArr != -1 {
				currCost = curr.LastArr - curr.FirstDep
			} else {
				currCost = 0
			}
		} else {
			currCost = curr.Cost
		}
		if curr.City == to && len(curr.Path) > 0 {
			foundRoutes = append(foundRoutes, curr.Path)
			foundCosts = append(foundCosts, currCost)
			continue
		}
		for _, t := range graph[curr.City] {
			// Avoid cycles
			if curr.Visited[t.To] {
				continue
			}
			// For duration, only allow next departure after current arrival
			if by == "duration" && curr.LastArr != -1 && t.Departure < curr.LastArr {
				continue
			}
			firstDep := curr.FirstDep
			if by == "duration" && firstDep == -1 {
				firstDep = t.Departure
			}
			nextCost := curr.Cost + t.Price
			if by == "duration" {
				nextCost = curr.Cost + t.Duration
			}
			nextVisited := make(map[string]bool)
			for k, v := range curr.Visited {
				nextVisited[k] = v
			}
			nextVisited[t.To] = true
			next := State{
				City:     t.To,
				Cost:     nextCost,
				Path:     append(append([]Transport{}, curr.Path...), t),
				Visited:  nextVisited,
				FirstDep: firstDep,
				LastArr:  t.Arrival,
			}
			// Insert in pq sorted by cost/duration
			var priority int
			if by == "duration" && firstDep != -1 {
				priority = t.Arrival - firstDep
			} else {
				priority = nextCost
			}
			inserted := false
			for i, item := range pq {
				if priority < item.Priority {
					pq = append(pq[:i], append([]PQItem{{State: next, Priority: priority}}, pq[i:]...)...)
					inserted = true
					break
				}
			}
			if !inserted {
				pq = append(pq, PQItem{State: next, Priority: priority})
			}
		}
	}
	return foundRoutes, foundCosts
}

// Find exactly the first k best routes (by price or duration, no cycles).
func getFirstKBestRoutes(transports []Transport, from string, to string, by string, k int) ([][]Transport, []int) {
	type State struct {
		City     string
		Cost     int
		Path     []Transport
		Visited  map[string]bool
		FirstDep int
		LastArr  int
	}
	graph := make(map[string][]Transport)
	for _, t := range transports {
		graph[t.From] = append(graph[t.From], t)
	}
	type PQItem struct {
		State
		Priority int
	}
	pq := []PQItem{{State: State{
		City: from, Cost: 0, Path: []Transport{}, Visited: map[string]bool{from: true}, FirstDep: -1, LastArr: -1,
	}, Priority: 0}}
	var foundRoutes [][]Transport
	var foundCosts []int

	for len(pq) > 0 && len(foundRoutes) < k {
		curr := pq[0]
		pq = pq[1:]
		var currCost int
		if by == "duration" {
			if curr.FirstDep != -1 && curr.LastArr != -1 {
				currCost = curr.LastArr - curr.FirstDep
			} else {
				currCost = 0
			}
		} else {
			currCost = curr.Cost
		}
		if curr.City == to && len(curr.Path) > 0 {
			foundRoutes = append(foundRoutes, curr.Path)
			foundCosts = append(foundCosts, currCost)
			continue
		}
		for _, t := range graph[curr.City] {
			if curr.Visited[t.To] {
				continue
			}
			if by == "duration" && curr.LastArr != -1 && t.Departure < curr.LastArr {
				continue
			}
			firstDep := curr.FirstDep
			if by == "duration" && firstDep == -1 {
				firstDep = t.Departure
			}
			nextCost := curr.Cost + t.Price
			if by == "duration" {
				nextCost = curr.Cost + t.Duration
			}
			nextVisited := make(map[string]bool)
			for k, v := range curr.Visited {
				nextVisited[k] = v
			}
			nextVisited[t.To] = true
			next := State{
				City:     t.To,
				Cost:     nextCost,
				Path:     append(append([]Transport{}, curr.Path...), t),
				Visited:  nextVisited,
				FirstDep: firstDep,
				LastArr:  t.Arrival,
			}
			var priority int
			if by == "duration" && firstDep != -1 {
				priority = t.Arrival - firstDep
			} else {
				priority = nextCost
			}
			inserted := false
			for i, item := range pq {
				if priority < item.Priority {
					pq = append(pq[:i], append([]PQItem{{State: next, Priority: priority}}, pq[i:]...)...)
					inserted = true
					break
				}
			}
			if !inserted {
				pq = append(pq, PQItem{State: next, Priority: priority})
			}
		}
	}
	return foundRoutes, foundCosts
}

// Wrapper for frontend: returns first k best routes as []Route slices
func getFirstKBestRoutesAsRoutes(from, to, by string, k int) ([]Route, []int) {
	transports := getTransports()
	paths, costs := getFirstKBestRoutes(transports, from, to, by, k)
	var result []Route
	for _, path := range paths {
		var route Route = Route{Path: "", Price: 0, Duration: 0, Departure: path[0].Departure, Arrival: path[len(path)-1].Arrival}
		var seenPaths map[string]bool = make(map[string]bool)
		var p []string = make([]string, 0)
		for _, t := range path {
			if !seenPaths[t.From] {
				seenPaths[t.From] = true
				p = append(p, t.From)
			}
			if !seenPaths[t.To] {
				seenPaths[t.To] = true
				p = append(p, t.To)
			}
			route.Price += t.Price
			route.Duration += t.Duration
		}
		for _, city := range p {
			if route.Path != "" {
				route.Path += "->"
			}
			route.Path += city
		}
		result = append(result, route)
	}
	for _, r := range result {
		fmt.Println("Route:", r.Path, "Price:", r.Price, "Duration:", r.Duration, "Departure:", r.Departure, "Arrival:", r.Arrival)
	}
	return result, costs
}
