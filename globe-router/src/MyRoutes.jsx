import React from 'react';

const routes = [
  {
    id: 1,
    from: 'Timișoara',
    to: 'București',
    date: '2025-07-01',
    airline: 'TAROM',
    flightNumber: 'RO123',
  },
  {
    id: 2,
    from: 'Cluj-Napoca',
    to: 'Londra',
    date: '2025-07-10',
    airline: 'Wizz Air',
    flightNumber: 'W6789',
  },
  {
    id: 3,
    from: 'Iași',
    to: 'Paris',
    date: '2025-08-15',
    airline: 'Air France',
    flightNumber: 'AF234',
  },
];

function MyRoutesPage() {
  return (
    <div style={styles.container}>
      <h1 style={styles.heading}>Rutele mele</h1>
      {routes.length === 0 ? (
        <p style={styles.noRoutes}>Nu ai rute disponibile.</p>
      ) : (
        <ul style={styles.list}>
          {routes.map((route) => (
            <li key={route.id} style={styles.card}>
              <p><strong>{route.from}</strong> → <strong>{route.to}</strong></p>
              <p>Data: {route.date}</p>
              <p>Companie: {route.airline}</p>
              <p>Zbor: {route.flightNumber}</p>
            </li>
          ))}
        </ul>
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
  },
  noRoutes: {
    textAlign: 'center',
    color: '#777',
  },
};

export default MyRoutesPage;
