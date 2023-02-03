import { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { useForm } from "react-hook-form";
import { useDispatch, useSelector } from "react-redux";
import { toast, ToastContainer } from "react-toastify";
import ISourceOfFund from "../../interfaces/source_of_fund";
import ITransaction from "../../interfaces/transaction";
import { RootState } from "../../store";
import { fetchUser, UserDispatch } from "../../store/slices/userSlice";
import TransactionModal from "../TransactionModal";
import "./style.scss";

type Payload = {
  amount: number;
  source_of_fund_id: number;
};

const TopupForm = () => {
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<Payload>();
  const { user } = useSelector((state: RootState) => state.user);
  const [sofs, setSofs] = useState<ISourceOfFund[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [transaction, setTransaction] = useState<ITransaction>();
  const [cookies] = useCookies(["token"]);
  const dispatch: UserDispatch = useDispatch();
  const api = process.env.REACT_APP_API_URL;

  useEffect(() => {
    dispatch(fetchUser(cookies.token));
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + cookies.token,
      },
    };

    fetch(api + "/transactions/source-of-funds", requestOptions)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch source of funds");
        return res.json();
      })
      .then((res) => {
        if (res.data) {
          setSofs(res.data);
        }
      })
      .catch((err) => toast.error(err.message));
  }, [api, cookies.token, dispatch]);

  useEffect(() => {
    setValue("source_of_fund_id", sofs[0]?.id);
  }, [setValue, sofs]);

  const onSubmit = (data: Payload) => {
    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + cookies.token,
      },
      body: JSON.stringify(data),
    };

    fetch(api + "/transactions/top-up", requestOptions)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to topup");
        return res.json();
      })
      .then((res) => {
        if (res.data) {
          dispatch(fetchUser(cookies.token));
          setShowModal(true);
          setTransaction(res.data);
        }
      })
      .catch((err) => toast.error(err.message));
  };

  return (
    <form className="form__topup" onSubmit={handleSubmit(onSubmit)}>
      <TransactionModal show={showModal} type="Top Up" data={transaction} />
      <ToastContainer />
      <div className="form-group">
        <label htmlFor="from">From</label>
        <select
          className="form-select"
          {...register("source_of_fund_id", {
            required: true,
            setValueAs: (value) => parseInt(value),
          })}
        >
          {sofs.map((sof) => (
            <option key={sof.id} value={sof.id}>
              {sof.name}
            </option>
          ))}
        </select>
        {errors.source_of_fund_id && (
          <span className="error">This field is required</span>
        )}
      </div>
      <div className="form-group">
        <label htmlFor="To">To</label>
        <input
          className="form-control"
          disabled
          type="number"
          defaultValue={user?.wallet_id}
        />
      </div>
      <div className="form-group">
        <label htmlFor="To">Amount</label>
        <input
          className="form-control form-control-lg"
          type="number"
          placeholder="Amount"
          min={50000}
          {...register("amount", {
            required: {
              value: true,
              message: "Please insert amount of transfer",
            },
            min: {
              value: 1000,
              message: "Minimum transfer amount is IDR 50.000",
            },
            max: {
              value: 50000000,
              message: "Maximum transfer amount is IDR 10.000.000",
            },
            setValueAs: (value) => parseInt(value),
          })}
        />
        {errors.amount && (
          <span className="error">{errors.amount?.message}</span>
        )}
      </div>
      <button type="submit" className="btn btn-lg w-100">
        Top Up
      </button>
    </form>
  );
};

export default TopupForm;
