import React from "react";
import AuthLayout from "../../components/AuthLayout";

import RegisterImg from "../../assets/images/undraw_authentication_re_svpt 1.png";
import RegisterForm from "../../components/Form/Register";

const Register = () => {
  return (
    <AuthLayout title="Register" imgSrc={RegisterImg}>
      <RegisterForm />
    </AuthLayout>
  );
};

export default Register;
