import React from "react";
import TransferForm from "../../components/Form/Transfer";
import "./style.scss";

const Transfer = () => {
  return (
    <div className="container">
      <div className="d-flex vh-90 flex-column align-items-center justify-content-center">
        <h2 className="transfer__title">Transfer</h2>
        <TransferForm />
      </div>{" "}
    </div>
  );
};

export default Transfer;
