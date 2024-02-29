import React, { useState } from 'react';
import axios from 'axios';
import Dashboard from './Dashboard';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [token, setToken] = useState(null);

  const handleLogin = async () => {
    try {
      const response = await axios.post('http://localhost:8080/login', {
        username: username,
        password: password,
      });
      const { token } = response.data;
      setToken(token);
      console.log(token)
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  if (token) {
    return <Dashboard token={token} />;
  }

  return (
    <div>
      <h2>Login</h2>
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
      <button onClick={handleLogin}>Login</button>
    </div>
  );
};

export default LoginPage;
