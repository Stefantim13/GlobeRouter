import React, { useState } from 'react';
import Cities from './Cities.json';

function RoutePlanner() {
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const [generalError, setGeneralError] = useState('');
  const [fromError, setFromError] = useState('');
  const [toError, setToError] = useState('');
  const [result, setResult] = useState(null);

  const handleSearch = () => {
    setGeneralError('');
    setFromError('');
    setToError('');
    setResult(null);

    // Validare
    if(!Cities.map(p => p.name.toLowerCase()).includes(from.toLowerCase())){
      setFromError('Orasul nu este valabil')
      return;
    }


    if(!Cities.map(p => p.name.toLowerCase()).includes(to.toLowerCase())){
      setToError('Orasul nu este valabil')
      return;
    }
      
    if (!from || !to) {
      setGeneralError('Te rugăm să completezi ambele locații.');
      return;
    }
      
    // Simulare
    const simulatedRoute = {
      path: ['București', 'Viena', 'Berlin', 'Londra'],
      price: '€150',
      duration: '12h 30m',
      changes: 2,
    };

    setResult(simulatedRoute);
  };

  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>GlobeRouter 🌍</h1>
      <p style={styles.subheading}>Planifică cea mai eficientă rută multimodală.</p>

      <div style={styles.form}>
        <input
          type="text"
          placeholder="Locație plecare (ex: București)"
          value={from}
          onChange={(e) => setFrom(e.target.value)}
          style={styles.input}
        />
        {fromError && <p style={styles.error}>{fromError}</p>}

        <input
          type="text"
          placeholder="Destinație (ex: Londra)"
          value={to}
          onChange={(e) => setTo(e.target.value)}
          style={styles.input}
        />
        {toError && <p style={styles.error}>{toError}</p>}
        
        <button onClick={handleSearch} style={styles.button}>Caută rută</button>
        {generalError && <p style={styles.error}>{generalError}</p>}
      </div>

      {result && (
        <div style={styles.result}>
          <h3>Rezultat rută:</h3>
          <p><strong>Traseu:</strong> {result.path.join(' → ')}</p>
          <p><strong>Preț:</strong> {result.price}</p>
          <p><strong>Durată:</strong> {result.duration}</p>
          <p><strong>Număr schimbări:</strong> {result.changes}</p>
        </div>
      )}
    </div>
  );
}

const styles = {
  container: {
    maxWidth: '600px',
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
    marginTop: '0.5rem',
  },
  result: {
    marginTop: '2rem',
    padding: '1rem',
    border: '1px solid #ccc',
    borderRadius: '5px',
    backgroundColor: '#fff',
  },
};

export default RoutePlanner;
