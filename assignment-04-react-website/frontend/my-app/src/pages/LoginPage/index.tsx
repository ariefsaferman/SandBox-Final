import React from "react";
import { NavLink, useNavigate } from "react-router-dom";
import "./index.scss";
import ImageLogin from "../../assets/images/login.png";
import { useForm } from "react-hook-form";
import { useCookies } from "react-cookie";
import { toast, ToastContainer } from "react-toastify";

interface LoginRequest {
  email: string;
  password: string;
}

export default function LoginPage() {
  const API_URL = process.env.REACT_APP_API_URL;
  const navigate = useNavigate();
  const [cookies, setCookie] = useCookies(["token"]);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginRequest>();

  const onSubmit = (data: LoginRequest) => {
    const loginRequest: LoginRequest = {
      email: data.email,
      password: data.password,
    };

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(loginRequest),
    };

    fetch(API_URL + "/login", requestOptions)
      .then((props) => {
        if (props.ok) {
          return props.json();
        }

        if (props.status === 400) {
          throw new Error("Bad Request");
        } else if (props.status === 401) {
          throw new Error("Wrong email or password");
        } else {
          throw new Error("Internal Server Error");
        }
      })
      .then((res) => {
        if (res.data) {
          setCookie("token", res.data.token, {
            path: "/",
            expires: new Date(Date.now() + 6 * 60 * 60 * 1000),
          });
        }
        navigate("/");
      })
      .catch((err) => {
        console.log(err);
        toast.error(err.message, {
          position: "top-right",
          autoClose: 5000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: "colored",
        });
      });
  };

  return (
    <div>
      <ToastContainer
        position="top-right"
        autoClose={5000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="colored"
      />
      <section className="navbar__login">
        <div className="navbar navbar-expand-lg bg-transparent fixed-top">
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
              <ul className="navbar-nav">
                <li className="nav-item ms-5">
                  <NavLink to={"/"} className="nav-link black fs-5">
                    Home
                  </NavLink>
                </li>
                <li className="nav-item ms-5">
                  <NavLink to={"/login"} className="nav-link text-white fs-5">
                    Login
                  </NavLink>
                </li>
                <li className="nav-item ms-5">
                  <NavLink
                    to={"/register"}
                    className="nav-link text-white fs-5"
                  >
                    Register
                  </NavLink>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      <section className="login__form">
        <div className="container">
          <div className="row vh-100">
            <div className="col-lg-6 d-flex flex-column justify-content-center">
              <div className="login__form__title text-start">
                <h1>Login</h1>
              </div>
              <div className="login__form__input mt-3">
                <form
                  className="border__control"
                  onSubmit={handleSubmit(onSubmit)}
                >
                  <div className="mb-3 text-start">
                    <label
                      htmlFor="exampleInputEmail1"
                      className="form-label login__text fs-4"
                    >
                      Email
                    </label>
                    <input
                      type="email"
                      className="form-control"
                      id="exampleInputEmail1"
                      aria-describedby="emailHelp"
                      style={{
                        height: "3rem",
                        borderColor: errors.email ? "red" : "",
                      }}
                      placeholder="asep@bc@gmail.com"
                      {...register("email", { required: true })}
                    />
                    {errors.email && (
                      <span className="text-danger">
                        This field is required
                      </span>
                    )}
                  </div>
                  <div className="mb-3 text-start">
                    <label
                      htmlFor="exampleInputPassword1"
                      className="form-label login__text fs-4"
                    >
                      Password
                    </label>
                    <input
                      type="password"
                      className="form-control"
                      id="exampleInputPassword1"
                      style={{
                        height: "3rem",
                        borderColor: errors.password ? "red" : "",
                      }}
                      placeholder="********"
                      {...register("password", { required: true })}
                    />
                    {errors.password && (
                      <span className="text-danger">
                        This field is required
                      </span>
                    )}
                  </div>
                  <button
                    type="submit"
                    className="btn btn-info mt-3 btn__submit text-white"
                  >
                    Submit
                  </button>
                </form>
              </div>
            </div>
            <div className="col-lg-6">
              <div className="bg__right">
                <div className="text-center">
                  <img
                    src={ImageLogin}
                    className="login__images"
                    alt="login images"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}
