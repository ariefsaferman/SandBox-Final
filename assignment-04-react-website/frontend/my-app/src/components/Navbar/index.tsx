import React from "react";
import { Cookies, useCookies } from "react-cookie";
import { NavLink, useNavigate } from "react-router-dom";

export default function Navbar() {
  const navigate = useNavigate();
  const [cookies, setCookie] = useCookies(["token"]);

  const logout = () => {
    if (window.confirm("Are you sure you want to logout?")) {
      setCookie("token", "");
      navigate("/login");
    }
    return;
  };
  return (
    <section className="home__page">
      <div className="navbar navbar-expand-lg">
        <div className="container">
          <div className="navbar-brand fw-bold">
            <h3>DigiWallet</h3>
          </div>
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
            <ul className="navbar-nav ">
              <li className="nav-item ms-5">
                <NavLink to={"/"} className="nav-link fs-5">
                  Home
                </NavLink>
              </li>
              <li className="nav-item ms-5 ">
                <NavLink to={"/transfer"} className="nav-link fs-5">
                  Transfer
                </NavLink>
              </li>
              <li className="nav-item ms-5">
                <NavLink to={"/topup"} className="nav-link fs-5">
                  Topup
                </NavLink>
              </li>
              <li className="nav-item ms-5">
                <NavLink to={"/games"} className="nav-link fs-5">
                  Games
                </NavLink>
              </li>
              <li className="nav-item ms-5">
                <NavLink
                  to={"/login"}
                  className="nav-link fs-5"
                  onClick={logout}
                >
                  Logout
                </NavLink>
              </li>
            </ul>
          </div>
        </div>
      </div>
    </section>
  );
}
