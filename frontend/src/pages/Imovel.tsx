import React, { useEffect, useState } from 'react';
import '../styles/Imovel.css'; // Import the CSS file for styling

const Imovel: React.FC = () => {
  const [properties, setProperties] = useState([]);

  useEffect(() => {
    const fetchProperties = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/properties');
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setProperties(data);
      } catch (error) {
        console.error('Error fetching properties:', error);
      }
    };

    fetchProperties();
  }, []);

  return (
    <div className="imovel-page">
      <h1 className="imovel-title">Available Properties</h1>
      <div className="property-list">
        {properties.length > 0 ? (
          properties.map((property) => (
            <div className="property-card" key={property.id}>
              <img
                src="https://via.placeholder.com/300" // Placeholder image
                alt={property.name}
                className="property-image"
              />
              <h2 className="property-name">{property.name}</h2>
              <button className="view-details-button">View Details</button>
            </div>
          ))
        ) : (
          <p>No properties available at the moment.</p>
        )}
      </div>
    </div>
  );
};

export default Imovel;