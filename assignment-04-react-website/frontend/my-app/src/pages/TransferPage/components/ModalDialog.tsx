import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";
// import { ReactComponent as Success } from "!file-loader!../../assets/images/success.svg";
import { ReactComponent as Success } from "../../../assets/images/success.svg";
import { ITopUp } from "../../../interfaces/ITopUp";
import { ITransaction } from "../../../interfaces/ITransactions";

interface Props {
  showModal: boolean;
  tData: ITransaction | undefined | ITopUp;
}

export default function ModalDialog(props: Props) {
  const navigate = useNavigate();
  const ref = React.useRef<HTMLButtonElement>(null);
  useEffect(() => {
    if (props.showModal) {
      ref.current?.click();
    }
  }, [props.showModal]);

  const handleClick = (e: any) => {
    e.preventDefault();
    window.print();
  };

  return (
    <div>
      <button
        ref={ref}
        type="button"
        className="btn btn-primary"
        style={{ display: "none" }}
        data-bs-toggle="modal"
        data-bs-target="#exampleModal"
      >
        Primary
      </button>

      <div
        className="modal fade wrapper"
        id="exampleModal"
        tabIndex={-1}
        aria-labelledby="exampleModalLabel"
        aria-hidden="true"
        data-bs-backdrop="static"
        data-bs-keyboard="false"
      >
        <div className="modal-dialog" style={{ marginTop: "11.5rem" }}>
          <div className="modal-content">
            <div className="modal-header border-bottom-0">
              <div className="col">
                <div className="img__icon " style={{ textAlign: "center" }}>
                  <Success />
                </div>
                <h1
                  className="modal-title text-success fw-bold"
                  style={{ textAlign: "center" }}
                >
                  Transfer Success
                </h1>
              </div>
            </div>
            <div className="modal-body px-3">
              <table
                className="table table-borderless"
                style={{ width: "100%" }}
              >
                <tbody>
                  <tr>
                    <td
                      style={{
                        textAlign: "left",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      Amount
                    </td>
                    <td
                      style={{
                        textAlign: "right",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      {props.tData?.amount}
                    </td>
                  </tr>
                  <tr>
                    <td
                      style={{
                        textAlign: "left",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      Transaction Id
                    </td>
                    <td
                      style={{
                        textAlign: "right",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      {props.tData?.id}
                    </td>
                  </tr>
                  <tr>
                    <td
                      style={{
                        textAlign: "left",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      From
                    </td>
                    <td
                      style={{
                        textAlign: "right",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      {props.tData?.source_of_fund_id
                        ? props.tData?.source_of_fund_id
                        : props.tData?.wallet_id}
                    </td>
                  </tr>
                  <tr>
                    <td
                      style={{
                        textAlign: "left",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      To
                    </td>
                    <td
                      style={{
                        textAlign: "right",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      {props.tData?.to_wallet_id
                        ? props.tData?.to_wallet_id
                        : props.tData?.wallet_id}
                    </td>
                  </tr>
                  <tr>
                    <td
                      style={{
                        textAlign: "left",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      Description
                    </td>
                    <td
                      style={{
                        textAlign: "right",
                        fontSize: "1.5rem",
                        width: "50%",
                      }}
                    >
                      {props.tData?.description}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div className="modal-footer justify-content-center gap-3 border-top-0 mb-4">
              <button
                type="button"
                className="btn btn-outline-info fs-4"
                style={{ width: "6rem" }}
                onClick={handleClick}
              >
                Print
              </button>
              <button
                type="button"
                className="btn btn-outline-info fs-4"
                data-bs-dismiss="modal"
                style={{ width: "6rem" }}
                onClick={() => navigate("/")}
              >
                Close
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
