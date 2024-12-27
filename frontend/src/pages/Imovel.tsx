import React, { useState, useEffect } from "react";
import "../styles/Imovel.css";

const Imovel: React.FC = () => {
  const [properties, setProperties] = useState([]);
  const [loading, setLoading] = useState(true); // Add loading state
  const [error, setError] = useState<string | null>(null); // Add error state

  useEffect(() => {
    const fetchProperties = async () => {
      const token = localStorage.getItem("jwt");
      if (!token) {
        setError("User not authenticated");
        setLoading(false);
        return;
      }

      try {
        const response = await fetch(
          "http://localhost:8080/api/protected/properties",
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
              Authorization: "Bearer " + localStorage.getItem("jwt"),
            },
          }
        );

        if (!response.ok) {
          throw new Error("Network response was not ok");
        }

        const data = await response.json();
        setProperties(data);
      } catch (error) {
        console.error("Error fetching properties:", error);
        setError("Failed to fetch properties");
      } finally {
        setLoading(false); // Ensure loading state is updated after the request
      }
    };

    fetchProperties();
  }, []);

  if (loading) return <p>Loading properties...</p>;
  if (error) return <p>{error}</p>;

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
