package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Citește orașele din worldcities.csv (doar orașe din Europa, unice)
func GetEuropeanCitiesFromCSV() ([]string, error) {
	file, err := os.Open("./worldcities.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	cityIdx, countryIdx, latIdx, lngIdx := -1, -1, -1, -1
	header := records[0]
	for i, col := range header {
		switch col {
		case "city":
			cityIdx = i
		case "country":
			countryIdx = i
		case "lat":
			latIdx = i
		case "lng":
			lngIdx = i
		}
	}
	if cityIdx == -1 || countryIdx == -1 || latIdx == -1 || lngIdx == -1 {
		return nil, fmt.Errorf("required columns not found")
	}
	citySet := make(map[string]struct{})
	for _, row := range records[1:] {
		if len(row) > latIdx && len(row) > lngIdx && len(row) > countryIdx && len(row) > cityIdx {
			lat, _ := strconv.ParseFloat(row[latIdx], 64)
			lng, _ := strconv.ParseFloat(row[lngIdx], 64)
			country := row[countryIdx]
			city := row[cityIdx]
			// Europa: lat 34-72, lng -25 to 45, sau country in lista de țări europene
			if (lat >= 34 && lat <= 72 && lng >= -25 && lng <= 45) || isEuropeanCountry(country) {
				if city != "" {
					citySet[city] = struct{}{}
				}
			}
		}
	}
	var cities []string
	for city := range citySet {
		cities = append(cities, city)
	}
	return cities, nil
}

// Listă simplă de țări europene (poți extinde după nevoie)
func isEuropeanCountry(country string) bool {
	europe := map[string]struct{}{
		"Romania": {}, "Germany": {}, "France": {}, "Austria": {}, "Hungary": {}, "Netherlands": {},
		"Belgium": {}, "Italy": {}, "Spain": {}, "Portugal": {}, "Switzerland": {}, "Poland": {},
		"Czechia": {}, "Czech Republic": {}, "Slovakia": {}, "Slovenia": {}, "Croatia": {}, "Serbia": {},
		"Bulgaria": {}, "Greece": {}, "Denmark": {}, "Sweden": {}, "Norway": {}, "Finland": {},
		"Ukraine": {}, "Belarus": {}, "Moldova": {}, "Russia": {}, "Estonia": {}, "Latvia": {},
		"Lithuania": {}, "Ireland": {}, "United Kingdom": {}, "UK": {}, "Bosnia and Herzegovina": {},
		"Albania": {}, "North Macedonia": {}, "Montenegro": {}, "Luxembourg": {}, "Monaco": {},
		"Andorra": {}, "San Marino": {}, "Vatican City": {}, "Liechtenstein": {}, "Kosovo": {},
		"Iceland": {},
	}
	_, ok := europe[country]
	return ok
}

// Populează tabela routes cu rute random, dar asigură conectivitate (graful să fie conex pentru Dijkstra)
func PopulateRandomRoutesFromCSV(n int) {
	cities, err := GetEuropeanCitiesFromCSV()
	if err != nil {
		fmt.Println("Eroare la citirea orașelor:", err)
		return
	}
	if len(cities) < 2 {
		fmt.Println("Nu sunt suficiente orașe europene în worldcities.csv")
		return
	}
	rand.Seed(time.Now().UnixNano())

	// 1. Creează un lanț (conexiune minimă) ca să fie sigur că toate orașele sunt conectate
	perm := rand.Perm(len(cities))
	for i := 0; i < len(cities)-1; i++ {
		from := cities[perm[i]]
		to := cities[perm[i+1]]
		departure := rand.Intn(24) * 100
		duration := 2 + rand.Intn(8)
		arrival := (departure + duration*100) % 2400
		price := 30 + rand.Intn(120)
		_, err := config.db.Exec(
			"INSERT INTO routes(fromCity, toCity, departure, arrival, duration, price) VALUES (?, ?, ?, ?, ?, ?)",
			from, to, departure, arrival, duration, price,
		)
		if err != nil {
			fmt.Printf("Error inserting route %s -> %s: %v\n", from, to, err)
		}
	}

	// 2. Adaugă rute suplimentare random pentru diversitate
	for i := 0; i < n; i++ {
		from := cities[rand.Intn(len(cities))]
		to := cities[rand.Intn(len(cities))]
		if from == to {
			continue
		}
		departure := rand.Intn(24) * 100
		duration := 2 + rand.Intn(8)
		arrival := (departure + duration*100) % 2400
		price := 30 + rand.Intn(120)
		_, err := config.db.Exec(
			"INSERT INTO routes(fromCity, toCity, departure, arrival, duration, price) VALUES (?, ?, ?, ?, ?, ?)",
			from, to, departure, arrival, duration, price,
		)
		if err != nil {
			fmt.Printf("Error inserting route %s -> %s: %v\n", from, to, err)
		}
	}
}

// Populează tabela routes cu un graf complet pentru n orașe din Europa, cu valori random
func PopulateCompleteGraphFromCSV(n int) {
	cities, err := GetEuropeanCitiesFromCSV()
	if err != nil {
		fmt.Println("Eroare la citirea orașelor:", err)
		return
	}
	if len(cities) < n {
		fmt.Printf("Nu sunt suficiente orașe europene în worldcities.csv (cerute: %d, găsite: %d)\n", n, len(cities))
		return
	}
	rand.Seed(time.Now().UnixNano())
	// Selectează random n orașe distincte
	perm := rand.Perm(len(cities))
	selected := make([]string, n)
	for i := 0; i < n; i++ {
		selected[i] = cities[perm[i]]
	}
	// Graf complet: fiecare oraș către fiecare alt oraș (orientat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			from := selected[i]
			to := selected[j]
			departure := rand.Intn(24) * 100
			duration := 2 + rand.Intn(8)
			arrival := (departure + duration*100) % 2400
			price := 30 + rand.Intn(120)
			_, err := config.db.Exec(
				"INSERT INTO routes(fromCity, toCity, departure, arrival, duration, price) VALUES (?, ?, ?, ?, ?, ?)",
				from, to, departure, arrival, duration, price,
			)
			if err != nil {
				fmt.Printf("Error inserting route %s -> %s: %v\n", from, to, err)
			}
		}
	}
}

// Generează un graf aproape complet (densitate mare, dar nu toate muchiile) cu 50 de orașe random din Europa
func PopulateAlmostCompleteGraphFromCSV() {
	const n = 50
	const density = 0.7 // 70% din muchiile posibile

	cities, err := GetEuropeanCitiesFromCSV()
	if err != nil {
		fmt.Println("Eroare la citirea orașelor:", err)
		return
	}
	if len(cities) < n {
		fmt.Printf("Nu sunt suficiente orașe europene în worldcities.csv (cerute: %d, găsite: %d)\n", n, len(cities))
		return
	}
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(cities))
	selected := make([]string, n)
	for i := 0; i < n; i++ {
		selected[i] = cities[perm[i]]
	}
	// Pentru fiecare pereche (i, j), adaugă muchie cu probabilitate density
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if rand.Float64() < density {
				from := selected[i]
				to := selected[j]
				departure := rand.Intn(24) * 100
				duration := 2 + rand.Intn(8)
				arrival := (departure + duration*100) % 2400
				price := 30 + rand.Intn(120)
				_, err := config.db.Exec(
					"INSERT INTO routes(fromCity, toCity, departure, arrival, duration, price) VALUES (?, ?, ?, ?, ?, ?)",
					from, to, departure, arrival, duration, price,
				)
				if err != nil {
					fmt.Printf("Error inserting route %s -> %s: %v\n", from, to, err)
				}
			}
		}
	}
}
