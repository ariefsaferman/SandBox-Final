import { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { useForm } from "react-hook-form";
import { useDispatch, useSelector } from "react-redux";
import { toast, ToastContainer } from "react-toastify";
import ITransaction from "../../interfaces/transaction";
import { RootState } from "../../store";
import { fetchUser, UserDispatch } from "../../store/slices/userSlice";
import TransactionModal from "../TransactionModal";
import "./style.scss";

type Payload = {
  amount: number;
  to: number;
  description: string;
};

const TransferForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Payload>();
  const { user } = useSelector((state: RootState) => state.user);
  const [showModal, setShowModal] = useState(false);
  const [transaction, setTransaction] = useState<ITransaction>();
  const [cookies] = useCookies(["token"]);
  const dispatch: UserDispatch = useDispatch();
  const api = process.env.REACT_APP_API_URL;

  useEffect(() => {
    dispatch(fetchUser(cookies.token));
  }, [cookies.token, dispatch]);

  const onSubmit = (data: Payload) => {
    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + cookies.token,
      },
      body: JSON.stringify(data),
    };

    fetch(api + "/transactions/transfer", requestOptions)
      .then(async (res) => {
        const resJson = await res.json();
        if (!res.ok) throw new Error(resJson.message);
        return resJson;
      })
      .then((res) => {
        if (res.data) {
          dispatch(fetchUser(cookies.token));
          setTransaction(res.data);
          setShowModal(true);
        }
      })
      .catch((err) => toast.error(err.message));
  };

  return (
    <form
      data-testid="transfer-form"
      className="form__transfer"
      onSubmit={handleSubmit(onSubmit)}
    >
      <TransactionModal show={showModal} type="Transfer" data={transaction} />
      <ToastContainer />
      <div className="form-group">
        <label htmlFor="from">From</label>
        <input
          data-testid="from-input"
          className="form-control"
          disabled
          type="number"
          value={user?.wallet_id}
        />
      </div>
      <div className="form-group">
        <label htmlFor="">To</label>
        <input
          className="form-control"
          type="number"
          placeholder="To"
          {...register("to", {
            required: true,
            setValueAs: (value) => parseInt(value),
          })}
          style={{ borderColor: errors.to && "red" }}
        />
        {errors.to && (
          <span data-testid="error" className="error">
            Please insert transfer target
          </span>
        )}
      </div>
      <div className="form-group">
        <label htmlFor="">Amount</label>
        <input
          className="form-control form-control-lg"
          type="number"
          placeholder="Amount"
          min={1000}
          max={50000000}
          {...register("amount", {
            required: {
              value: true,
              message: "Please insert amount of transfer",
            },
            min: {
              value: 1000,
              message: "Minimum transfer amount is IDR 1.000",
            },
            max: {
              value: 50000000,
              message: "Maximum transfer amount is IDR 50.000.000",
            },
            setValueAs: (value) => parseInt(value),
          })}
          style={{ borderColor: errors.amount && "red" }}
        />
        {errors.amount && (
          <span data-testid="error" className="error">
            {errors.amount?.message}
          </span>
        )}
      </div>
      <div className="form-group">
        <label htmlFor="">Description</label>
        <textarea
          className="form-control"
          placeholder="Description"
          {...register("description", { required: true })}
          style={{ borderColor: errors.description && "red" }}
        />
        {errors.description && (
          <span data-testid="error" className="error">
            Please insert description
          </span>
        )}
      </div>
      <button
        data-testid="btn-submit"
        type="submit"
        className="btn btn-lg w-100"
      >
        Send
      </button>
    </form>
  );
};

export default TransferForm;
