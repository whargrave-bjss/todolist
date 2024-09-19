import React, { useState } from 'react';
import { register } from '../apiservice/register';
import { login } from '../apiservice/login';

const LoginForm = ({ setIsLoggedIn, setCurrentUser }) => {
    const [isRegistering, setIsRegistering] = useState(false);
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleRegister = async (e) => {
        e.preventDefault();
        try {
            const user = await register(username, password);
            console.log('Registration successful:', user);
            setIsLoggedIn(true); 
        } catch (error) {
            console.error('Registration failed:', error);
        }
    }

    const handleLogin = async (e) => {
        e.preventDefault();
        try {
            const user = await login(username, password);
            console.log('Login successful:', user);
            setIsLoggedIn(true); 
            setCurrentUser(user);
        } catch (error) {
            console.error('Login failed:', error);
        }
    }

    return (
        <div>
            <form>
                <input 
                    type="text" 
                    placeholder="Username" 
                    value={username} 
                    onChange={(e) => setUsername(e.target.value)} 
                />
                <input 
                    type="password" 
                    placeholder="Password" 
                    value={password} 
                    onChange={(e) => setPassword(e.target.value)} 
                />
                {isRegistering ? (
                    <>
                        <button type="submit" onClick={handleRegister}>Register</button>
                        <button type="button" onClick={() => setIsRegistering(false)}>Cancel</button>
                    </>
                ) : (
                    <>
                        <button type="submit" onClick={handleLogin}>Login</button>
                        <button type="button" onClick={() => setIsRegistering(true)}>Register</button>
                    </>
                )}
            </form>
        </div>
    )
}

export default LoginForm;