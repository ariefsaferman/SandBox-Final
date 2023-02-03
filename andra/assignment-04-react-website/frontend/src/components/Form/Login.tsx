import { useCookies } from "react-cookie";
import { useForm } from "react-hook-form";
import { useLocation, useNavigate } from "react-router-dom";

import "./style.scss";

type FormValues = {
  email: string;
  password: string;
};

const LoginForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>();
  const api = process.env.REACT_APP_API_URL;
  const [cookies, setCookie] = useCookies(["token"]);
  const navigate = useNavigate();
  const location = useLocation();

  let from = location.state?.from?.pathname || "/";

  const onSubmit = (data: FormValues) => {
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };

    fetch(api + "/login", requestOptions)
      .then(async (response) => {
        const resJson = await response.json();
        if (!response.ok) throw new Error(resJson.message);
        return resJson;
      })
      .then((res) => {
        if (res.data) {
          setCookie("token", res.data.token, {
            path: "/",
            expires: new Date(Date.now() + 6 * 60 * 60 * 1000),
          });
          navigate(from, { replace: true });
        }
      })
      .catch((err) => {
        console.log(err.message);
      });
  };

  return (
    <form className="form__login" onSubmit={handleSubmit(onSubmit)}>
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
          {...register("password", { required: true })}
        />
        {errors.password && <span>This field is required</span>}
      </div>
      <button type="submit" className="btn btn-lg w-100">
        Submit
      </button>
    </form>
  );
};

export default LoginForm;
