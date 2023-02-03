import React from "react";
import { useCookies } from "react-cookie";
import { NavLink, useNavigate } from "react-router-dom";
import "./style.scss";

const Navbar = () => {
  const [cookies, setCookies, removeCookies] = useCookies(["token"]);
  const navigate = useNavigate();

  const handleLogout = () => {
    if (window.confirm("Are you sure you want to logout?")) {
      removeCookies("token", { path: "/" });
      navigate("/login");
    }
  };

  return (
    <nav
      className={`navbar navbar-expand-lg bg-transparent ${
        cookies.token ? "" : "fixed-top"
      }`}
    >
      <div className="container">
        <h1 className="navbar-brand m-0">DigiWallet</h1>
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
        <div
          className="collapse navbar-collapse justify-content-end"
          id="navbarSupportedContent"
        >
          <ul className="navbar-nav">
            <li className="nav-item">
              <NavLink to={"/"} className="nav-link">
                Home
              </NavLink>
            </li>
            {!cookies.token ? (
              <>
                <li className="nav-item ms-lg-5">
                  <NavLink to={"/login"} className="nav-link text-white">
                    Login
                  </NavLink>
                </li>
                <li className="nav-item ms-lg-5">
                  <NavLink to={"/register"} className="nav-link text-white">
                    Register
                  </NavLink>
                </li>
              </>
            ) : (
              <>
                <li className="nav-item ms-lg-5">
                  <NavLink to={"/transfer"} className="nav-link">
                    Transfer
                  </NavLink>
                </li>
                <li className="nav-item ms-lg-5">
                  <NavLink to={"/topup"} className="nav-link">
                    Topup
                  </NavLink>
                </li>
                <li className="nav-item ms-lg-5">
                  <NavLink to={"/games"} className="nav-link">
                    Games
                  </NavLink>
                </li>
                <li className="nav-item ms-lg-5">
                  <button className="btn nav-link" onClick={handleLogout}>
                    Logout
                  </button>
                </li>
              </>
            )}
          </ul>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
