import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import "../styles/Register.css";

const Register: React.FC = () => {
  const [formData, setFormData] = useState({
    nome: "",
    email: "",
    endereco: "",
    numero: "",
    senha: "",
  });
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError(null); // Reset error state

    try {
      const response = await fetch("http://localhost:8080/api/addUsers", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (!response.ok) {
        throw new Error("Registration failed. Please check the input data.");
      }

      // Redirect to login page upon success
      navigate("/login");
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "An unexpected error occurred."
      );
    }
  };

  return (
    <div className="container mt-5">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <div className="card shadow-sm">
            <div className="card-header bg-primary text-white">
              <h3 className="text-center">Register</h3>
            </div>
            <div className="card-body">
              {error && <div className="alert alert-danger">{error}</div>}
              <form onSubmit={handleSubmit}>
                <div className="form-group mb-3">
                  <label htmlFor="nome">Name</label>
                  <input
                    type="text"
                    className="form-control"
                    id="nome"
                    name="nome"
                    value={formData.nome}
                    onChange={handleChange}
                    placeholder="Enter your name"
                    required
                  />
                </div>
                <div className="form-group mb-3">
                  <label htmlFor="email">Email</label>
                  <input
                    type="email"
                    className="form-control"
                    id="email"
                    name="email"
                    value={formData.email}
                    onChange={handleChange}
                    placeholder="Enter your email"
                    required
                  />
                </div>
                <div className="form-group mb-3">
                  <label htmlFor="endereco">Address</label>
                  <input
                    type="text"
                    className="form-control"
                    id="endereco"
                    name="endereco"
                    value={formData.endereco}
                    onChange={handleChange}
                    placeholder="Enter your address"
                  />
                </div>
                <div className="form-group mb-3">
                  <label htmlFor="numero">Phone Number</label>
                  <input
                    type="text"
                    className="form-control"
                    id="numero"
                    name="numero"
                    value={formData.numero}
                    onChange={handleChange}
                    placeholder="Enter your phone number"
                  />
                </div>
                <div className="form-group mb-3">
                  <label htmlFor="senha">Password</label>
                  <input
                    type="password"
                    className="form-control"
                    id="senha"
                    name="senha"
                    value={formData.senha}
                    onChange={handleChange}
                    placeholder="Enter your password"
                    required
                  />
                </div>
                <button type="submit" className="btn btn-primary w-100">
                  Register
                </button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Register;
