import React, { useEffect, useState } from "react";
import "../style.scss";
import Navbar from "../../components/Navbar";
import "./index.scss";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../store";
import { fetchUser, UserDispatch } from "../../store/slices/userSlice";
import { useForm } from "react-hook-form";
import { useCookies } from "react-cookie";
import { ITopUp } from "../../interfaces/ITopUp";
import { toast, ToastContainer } from "react-toastify";
import ModalDialog from "../TransferPage/components/ModalDialog";

interface ITopRequest {
  amount: number;
  source_of_fund_id: number;
}

interface ISourceOfFund {
  id: number;
  name: string;
}

export default function TopupPage() {
  const API_URL = process.env.REACT_APP_API_URL;
  const { user } = useSelector((state: RootState) => state.user);
  const dispatch: UserDispatch = useDispatch();
  const [cookies] = useCookies(["token"]);
  const [sourceOfFund, setSourceOfFund] = useState<ISourceOfFund[]>([]);
  const [topup, setTopup] = useState<ITopUp>();
  const [modal, setModal] = useState<boolean>(false);

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<ITopRequest>();

  useEffect(() => {
    setValue("source_of_fund_id", sourceOfFund[0]?.id);
  }, [sourceOfFund, setValue]);

  useEffect(() => {
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${cookies.token}`,
      },
    };

    fetch(`${API_URL}/transactions/source-of-funds`, requestOptions)
      .then((response) => response.json())
      .then((data) => {
        setSourceOfFund(data.data);
        console.log(data.data);
      })
      .catch((error) => console.log(error));
  }, [API_URL, cookies.token]);

  useEffect(() => {
    dispatch(fetchUser(cookies.token));
  }, [dispatch, cookies.token]);

  const onSubmit = (data: ITopRequest) => {
    const topRequest: ITopRequest = {
      amount: data.amount,
      source_of_fund_id: data.source_of_fund_id,
    };

    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${cookies.token}`,
      },
      body: JSON.stringify({
        amount: Number(topRequest.amount),
        source_of_fund_id: Number(topRequest.source_of_fund_id),
      }),
    };

    fetch(`${API_URL}/transactions/top-up`, requestOptions)
      .then((response) => {
        if (response.ok) {
          return response.json();
        }
        throw new Error("bad request");
      })
      .then((data) => {
        console.log(data);
        setTopup(data.data);
        setModal(true);
        dispatch(fetchUser(cookies.token));
      })
      .catch((error) => {
        toast.error(error, {
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
      <Navbar />
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
      <ModalDialog showModal={modal} tData={topup} />

      <section className="transfer__form">
        <div className="container">
          <div className="d-flex vertical flex-column align-items-center justify-content-center">
            <h1 className="fw-bold">Top Up</h1>
            <form onSubmit={handleSubmit(onSubmit)}>
              <label
                htmlFor="exampleFormControlInput1"
                className="form-label fw-bold mt-5"
              >
                From
              </label>
              <select
                className="form-select form-select-lg mb-3 from__transfer border border-dark"
                aria-label=".form-select-lg example"
                id="exampleFormControlInput1"
                {...register("source_of_fund_id", { required: true })}
                defaultValue={`${sourceOfFund[0]}`}
              >
                {sourceOfFund.map((item) => (
                  <option value={item.id} key={item.id}>
                    {item.name}
                  </option>
                ))}
              </select>

              <fieldset disabled>
                <label
                  htmlFor="exampleFormControlInput2"
                  className="form-label fw-bold"
                >
                  To
                </label>
                <div className="form-group mb-3 from__transfer">
                  <input
                    type="number"
                    placeholder={user?.wallet_id.toString()}
                    id="exampleFormControlInput2"
                    className="form-control control_login border border-dark"
                    aria-label="Username"
                    aria-describedby="basic-addon1"
                    style={{ height: "3rem" }}
                  />
                </div>
              </fieldset>

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
                      value: 50000,
                      message: "amount of money cannot be less than 50000",
                    },
                    max: {
                      value: 10000000,
                      message: "amount of money cannot be more than 10000000",
                    },
                  })}
                />
                {errors.amount?.type === "required" ? (
                  <span className="text-danger text-start">
                    Please insert amount money
                  </span>
                ) : (
                  <span className="text-danger text-start">
                    {errors.amount?.message}
                  </span>
                )}
              </div>

              <div className="form-group mb-3 justify-content-center">
                <button
                  type="submit"
                  className="btn btn-info btn-lg w-100 text-white"
                >
                  Top Up
                </button>
              </div>
            </form>
          </div>
        </div>
      </section>
    </div>
  );
}
