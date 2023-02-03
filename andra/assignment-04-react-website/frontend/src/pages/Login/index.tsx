import React from "react";
import AuthLayout from "../../components/AuthLayout";

import LoginImg from "../../assets/images/undraw_login_re_4vu2 1.png";
import LoginForm from "../../components/Form/Login";

const Login = () => {
  return (
    <AuthLayout title="Login" imgSrc={LoginImg}>
      <LoginForm />
    </AuthLayout>
  );
};

export default Login;
