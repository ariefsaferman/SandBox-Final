import React from "react";
import { NavLink, useNavigate } from "react-router-dom";
import ImageRegister from "../../assets/images/register.png";
import { useForm } from "react-hook-form";
import "./index.scss";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

interface RegisterRequest {
  name: string;
  email: string;
  password: string;
}

interface RegisterReqs {
  email: string;
  password: string;
  first_name: string;
  last_name: string;
}

export default function RegisterPage() {
  const API_URL = process.env.REACT_APP_API_URL;
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterRequest>();

  const onSubmit = (data: RegisterRequest) => {
    const arr = data.name.split(" ");
    const registerRequest: RegisterReqs = {
      email: data.email,
      password: data.password,
      first_name: arr[0],
      last_name: arr.length > 1 ? arr[1] : arr[0],
    };

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(registerRequest),
    };

    fetch(API_URL + "/register", requestOptions)
      .then((props) => {
        if (props.ok) {
          return props.json();
        }

        if (props.status === 400) {
          throw new Error("Bad Request");
        } else if (props.status === 409) {
          throw new Error("Email already been registered");
        } else {
          throw new Error("Internal Server Error");
        }
      })
      .then((data) => {
        if (data.error) {
        }
        console.log(data);
        navigate("/login");
      })
      .catch((error) => {
        toast.error(error.toString(), {
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

      <section className="register__form">
        <div className="container">
          <div className="row vh-100">
            <div className="col-lg-6 d-flex flex-column justify-content-center">
              <div className="login__form__title text-start">
                <h1>Register</h1>
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
                      Name
                    </label>
                    <input
                      type="text"
                      className="form-control"
                      id="exampleInputEmail1"
                      aria-describedby="emailHelp"
                      style={{
                        height: "3rem",
                        borderColor: errors.name ? "red" : "",
                      }}
                      placeholder="Asep Budiantoro Chandradiman"
                      {...register("name", { required: true })}
                    />
                    {errors.name && (
                      <span className="text-danger">required field</span>
                    )}
                  </div>
                  <div className="mb-3 text-start">
                    <label
                      htmlFor="exampleInputEmail2"
                      className="form-label login__text fs-4"
                    >
                      Email
                    </label>
                    <input
                      type="email"
                      className="form-control"
                      id="exampleInputEmail2"
                      aria-describedby="emailHelp"
                      style={{
                        height: "3rem",
                        borderColor: errors.email ? "red" : "",
                      }}
                      placeholder="asep.bc@gmail.com"
                      {...register("email", { required: true })}
                    />
                    {errors.email && (
                      <span className="text-danger">required field</span>
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
                      {...register("password", {
                        required: true,
                        minLength: {
                          value: 8,
                          message: "Password must have at least 8 characters",
                        },
                      })}
                    />
                    {errors.password?.type === "required" ? (
                      <span className="text-danger">required field</span>
                    ) : (
                      <span className="text-danger">
                        {errors.password?.message}
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
                    src={ImageRegister}
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
