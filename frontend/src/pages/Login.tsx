import React, { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import "../styles/Login.css"; // Import the CSS file for styling

const Login: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(""); // Clear error message before making a new request

    try {
      const response = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        throw new Error("Login failed. Please check your credentials.");
      }

      const data = await response.json();
      console.log(data);
      const token: string = data.userToken;
      console.log(token);
      // Store token in localStorage
      localStorage.setItem("jwt", token);

      // Redirect to the homepage or dashboard
      navigate("/"); // Adjust the route if needed
    } catch (err) {
      setError(err.message);
    }
  };

  const handleInputChange = () => {
    // Clear the error message when user starts typing
    if (error) setError("");
  };

  return (
    <div className="login-page">
      <h1>Login</h1>
      <form onSubmit={handleLogin}>
        <div className="form-group">
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            onInput={handleInputChange} // Clear error when typing
            required
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password:</label>
          <input
            type="password"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            onInput={handleInputChange} // Clear error when typing
            required
          />
        </div>
        {error && <p className="error-message">{error}</p>}{" "}
        {/* Display error if exists */}
        <button type="submit">Login</button>
        <div className="signup-link">
          <p>Don't have an account?</p>
          <Link to="/register">
            <button type="button">Sign Up</button>
          </Link>
        </div>
      </form>
    </div>
  );
};

export default Login;
