import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

function LoginPage({ setIsLoggedIn, setUserEmail }) {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');
    const [isLoggedIn, setLocalIsLoggedIn] = useState(false);
    const navigate = useNavigate();

    useEffect(() => {
        if (localStorage.getItem('isLoggedIn') === 'true') {
            navigate('/');
        }
    }, [navigate]);

    const handleLogin = async (e) => {
        e.preventDefault();
        setError('');
        if (!email.includes('@')) {
            setError('Adresa de email nu este validă.');
            return;
        }

        if (!email || !password) {
            setError('Te rugăm să completezi toate câmpurile.');
            return;
        }

        if (email === 'user@example.com' && password === 'parola123') {
            if (setIsLoggedIn) setIsLoggedIn(true);
            if (setUserEmail) setUserEmail(email);
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('email', email);
            return;
        }

        try {
            const res = await fetch(
                'http://localhost:8080/login',
                {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ email, password, confirmationPassword: password }),
                }
            );
            const data = await res.json();
            if (data.error) {
                setError(data.error);
            } else {
                if (setIsLoggedIn) setIsLoggedIn(true);
                if (setUserEmail) setUserEmail(email);
                localStorage.setItem('isLoggedIn', 'true');
                localStorage.setItem('email', email);
                navigate('/');
            }
        } catch (err) {
            setError('Eroare la conectarea cu serverul.');
        }
    };

    if (isLoggedIn) {
        return (
            <div style={styles.container}>
                <h2>Bine ai venit, {email}!</h2>
                <p>Ai fost autentificat cu succes!</p>
            </div>
        );
    }

    return (
        <div style={styles.container}>
            <h1 style={styles.heading}>Autentificare</h1>
            <form style={styles.form}>
                <input
                    type="email"
                    placeholder="Email"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    style={styles.input}
                />
                <input
                    type="password"
                    placeholder="Parolă"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    style={styles.input}
                />
                {error && <p style={styles.error}>{error}</p>}
                <button onClick={handleLogin} type="submit" style={styles.button}>
                    Login
                </button>
            </form>
        </div>
    );
}

const styles = {
    container: {
        maxWidth: '400px',
        margin: '4rem auto',
        padding: '2rem',
        fontFamily: 'Arial, sans-serif',
        backgroundColor: '#f0f0f0',
        borderRadius: '10px',
        boxShadow: '0 0 10px rgba(0,0,0,0.1)',
        textAlign: 'center',
    },
    heading: {
        marginBottom: '1.5rem',
        fontSize: '1.5rem',
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
        fontSize: '0.9rem',
    },
};

export default LoginPage;
