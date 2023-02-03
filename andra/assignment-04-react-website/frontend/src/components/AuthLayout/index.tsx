import React from "react";

import "./style.scss";

type Props = {
  children: React.ReactNode;
  title: string;
  imgSrc: string;
};

const AuthLayout = (props: Props) => {
  return (
    <div className="container">
      <div className="row vh-100">
        <div className="col-lg-6 d-flex flex-column justify-content-center">
          <h2 className="title">{props.title}</h2>
          {props.children}
        </div>
        <div className="col-lg-6 auth__col">
          <div className="auth__hero">
            <img src={props.imgSrc} alt="Auth" />
          </div>
        </div>
      </div>
    </div>
  );
};

export default AuthLayout;
