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
        <div className="min-h-screen bg-gradient-to-br from-purple-400 via-pink-500 to-red-500 py-8">
        <div className="max-w-md mx-auto bg-white rounded-xl shadow-md overflow-hidden md:max-w-2xl">
          <div className="p-8">
            <h1 className="text-3xl font-bold text-center text-gray-800 mb-8">
              {isRegistering ? "Register" : "Login"}
            </h1>
            <form className="space-y-6">
              <div>
                <input
                  type="text"
                  placeholder="Username"
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                  className="w-full px-3 py-2 placeholder-gray-300 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-indigo-100 focus:border-indigo-300"
                />
              </div>
              <div>
                <input
                  type="password"
                  placeholder="Password"
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full px-3 py-2 placeholder-gray-300 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-indigo-100 focus:border-indigo-300"
                />
              </div>
              {isRegistering ? (
                <div className="flex space-x-4">
                  <button
                    type="submit"
                    onClick={handleRegister}
                    className="w-full px-3 py-4 text-white bg-teal-500 rounded-md focus:bg-teal-600 focus:outline-none"
                  >
                    Register
                  </button>
                  <button
                    type="button"
                    onClick={() => setIsRegistering(false)}
                    className="w-full px-3 py-4 text-white bg-red-500 rounded-md focus:bg-red-600 focus:outline-none"
                  >
                    Cancel
                  </button>
                </div>
              ) : (
                <div className="flex space-x-4">
                  <button
                    type="submit"
                    onClick={handleLogin}
                    className="w-full px-3 py-4 text-white bg-teal-500 rounded-md focus:bg-teal-600 focus:outline-none"
                  >
                    Login
                  </button>
                  <button
                    type="button"
                    onClick={() => setIsRegistering(true)}
                    className="w-full px-3 py-4 text-white bg-purple-500 rounded-md focus:bg-purple-600 focus:outline-none"
                  >
                    Register
                  </button>
                </div>
              )}
            </form>
          </div>
        </div>
      </div>
    )
}

export default LoginForm;