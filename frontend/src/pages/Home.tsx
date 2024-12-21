import React from 'react';
import '../styles/Home.css';
import '../pages/Imovel'
import { Link } from 'react-router-dom';

const HomePage: React.FC = () => {
  return (
    <div className="home-page">
      <header className="home-header">
        <div className="login-section">
          <button onClick={() => console.log("waiting for login")}>Login</button>
        </div>
        <div className="welcome-text">
          <h1>Imobiliária</h1>
          <p>O seu lugar para encontrar o imóvel dos sonhos!</p>
        </div>
      </header>

      <main className="home-main">
        <section className="hero-section">
          <h2>Encontre o imóvel perfeito</h2>
          <p>Explore uma vasta seleção de casas, apartamentos e terrenos disponíveis para venda ou aluguel.</p>
          <button onClick={() => alert('Procurando imóveis...')}>Começar Agora</button>
        </section>

        <section className="featured-properties">
          <h2>Imóveis em Destaque</h2>
          <div className="property-list">
            <Link to="/Imovel">
              <div className="property">
                <img src="https://via.placeholder.com/150" alt="Casa 1" />
                <h3>Casa de Luxo</h3>
                <p>3 quartos, 2 banheiros - R$ 750.000</p>
              </div>
            </Link>
            <div className="property">
              <img src="https://via.placeholder.com/150" alt="Casa 2" />
              <h3>Apartamento Moderno</h3>
              <p>2 quartos, 1 banheiro - R$ 400.000</p>
            </div>
            <div className="property">
              <img src="https://via.placeholder.com/150" alt="Casa 3" />
              <h3>Terreno Amplo</h3>
              <p>Área: 500m² - R$ 250.000</p>
            </div>
          </div>
        </section>
      </main>

      <footer className="home-footer">
        <p>&copy; 2024 Imobiliária. Todos os direitos reservados.</p>
      </footer>
    </div>
  );
};

export default HomePage;
