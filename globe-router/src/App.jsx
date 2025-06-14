import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom';

import RoutePlanner from './RoutePlanner';
import LoginPage from './LoginPage';
import SignUp from './SignUp';
import MyRoutesPage from './MyRoutes';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(() => {
    return localStorage.getItem('isLoggedIn') === 'true';
  });

  React.useEffect(() => {
    localStorage.setItem('isLoggedIn', isLoggedIn ? 'true' : 'false');
  }, [isLoggedIn]);
  const [userEmail, setUserEmail] = useState(() => localStorage.getItem('email') || '');

  const handleLogout = () => {
    setIsLoggedIn(false);
    setUserEmail('');
    localStorage.removeItem('isLoggedIn');
    localStorage.removeItem('email');
  };

  return (
    <Router>
      <div>
        <nav style={styles.navbar}>
          <div style={styles.leftNav}>
            <Link to="/" style={styles.routePlannerButton}>Route Planner</Link>
          </div>
          <div style={styles.rightNav}>
            {!isLoggedIn ? (
              <>
                <Link to="/login" style={styles.navButton}>Login</Link>
                <Link to="/signup" style={styles.navButton}>Register</Link>
              </>
            ) : (
              <>
                <Link to="/my-routes" style={styles.navButton}>My Routes</Link>
                <button onClick={handleLogout} style={styles.navButton}>Logout</button>
              </>
            )}
          </div>
        </nav>

        <Routes>
          <Route path="/" element={<RoutePlanner />} />
          <Route path="/login" element={<LoginPage setIsLoggedIn={setIsLoggedIn} setUserEmail={setUserEmail} />} />
          <Route path="/signup" element={<SignUp setIsLoggedIn={setIsLoggedIn} setUserEmail={setUserEmail} />} />
          <Route
            path="/my-routes"
            element={isLoggedIn ? <MyRoutesPage email={userEmail} /> : <Navigate to="/login" />}
          />
        </Routes>
      </div>
    </Router>
  );
}

const styles = {
  navbar: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: '1rem 2rem',
    backgroundColor: '#eee',
  },
  leftNav: {
    display: 'flex',
    alignItems: 'center',
  },
  rightNav: {
    display: 'flex',
    alignItems: 'center',
    gap: '1rem',
  },
  routePlannerButton: {
    textDecoration: 'none',
    fontWeight: 'bold',
    fontSize: '1.1rem',
    color: '#007bff',
  },
  navButton: {
    textDecoration: 'none',
    padding: '0.7rem 1rem',
    backgroundColor: '#007bff',
    color: 'white',
    borderRadius: '5px',
    border: 'none',
    cursor: 'pointer',
    font: 'inherit',
    lineHeight: '1',
    display: 'inline-block',
  },
};

export default App;
