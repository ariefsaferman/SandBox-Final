import React, { useEffect, useState } from "react";
import "../style.scss";
import Navbar from "../../components/Navbar";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../store";
import { UserDispatch, fetchUser } from "../../store/slices/userSlice";
import {
  transactionDispatch,
  fetchTransaction,
} from "../../store/slices/transactionSlice";
import icnSearch from "../../../src/assets/images/icnSearch.png";
import { useCookies } from "react-cookie";
import "./index.scss";
import { IFilterRequest } from "../../interfaces/IFilterRequest";
import useDebounce from "../../hooks/useDebounce";

function setDate(date: string) {
  const dateObj = new Date(date);
  const month = dateObj.toLocaleString("default", { month: "long" });
  const day = dateObj.getDate();
  const year = dateObj.getFullYear();
  const time = dateObj.toLocaleTimeString([], {
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  });
  return `${time} - ${day} ${month} ${year} `;
}

function setCurrency(num: number | undefined): string | undefined {
  if (num) {
    return num.toLocaleString("id-ID", { minimumFractionDigits: 2 });
  }
  return "";
}

export default function HomePage() {
  const [cookies] = useCookies(["token"]);

  const { user, userError, userLoading } = useSelector(
    (state: RootState) => state.user
  );
  const { transactions, transactionError, transactionLoading } = useSelector(
    (state: RootState) => state.transaction
  );

  const [pageCount, setPageCount] = useState<number>(0);
  const [pageNumber, setPageNumber] = useState<number>(1);
  const [transactionSortBy, setTransactionSortBy] = useState<string>("date");
  const [transactionSortDir, setTransactionSortDir] = useState<string>("desc");
  const [search, setSearch] = useState<string>("");
  const [showLast, setShowLast] = useState<string>("");
  const debouncedValue = useDebounce<string>(search, 500);

  const dispatch: UserDispatch = useDispatch();
  const dispatch2: transactionDispatch = useDispatch();

  useEffect(() => {
    const request: IFilterRequest = {
      page: pageNumber,
      size: 10,
      sortBy: transactionSortBy,
      sortDir: transactionSortDir,
      search: debouncedValue,
      token: cookies.token,
      last: showLast,
    };

    dispatch(fetchUser(cookies.token));
    dispatch2(fetchTransaction(request));
  }, [
    dispatch,
    dispatch2,
    cookies.token,
    pageNumber,
    transactionSortBy,
    transactionSortDir,
    debouncedValue,
    showLast,
  ]);

  useEffect(() => {
    if (transactions) {
      setPageCount(Math.ceil(transactions?.count / transactions?.size));
    }
  }, [transactions]);

  return (
    <div>
      <Navbar />

      <section className="greet__section mt-5">
        <div className="container greet__title">
          <h1 className="fw-bold">Good Morning, {user?.first_name}!</h1>
          <h4>Account: {user?.wallet_id}</h4>
        </div>
      </section>

      <section className="balance__section">
        <div className="container balance__title">
          <h5>Balance:</h5>
          <h1 className="mt-4">IDR {setCurrency(user?.wallet.balance)}</h1>
        </div>
      </section>

      <section className="sort__section">
        <div className="container wrapp__boundary">
          <div className="row">
            <div className="col-lg-2  mt-2">
              <h3 className="text__color ">Show</h3>
            </div>
            <div className="col first__sort">
              <select
                className="form-select text-secondary show__sort"
                aria-label="Default select example"
                onChange={(e) => {
                  e.preventDefault();
                  setShowLast(e.target.value);
                }}
              >
                <option value="">Last 10 Transactions</option>
                <option value="thismonth">This month</option>
                <option value="lastmonth">Last month</option>
                <option value="thisyear">This year</option>
                <option value="lastyear">Last year</option>
              </select>
            </div>
            <div className="col">
              <h3 className="text__color text-end mt-2">Sort</h3>
            </div>
            <div className="col">
              <select
                className="form-select text-secondary sortBy__sort"
                aria-label="Default select example"
                onChange={(e) => {
                  e.preventDefault();
                  setTransactionSortBy(e.target.value);
                }}
              >
                <option value="date">Date</option>
                <option value="amount">Amount</option>
                <option value="to">To</option>
              </select>
            </div>
            <div className="col">
              <select
                className="form-select text-secondary sortDir__sort"
                aria-label="Default select example"
                onChange={(e) => {
                  e.preventDefault();
                  setTransactionSortDir(e.target.value);
                }}
              >
                <option value="desc">Descending</option>
                <option value="asc">Ascending</option>
              </select>
            </div>

            <div className="col">
              <div className="input-group mb-3 text-secondary ">
                <span className="input-group-text bg__color" id="basic-addon1">
                  <img src={icnSearch} alt="search icon" />
                </span>
                <input
                  type="text"
                  className="form-control search__sort  border__search"
                  placeholder="Search"
                  aria-label="Username"
                  aria-describedby="basic-addon1"
                  onChange={(e) => {
                    e.preventDefault();
                    setSearch(e.target.value);
                  }}
                />
              </div>
            </div>
          </div>
        </div>
      </section>

      <section className="table__section">
        <div className="container">
          <table className="table table-striped table-bordered mt-3">
            <thead>
              <tr>
                <th scope="col">Date & Time</th>
                <th scope="col">Type</th>
                <th scope="col">From/To</th>
                <th scope="col">Description</th>
                <th scope="col">Amount</th>
              </tr>
            </thead>
            <tbody>
              {transactions?.data.map((transaction) => {
                return (
                  <tr key={transaction.id} style={{ borderTopWidth: "4px" }}>
                    <td style={{ textAlign: "start", padding: "1rem" }}>
                      {setDate(transaction.created_at)}
                    </td>
                    <td style={{ textAlign: "start", padding: "1rem" }}>
                      {transaction.to_wallet_id ? "DEBIT" : "CREDIT"}
                    </td>
                    <td style={{ textAlign: "start", padding: "1rem" }}>
                      {transaction.to_wallet_id
                        ? transaction.to_wallet_id
                        : transaction.wallet_id}
                    </td>
                    <td style={{ textAlign: "start", padding: "1rem" }}>
                      {transaction.description}
                    </td>
                    <td
                      style={{ textAlign: "start", padding: "1rem" }}
                      className={transaction.to_wallet_id ? "" : "text-success"}
                    >
                      {transaction.to_wallet_id
                        ? `- ${setCurrency(transaction.amount)}`
                        : `+ ${setCurrency(transaction.amount)}`}
                    </td>
                  </tr>
                );
              })}
            </tbody>
          </table>
        </div>
      </section>

      <section className="footer__section">
        <div className="container">
          <nav aria-label="Page navigation example">
            <ul className="pagination justify-content-center">
              <li className={`page-item ${pageNumber === 1 ? "disabled" : ""}`}>
                <button
                  className="page-link"
                  onClick={() => setPageNumber(pageNumber - 1)}
                  {...(pageNumber === 1 ? { disabled: true } : {})}
                >
                  Previous
                </button>
              </li>
              {Array.from(Array(pageCount), (e, i) => {
                return (
                  <li className="page-item">
                    <button
                      onClick={() => setPageNumber(i + 1)}
                      className={`page-link ${
                        pageNumber === i + 1 ? "active" : ""
                      }`}
                    >
                      {i + 1}
                    </button>
                  </li>
                );
              })}
              <li
                className={`page-item ${
                  pageNumber === pageCount ? "disabled" : ""
                }`}
              >
                <button
                  className="page-link"
                  onClick={() => setPageNumber(pageNumber + 1)}
                >
                  Next
                </button>
              </li>
            </ul>
          </nav>
        </div>
      </section>
    </div>
  );
}
