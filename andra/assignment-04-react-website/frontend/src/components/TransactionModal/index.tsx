import { useEffect, useRef } from "react";
import ITransaction from "../../interfaces/transaction";
import { ReactComponent as SuccessIcon } from "../../assets/images/success-icon.svg";
import "./style.scss";
import { useNavigate } from "react-router-dom";

type Props = {
  show: boolean;
  type: string;
  data: ITransaction | undefined;
};

const TransactionModal = (props: Props) => {
  const ref = useRef<HTMLButtonElement>(null);
  const modalRef = useRef<HTMLDivElement>(null);
  const navigate = useNavigate();

  useEffect(() => {
    if (props.show) {
      ref.current?.click();
    }
  }, [props.show]);

  const handleClose = () => {
    navigate("/");
  };

  return (
    <>
      <button
        ref={ref}
        className="d-none"
        type="button"
        data-bs-toggle="modal"
        data-bs-target="#modal"
      ></button>
      <div
        className="modal fade"
        data-testid="transaction-modal"
        ref={modalRef}
        data-bs-backdrop="static"
        data-bs-keyboard="false"
        tabIndex={-1}
        id="modal"
      >
        <div className="modal-dialog modal-dialog-centered">
          <div className="modal-content">
            <div className="modal-header">
              <SuccessIcon />
              <h5 className="modal-title text-center">{props.type} Success</h5>
            </div>
            <div className="modal-body">
              <table className="table table-borderless">
                <tbody>
                  <tr>
                    <td className="align-bottom">Amount</td>
                    <td className="transaction__amount">
                      {props.data?.amount.toLocaleString("id-ID")}
                    </td>
                  </tr>
                  <tr>
                    <td>Transaction Id</td>
                    <td>{props.data?.id}</td>
                  </tr>
                  <tr>
                    <td>From</td>
                    <td>
                      {props.data?.source_of_fund_id ?? props.data?.wallet_id}
                    </td>
                  </tr>
                  <tr>
                    <td>To</td>
                    <td>{props.data?.to_wallet_id ?? props.data?.wallet_id}</td>
                  </tr>
                  <tr>
                    <td>Description</td>
                    <td>{props.data?.description}</td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div className="modal-footer">
              <button
                type="button"
                className="btn"
                onClick={() => window.print()}
              >
                Print
              </button>
              <button
                type="button"
                className="btn"
                data-bs-dismiss="modal"
                onClick={handleClose}
              >
                Close
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default TransactionModal;
