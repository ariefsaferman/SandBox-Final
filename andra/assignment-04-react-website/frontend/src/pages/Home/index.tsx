import { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { toast, ToastContainer } from "react-toastify";
import FilterTransaction from "../../components/FilterTransaction";
import Hero from "../../components/Hero";
import Pagination from "../../components/Pagination";
import TransactionTable from "../../components/TransactionTable";
import useDebounce from "../../hooks/useDebounce";
import { ShowFilters } from "../../interfaces/show_filter";
import ITransaction from "../../interfaces/transaction";

const Home = () => {
  const [transactions, setTransactions] = useState<ITransaction[]>([]);
  const [page, setPage] = useState<number>(1);
  const size: number = 10;
  const [totalPage, setTotalPage] = useState<number>(0);

  const [showFilter, setShowFilter] = useState<ShowFilters>(
    ShowFilters.LAST_TEN
  );
  const [sortBy, setSortBy] = useState<string>("date");
  const [sortDirection, setSortDirection] = useState<string>("desc");
  const [search, setSearch] = useState<string>("");
  const debouncedSearch = useDebounce<string>(search);

  const [cookies] = useCookies(["token"]);
  const api = process.env.REACT_APP_API_URL;

  useEffect(() => {
    fetchTransactions();
  }, [
    page,
    cookies.token,
    api,
    sortBy,
    sortDirection,
    debouncedSearch,
    showFilter,
  ]);

  const fetchTransactions = () => {
    let url =
      api +
      "/transactions?page=" +
      page +
      "&size=" +
      size +
      "&sortBy=" +
      sortBy +
      "&sortDir=" +
      sortDirection +
      "&search=" +
      debouncedSearch +
      "&last=" +
      showFilter;

    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + cookies.token,
      },
    };

    fetch(url, requestOptions)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch transactions");
        return res.json();
      })
      .then((res) => {
        if (res.data) {
          setTransactions(res.data.data);
          changeTotalPage(res.data.count);
        }
      })
      .catch((err) => toast.error(err.message));
  };

  const changeTotalPage = (totalData: number) => {
    setTotalPage(Math.ceil(totalData / size));
  };

  const changePage = (page: number) => {
    setPage(page);
  };

  return (
    <div className="container">
      <ToastContainer />
      <Hero />
      <FilterTransaction
        showFilter={showFilter}
        setShowFilter={setShowFilter}
        sortBy={sortBy}
        setSortBy={setSortBy}
        sortDirection={sortDirection}
        setSortDirection={setSortDirection}
        search={search}
        setSearch={setSearch}
      />
      {transactions.length > 0 ? (
        <TransactionTable data={transactions} />
      ) : (
        <div className=" mt-5 text-secondary text-center">No data</div>
      )}
      <div className="float-end">
        <Pagination
          totalPage={totalPage}
          currentPage={page}
          onChangePage={changePage}
        />
      </div>
    </div>
  );
};

export default Home;
