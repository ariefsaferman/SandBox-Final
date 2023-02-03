import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { toast, ToastContainer } from "react-toastify";

import "./style.scss";

type FormValues = {
  name: string;
  email: string;
  password: string;
};

type Payload = {
  email: string;
  password: string;
  first_name: string;
  last_name: string;
};

const RegisterForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>();
  const api = process.env.REACT_APP_API_URL;
  const navigate = useNavigate();

  const onSubmit = (data: FormValues) => {
    const splitName = data.name.split(" ");
    const payload: Payload = {
      email: data.email,
      password: data.password,
      first_name: splitName[0],
      last_name: splitName.length > 1 ? splitName[1] : splitName[0],
    };

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(payload),
    };

    fetch(api + "/register", requestOptions)
      .then(async (response) => {
        const resJson = await response.json();
        if (!response.ok) throw new Error(resJson.message);
        navigate("/login");
      })
      .catch((err) => {
        toast.error(err.message);
      });
  };

  return (
    <form className="form__register" onSubmit={handleSubmit(onSubmit)}>
      <ToastContainer />
      <div className="form-group">
        <label htmlFor="name" className="fw-bold">
          Name
        </label>
        <input
          type="text"
          className="form-control form-control-lg"
          id="name"
          placeholder="Asep Budiantoro Chandradiman"
          {...register("name", { required: true })}
        />
        {errors.name && <span>This field is required</span>}
      </div>
      <div className="form-group">
        <label htmlFor="email" className="fw-bold">
          Email
        </label>
        <input
          type="email"
          className="form-control form-control-lg"
          id="email"
          placeholder="asep.bc@gmail.com"
          {...register("email", { required: true })}
        />
        {errors.email && <span>This field is required</span>}
      </div>
      <div className="form-group">
        <label htmlFor="password" className="fw-bold">
          Password
        </label>
        <input
          type="password"
          className="form-control form-control-lg"
          id="password"
          placeholder="********"
          {...register("password", {
            required: {
              value: true,
              message: "Password is required",
            },
            minLength: {
              value: 8,
              message: "Password must be at least 8 characters",
            },
          })}
        />
        {errors.password && <span>{errors.password.message}</span>}
      </div>
      <button type="submit" className="btn btn-lg w-100">
        Submit
      </button>
    </form>
  );
};

export default RegisterForm;
