import React, { useEffect, useState, useRef } from 'react';

function MyRoutesPage() {
  const [routes, setRoutes] = useState([]);
  const [showMap, setShowMap] = useState(false);
  const [mapRoute, setMapRoute] = useState(null);
  const [cityCoords, setCityCoords] = useState({});
  const mapRef = useRef(null);
  const leafletRef = useRef(null);
  const email = localStorage.getItem('email');

  useEffect(() => {
    async function fetchRoutes() {
      if (!email) return;
      const res = await fetch(`http://localhost:8080/${encodeURIComponent(email)}`);
      const data = await res.json();
      if (data.travels) setRoutes(data.travels);
    }
    fetchRoutes();
  }, [email]);

  // Fetch city coordinates from a CSV file (cities.csv) in public folder
  useEffect(() => {
    async function fetchCoords() {
      const allCities = new Set();
      routes.forEach(route => {
        (route.cities || '').split('->').forEach(city => allCities.add(city.trim()));
      });
      // Citește CSV-ul cu fetch
      const res = await fetch('./worldcities.csv');
      const text = await res.text();
      const lines = text.split('\n');
      // CSV-ul are header: city,city_ascii,lat,lng,country,...
      const header = lines[0].replace(/"/g, '').split(',');
      const cityIndex = header.indexOf('city');
      const cityAsciiIndex = header.indexOf('city_ascii');
      const countryIndex = header.indexOf('country');
      const latIndex = header.indexOf('lat');
      const lngIndex = header.indexOf('lng');
      const coordsObj = {};
      for (let i = 1; i < lines.length; i++) {
        // Folosește regex pentru a separa corect valorile cu virgule în ghilimele
        const cols = lines[i].match(/(".*?"|[^",\s]+)(?=\s*,|\s*$)/g);
        if (!cols || cols.length < Math.max(cityIndex, cityAsciiIndex, countryIndex, latIndex, lngIndex) + 1) continue;
        // Ia atât city cât și city_ascii pentru potrivire flexibilă, și țara
        const cityName = cols[cityIndex].replace(/"/g, '').trim();
        const cityAscii = cols[cityAsciiIndex].replace(/"/g, '').trim();
        const country = cols[countryIndex].replace(/"/g, '').trim();
        const lat = parseFloat(cols[latIndex].replace(/"/g, ''));
        const lng = parseFloat(cols[lngIndex].replace(/"/g, ''));
        if (cityName && country && !isNaN(lat) && !isNaN(lng)) {
          // Cheie: city|country și city_ascii|country
          coordsObj[`${cityName}|${country}`] = [lat, lng];
          coordsObj[`${cityAscii}|${country}`] = [lat, lng];
        }
      }
      // Completează doar orașele din rute
      const finalCoords = {};
      for (const city of allCities) {
        // Permite utilizatorului să scrie și "Berlin, Germany" sau "Berlin|Germany"
        let cityKey = city;
        let country = '';
        if (city.includes('|')) {
          [cityKey, country] = city.split('|').map(s => s.trim());
        } else if (city.includes(',')) {
          [cityKey, country] = city.split(',').map(s => s.trim());
        }
        // Caută cu și fără diacritice, cu și fără țară
        let found = false;
        if (country) {
          if (coordsObj[`${cityKey}|${country}`]) {
            finalCoords[city] = coordsObj[`${cityKey}|${country}`];
            found = true;
          } else if (coordsObj[`${cityKey.normalize('NFD').replace(/[\u0300-\u036f]/g, '')}|${country}`]) {
            finalCoords[city] = coordsObj[`${cityKey.normalize('NFD').replace(/[\u0300-\u036f]/g, '')}|${country}`];
            found = true;
          }
        }
        if (!found) {
          // Caută fără țară
          const allKeys = Object.keys(coordsObj).filter(k => k.startsWith(cityKey + '|'));
          if (allKeys.length === 1) {
            finalCoords[city] = coordsObj[allKeys[0]];
          } else if (allKeys.length > 1) {
            // Dacă există mai multe, preferă una din Europa sau România
            const prefer = allKeys.find(k => k.endsWith('|Romania') || k.endsWith('|Germany') || k.endsWith('|France') || k.endsWith('|Austria'));
            finalCoords[city] = coordsObj[prefer] || coordsObj[allKeys[0]];
          } else if (coordsObj[`${cityKey}|`]) {
            finalCoords[city] = coordsObj[`${cityKey}|`];
          } else if (coordsObj[`${cityKey.normalize('NFD').replace(/[\u0300-\u036f]/g, '')}|`]) {
            finalCoords[city] = coordsObj[`${cityKey.normalize('NFD').replace(/[\u0300-\u036f]/g, '')}|`];
          } else {
            finalCoords[city] = [44.4268, 26.1025]; // fallback București
          }
        }
      }
      setCityCoords(finalCoords);
    }
    fetchCoords();
  }, [routes]);

  // Funcție pentru ștergerea unei rute salvate (un-bookmark)
  const deleteRouteFromServer = async (route) => {
    try {
      const routeToSend = {
        path: route.cities.includes('->') ? route.cities : route.cities.split(' → ').join('->'),
        price: route.totalCost,
        duration: route.duration,
      };
      await fetch('http://localhost:8080/delete-route', {
        method: 'DELETE',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, route: routeToSend }),
      });
      // Refresh la pagină după ștergere
      window.location.reload();
    } catch (err) {
      alert('Eroare la conectarea cu serverul.');
    }
  };

  // Harta Leaflet în modal
  useEffect(() => {
    if (showMap && mapRoute && mapRef.current && Object.keys(cityCoords).length > 0) {
      // Lazy load Leaflet dacă nu e deja încărcat
      if (!leafletRef.current) {
        const script = document.createElement('script');
        script.src = 'https://unpkg.com/leaflet@1.9.3/dist/leaflet.js';
        script.onload = () => {
          leafletRef.current = window.L;
          renderMap();
        };
        document.body.appendChild(script);
      } else {
        renderMap();
      }
    }
    // eslint-disable-next-line
    function renderMap() {
      const L = leafletRef.current || window.L;
      if (!L) return;
      // Șterge orice hartă anterioară
      if (mapRef.current._leaflet_id) {
        mapRef.current._leaflet_id = null;
        mapRef.current.innerHTML = '';
      }
      // Exemplu: parsează cities și pune puncte pe hartă (fără coordonate reale)
      const cities = mapRoute.cities.split('->').map(s => s.trim());
      const coords = cities.map(city => cityCoords[city] || [44.4268, 26.1025]);
      const map = L.map(mapRef.current).setView(coords[0], 6);
      L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
        attribution: '© OpenStreetMap contributors'
      }).addTo(map);
      const poly = L.polyline(coords, { weight: 5, opacity: 0.8, color: 'blue', dashArray: '10,5' }).addTo(map);
      map.fitBounds(poly.getBounds());
      coords.forEach((coord, idx) => {
        const marker = L.marker(coord).addTo(map);
        marker.bindTooltip(cities[idx], { permanent: true, direction: 'right', offset: [10, 0] });
      });
    }
  }, [showMap, mapRoute, cityCoords]);

  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>Rutele mele</h1>
      {routes.length === 0 ? (
        <p style={styles.noRoutes}>Nu ai rute disponibile.</p>
      ) : (
        <ul style={styles.list}>
          {routes.map((route, idx) => (
            <li key={idx} style={styles.card}>
              <button
                onClick={() => deleteRouteFromServer(route)}
                style={{
                  position: 'absolute',
                  top: '1rem',
                  right: '1rem',
                  border: 'none',
                  borderRadius: '4px',
                  padding: '0.5rem 1rem',
                  cursor: 'pointer',
                  fontSize: '1.3rem',
                  backgroundColor: '#ffd700',
                  color: '#333',
                  zIndex: 2,
                }}
                aria-label="Șterge din favorite"
                title="Șterge din favorite"
              >
                ★
              </button>
              {/* Butonul de hartă fix sub butonul de save, în dreapta jos */}
              <button
                onClick={() => { setMapRoute(route); setShowMap(true); }}
                style={{
                  position: 'absolute',
                  right: '1rem',
                  bottom: '1rem',
                  border: 'none',
                  borderRadius: '4px',
                  padding: '0.5rem 1rem',
                  cursor: 'pointer',
                  fontSize: '1.1rem',
                  backgroundColor: '#007bff',
                  color: '#fff',
                  zIndex: 2,
                }}
                aria-label="Vezi pe hartă"
                title="Vezi pe hartă"
              >
                🗺️
              </button>
              <p>
                <strong>Traseu:</strong> {(route.cities || '').split(',').join(' → ')}
              </p>
              <p>Preț: {route.totalCost} €</p>
              <p>Durată: {route.duration} ore</p>
              <p>Plecare: {route.departure}</p>
              <p>Sosire: {route.arrival}</p>
            </li>
          ))}
        </ul>
      )}
      {showMap && (
        <div style={styles.modalOverlay} onClick={() => setShowMap(false)}>
          <div style={styles.modalContent} onClick={e => e.stopPropagation()}>
            <div ref={mapRef} style={{ width: '600px', height: '400px' }} />
            <button style={styles.closeButton} onClick={() => setShowMap(false)}>Închide</button>
            <link rel="stylesheet" href="https://unpkg.com/leaflet@1.9.3/dist/leaflet.css" />
          </div>
        </div>
      )}
    </div>
  );
}

const styles = {
  container: {
    maxWidth: '600px',
    margin: '4rem auto',
    padding: '2rem',
    backgroundColor: '#f9f9f9',
    borderRadius: '10px',
    fontFamily: 'Arial, sans-serif',
    boxShadow: '0 0 10px rgba(0,0,0,0.1)',
  },
  heading: {
    fontSize: '1.8rem',
    marginBottom: '1.5rem',
    textAlign: 'center',
  },
  list: {
    listStyle: 'none',
    padding: 0,
    margin: 0,
  },
  card: {
    backgroundColor: '#fff',
    padding: '1rem',
    marginBottom: '1rem',
    borderRadius: '8px',
    border: '1px solid #ddd',
    textAlign: 'left',
    position: 'relative',
  },
  noRoutes: {
    textAlign: 'center',
    color: '#777',
  },
  modalOverlay: {
    position: 'fixed',
    top: 0, left: 0, right: 0, bottom: 0,
    background: 'rgba(0,0,0,0.5)',
    zIndex: 1000,
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
  },
  modalContent: {
    background: '#fff',
    borderRadius: '10px',
    padding: '1rem',
    position: 'relative',
    boxShadow: '0 0 20px #333',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
  },
  closeButton: {
    marginTop: '1rem',
    padding: '0.5rem 1.5rem',
    background: '#007bff',
    color: '#fff',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
    fontSize: '1rem',
  },
};

export default MyRoutesPage;
