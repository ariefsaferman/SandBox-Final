import { Route, Routes } from "react-router-dom";

import MainLayout from "./components/MainLayout";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";

import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";
import "react-toastify/dist/ReactToastify.css";
import "@fontsource/montserrat";
import "bootstrap/dist/js/bootstrap.min.js";
import Topup from "./pages/Topup";
import Transfer from "./pages/Transfer";
import ProtectedPage from "./pages/ProtectedPage";
import UnauthenticatedPage from "./pages/UnauthenticatedPage";
import Games from "./pages/Games";

function App() {
  return (
    <div className="App">
      <Routes>
        <Route element={<MainLayout />}>
          <Route element={<UnauthenticatedPage />}>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
          </Route>
          <Route path="/" element={<ProtectedPage />}>
            <Route index element={<Home />} />
            <Route path="/topup" element={<Topup />} />
            <Route path="/transfer" element={<Transfer />} />
            <Route path="/games" element={<Games />} />
          </Route>
        </Route>
        <Route path="*" element={<h2>404 Not Found</h2>} />
      </Routes>
    </div>
  );
}

export default App;
