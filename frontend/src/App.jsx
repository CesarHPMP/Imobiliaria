import React, { useState, useEffect } from 'react';
import './styles/App.css';

function App() {
  const [properties, setProperties] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/api/properties')
      .then(response => response.json())
      .then(data => setProperties(data));
  }, []);

  return (
    <div className="App">
      <h1>Real Estate CRM</h1>
      <ul>
        {properties.map(property => (
          <li key={property.id}>{property.name}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
