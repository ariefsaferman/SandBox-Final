import React from "react";
import LoginPage from "./pages/LoginPage/index";
import { Route, Routes } from "react-router-dom";
import "./App.css";
import "https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js";
import "bootstrap/dist/css/bootstrap.min.css";
import RegisterPage from "./pages/RegisterPage";
import HomePage from "./pages/HomePage";
import ProtectedPage, { ProtectReversePage } from "./pages/ProtectedPage";
import TransferPage from "./pages/TransferPage";
import TopupPage from "./pages/TopupPage";
import GamesPage from "./pages/GamesPage";

import "react-toastify/dist/ReactToastify.css";

function App() {
  return (
    <div className="App">
      <Routes>
        <Route element={<ProtectedPage />}>
          <Route path="/" element={<HomePage />}></Route>
          <Route path="/transfer" element={<TransferPage />}></Route>
          <Route path="/topup" element={<TopupPage />}></Route>
          <Route path="/games" element={<GamesPage />}></Route>
        </Route>
        <Route element={<ProtectReversePage />}>
          <Route path="/login" element={<LoginPage />}></Route>
          <Route path="/register" element={<RegisterPage />}></Route>
        </Route>
        <Route path="*" element={<h2> not found page</h2>}></Route>
      </Routes>
    </div>
  );
}

export default App;
