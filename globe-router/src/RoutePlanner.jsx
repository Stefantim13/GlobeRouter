import React, { useState } from 'react';
import Cities from './Cities.json';

function RoutePlanner() {
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const [travelDate, setTravelDate] = useState('');
  const [fromSuggestions, setFromSuggestions] = useState([]);
  const [toSuggestions, setToSuggestions] = useState([]);
  const [fromError, setFromError] = useState('');
  const [toError, setToError] = useState('');
  const [dateError, setDateError] = useState('');
  const [generalError, setGeneralError] = useState('');
  const [results, setResults] = useState([]);

  const [sortOption, setSortOption] = useState('');
  const [savedRoutes, setSavedRoutes] = useState([]);

  const getTodayISO = () => {
    const today = new Date();
    return today.toISOString().split('T')[0];
  };

  const formatDate = (isoDate) => {
    const [year, month, day] = isoDate.split('-');
    return `${day} ${month} ${year}`;
  };

  const getSuggestions = (value) => {
    const input = value.toLowerCase();
    return Cities.filter(city =>
      city.name.toLowerCase().startsWith(input)
    ).slice(0, 5);
  };

  const handleSearch = () => {
    setFromError('');
    setToError('');
    setDateError('');
    setGeneralError('');
    setResults([]);

    if (!from || !to) {
      setGeneralError('Te rugăm să completezi ambele locații.');
      return;
    }

    if(from === to){
      setGeneralError('Nu ai cum sa calatoresti in aceeasi locatie');
      return;
    }

    const cityNames = Cities.map(c => c.name.toLowerCase());

    if (!cityNames.includes(from.toLowerCase())) {
      setFromError('Orașul de plecare nu este valid.');
      return;
    }

    if (!cityNames.includes(to.toLowerCase())) {
      setToError('Orașul de destinație nu este valid.');
      return;
    }

    if (!travelDate) {
      setDateError('Te rugăm să selectezi o dată pentru călătorie.');
      return;
    }

    let simulatedRoutes = [
      { path: [from, 'Viena', to], price: 120, duration: 10, changes: 1 },
      { path: [from, 'Berlin', 'Amsterdam', to], price: 200, duration: 7, changes: 2 },
      { path: [from, 'Budapesta', 'Paris', to], price: 180, duration: 13, changes: 3 },
      { path: [from, 'Praga', to], price: 90, duration: 9, changes: 1 },
    ];

    switch (sortOption) {
      case 'price-asc':
        simulatedRoutes.sort((a, b) => a.price - b.price);
        break;
      case 'duration-asc':
        simulatedRoutes.sort((a, b) => a.duration - b.duration);
        break;
      default:
        break;
    }

    simulatedRoutes = simulatedRoutes.map(route => ({
      ...route,
      date: formatDate(travelDate)
    }));

    setResults(simulatedRoutes);
  };

  const toggleBookmark = (routeIndex) => {
    setSavedRoutes((prev) => {
      if (prev.includes(routeIndex)) {
        return prev.filter((idx) => idx !== routeIndex);
      } else {
        return [...prev, routeIndex];
      }
    });
  };

  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>GlobeRouter 🌍</h1>
      <p style={styles.subheading}>Planifică cea mai eficientă rută multimodală.</p>

      <div style={styles.form}>
        <input
          type="text"
          placeholder="Locație plecare"
          value={from}
          onChange={(e) => {
            setFrom(e.target.value);
            setFromSuggestions(getSuggestions(e.target.value));
          }}
          style={styles.input}
        />
        {fromSuggestions.length > 0 && (
          <ul style={styles.suggestions}>
            {fromSuggestions.map((city, idx) => (
              <li
                key={idx}
                style={styles.suggestionItem}
                onClick={() => {
                  setFrom(city.name);
                  setFromSuggestions([]);
                }}
              >
                {city.name}
              </li>
            ))}
          </ul>
        )}
        {fromError && <p style={styles.error}>{fromError}</p>}

        <input
          type="text"
          placeholder="Destinație"
          value={to}
          onChange={(e) => {
            setTo(e.target.value);
            setToSuggestions(getSuggestions(e.target.value));
          }}
          style={styles.input}
        />
        {toSuggestions.length > 0 && (
          <ul style={styles.suggestions}>
            {toSuggestions.map((city, idx) => (
              <li
                key={idx}
                style={styles.suggestionItem}
                onClick={() => {
                  setTo(city.name);
                  setToSuggestions([]);
                }}
              >
                {city.name}
              </li>
            ))}
          </ul>
        )}
        {toError && <p style={styles.error}>{toError}</p>}
        <input 
          type="date"
          value={travelDate}
          min={getTodayISO()}
          onChange={(e) => setTravelDate(e.target.value)}
          style={styles.input}
        />
        {dateError && <p style={styles.error}>{dateError}</p>}

        <select value={sortOption} onChange={(e) => setSortOption(e.target.value)} style={styles.input}>
          <option value="">Sortează după...</option>
          <option value="price-asc">Preț crescător</option>
          <option value="duration-asc">Durată crescătoare</option>
        </select>

        <button onClick={handleSearch} style={styles.button}>Caută rută</button>
        {generalError && <p style={styles.error}>{generalError}</p>}
      </div>

      {results.length > 0 && (
        <div style={styles.result}>
          <h3>Rezultate rute:</h3>
          <hr style={{ margin: '1rem 0' }} />
          {results.map((route, index) => (
            <div key={index} style={{ marginBottom: '1rem', position: 'relative' }}>
              <button
                onClick={() => toggleBookmark(index)}
                style={{
                  ...styles.saveButton,
                  backgroundColor: savedRoutes.includes(index) ? '#ffd700' : '#eee',
                  color: savedRoutes.includes(index) ? '#333' : '#888',
                }}
                aria-label={savedRoutes.includes(index) ? 'Anulează salvarea' : 'Salvează ruta'}
              >
                {savedRoutes.includes(index) ? '★' : '☆'}
              </button>
              <p style={{paddingRight: '2.5rem'}}><strong>Traseu:</strong> {route.path.join(' → ')}</p>
              <p><strong>Preț:</strong> €{route.price}</p>
              <p><strong>Durată:</strong> {route.duration} ore</p>
              <p><strong>Schimbări:</strong> {route.changes}</p>
              <p><strong>Data:</strong> {route.date}</p>
              <hr />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles = {
  container: {
    maxWidth: '650px',
    margin: '2rem auto',
    padding: '2rem',
    fontFamily: 'Arial, sans-serif',
    backgroundColor: '#f9f9f9',
    borderRadius: '10px',
    boxShadow: '0 0 10px #ddd',
  },
  heading: {
    fontSize: '2rem',
    marginBottom: '0.5rem',
  },
  subheading: {
    fontSize: '1rem',
    marginBottom: '1.5rem',
    color: '#666',
  },
  form: {
    display: 'flex',
    flexDirection: 'column',
    gap: '1rem',
  },
  input: {
    padding: '0.75rem',
    fontSize: '1rem',
    border: '1px solid #ccc',
    borderRadius: '5px',
  },
  
  suggestions: {
    listStyle: 'none',
    padding: 0,
    margin: 0,
    backgroundColor: '#fff',
    border: '1px solid #ccc',
    borderRadius: '5px',
    maxHeight: '150px',
    overflowY: 'auto',
  },
  suggestionItem: {
    padding: '0.5rem',
    cursor: 'pointer',
  },
  button: {
    padding: '0.75rem',
    fontSize: '1rem',
    backgroundColor: '#007bff',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
  error: {
    color: 'red',
    marginTop: '-0.5rem',
    marginBottom: '0.5rem',
  },
  result: {
    marginTop: '2rem',
    padding: '1rem',
    border: '1px solid #ccc',
    borderRadius: '5px',
    backgroundColor: '#fff',
  },
  saveButton: {
    position: 'absolute',
    top: '1rem',
    right: '1rem',
    border: 'none',
    borderRadius: '4px',
    padding: '0.5rem 1rem',
    cursor: 'pointer',
    fontSize: '1.3rem',
    backgroundColor: '#eee', 
    color: '#888', 
    zIndex: 2,
  },
};

export default RoutePlanner;
