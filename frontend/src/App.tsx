import React from "react";
import { Routes, Route } from "react-router-dom";
import HomePage from "./pages/Home";
import Login from "./pages/Login";
import Imovel from "./pages/Imovel";
import Register from "./pages/Register";

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePage />} />
      <Route path="/Login" element={<Login />} />
      <Route path="/Imovel" element={<Imovel />} />
      <Route path="/Register" element={<Register />} />
    </Routes>
  );
}

export default App;
