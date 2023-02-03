import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import NavbarLogin from "../../components/Navbar";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../store";
import { UserDispatch, fetchUser } from "../../store/slices/userSlice";

import "../style.scss";
import "./index.scss";
import { useCookies } from "react-cookie";
import { toast, ToastContainer } from "react-toastify";
import ModalDialog from "./components/ModalDialog";
import { ITransaction } from "../../interfaces/ITransactions";

interface TransferRequest {
  amount: number;
  to: number;
  description: string;
}

export default function TransferPage() {
  const API_URL = process.env.REACT_APP_API_URL;
  const { user } = useSelector((state: RootState) => state.user);
  const dispatch: UserDispatch = useDispatch();
  const [modal, setModal] = useState<boolean>(false);
  const [Tdata, setTdata] = useState<ITransaction>();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<TransferRequest>();

  const [cookies] = useCookies(["token"]);

  const onSubmit = (data: TransferRequest) => {
    const transferRequest: TransferRequest = {
      amount: data.amount,
      to: data.to,
      description: data.description,
    };

    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${cookies.token}`,
      },
      body: JSON.stringify({
        amount: Number(transferRequest.amount),
        to: Number(transferRequest.to),
        description: transferRequest.description,
      }),
    };

    fetch(API_URL + "/transactions/transfer", requestOptions)
      .then((props) => {
        if (props.ok) {
          return props.json();
        }
        throw new Error("failed to transfer");
      })
      .then((res) => {
        if (res.data) {
          setModal(true);
          console.log(data);
          setTdata(res.data);
          dispatch(fetchUser(cookies.token));
        }
      })
      .catch((err) => {
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

  useEffect(() => {
    dispatch(fetchUser(cookies.token));
  }, [dispatch, cookies.token]);

  return (
    <div>
      <NavbarLogin />
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
      <ModalDialog showModal={modal} tData={Tdata} />

      <section className="transfer__form">
        <div className="container">
          <div className="d-flex vertical flex-column align-items-center justify-content-center">
            <h1 className="fw-bold">Transfer</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
              <fieldset disabled>
                <label
                  htmlFor="exampleFormControlInput1"
                  className="form-label fw-bold mt-5"
                >
                  From
                </label>
                <div className="form-group mb-3">
                  <input
                    type="number"
                    placeholder={user?.id.toString()}
                    id="exampleFormControlInput1"
                    className="form-control control_login border border-dark"
                    aria-label="Username"
                    aria-describedby="basic-addon1"
                    style={{
                      width: "25rem",
                      height: "3rem",
                    }}
                  />
                </div>
              </fieldset>

              <label
                htmlFor="exampleFormControlInput2"
                className="form-label fw-bold "
              >
                To
              </label>
              <div className="form-group mb-3">
                <input
                  data-testid="toForm"
                  type="number"
                  id="exampleFormControlInput2"
                  className="form-control control_login border border-dark"
                  aria-label="Username"
                  aria-describedby="basic-addon1"
                  style={{
                    width: "25rem",
                    height: "3rem",
                    borderColor: errors.to ? "red" : "",
                  }}
                  {...register("to", { required: true })}
                />
                {errors.to && (
                  <span className="text-danger text-start">
                    Please insert destination wallet
                  </span>
                )}
              </div>

              <label
                htmlFor="exampleFormControlInput3"
                className="form-label fw-bold "
              >
                Amount
              </label>
              <div className="form-group mb-3">
                <input
                  type="number"
                  id="exampleFormControlInput3"
                  className="form-control control_login border border-dark"
                  aria-label="Username"
                  aria-describedby="basic-addon1"
                  placeholder="1.000.000"
                  style={{
                    width: "25rem",
                    height: "5rem",
                    fontSize: "2rem",
                    borderColor: errors.amount ? "red" : "",
                  }}
                  {...register("amount", {
                    required: true,
                    min: {
                      value: 1000,
                      message: "Minimal transfer Rp. 1.000",
                    },
                    max: {
                      value: 50000000,
                      message: "Maximal transfer Rp. 50.000.000",
                    },
                  })}
                />
                {errors.amount?.type === "required" ? (
                  <span data-testid="error" className="text-danger text-start">
                    Please insert amount
                  </span>
                ) : (
                  <span data-testid="error" className="text-danger text-start">
                    {errors.amount?.message}
                  </span>
                )}
              </div>

              <label
                htmlFor="exampleFormControlInput4"
                className="form-label fw-bold "
              >
                Description
              </label>
              <div className="form-group mb-3">
                <textarea
                  id="exampleFormControlInput4"
                  className="form-control control_login border border-dark"
                  aria-label="Username"
                  placeholder="Bayar Hutang"
                  aria-describedby="basic-addon1"
                  style={{
                    width: "25rem",
                    height: "5rem",
                    borderColor: errors.description ? "red" : "",
                  }}
                  {...register("description", { required: true })}
                />
                <span className="text-danger text-start">
                  {errors.description?.type === "required" ? (
                    <span className="text-danger text-start">
                      Please insert description
                    </span>
                  ) : (
                    <span className="text-danger text-start">
                      {errors.description?.message}
                    </span>
                  )}
                </span>
              </div>

              <div className="form-group mb-3 justify-content-center">
                <button
                  data-testid="submitBtn"
                  type="submit"
                  className="btn btn-info btn-lg w-100 text-white"
                >
                  Send
                </button>
              </div>
            </form>
          </div>
        </div>
      </section>
    </div>
  );
}
