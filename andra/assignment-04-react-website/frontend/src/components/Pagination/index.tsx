import { useEffect, useState } from "react";

type Props = {
  totalPage: number;
  currentPage: number;
  onChangePage: (page: number) => void;
};

const Pagination = (props: Props) => {
  const [start, setStart] = useState<number>(1);
  const [end, setEnd] = useState<number>(props.totalPage);
  const [arr, setArr] = useState<number[]>([]);

  useEffect(() => {
    setArr(
      Array.from({ length: end - start + 1 }, (_, index) => start + index)
    );
  }, [start, end]);

  useEffect(() => {
    if (props.totalPage > 5) {
      setEnd(5);
    }
  }, [props.totalPage]);

  useEffect(() => {
    if (props.currentPage > 3) {
      if (props.currentPage + 2 > props.totalPage) {
        setStart(props.totalPage - 4);
        setEnd(props.totalPage);
      } else {
        setStart(props.currentPage - 2);
        setEnd(props.currentPage + 2);
      }
    } else {
      setStart(1);
      setEnd(props.totalPage > 5 ? 5 : props.totalPage);
    }
  }, [props.currentPage, props.totalPage]);

  return (
    <nav aria-label="Page navigation">
      <ul className="pagination">
        {props.currentPage > 1 && (
          <li className="page-item">
            <button
              className="page-link"
              aria-label="Previous"
              onClick={() => props.onChangePage(1)}
            >
              <span aria-hidden="true">First</span>
            </button>
          </li>
        )}
        {arr.map((v) => (
          <li key={v} className="page-item">
            <button
              className={
                props.currentPage === v ? "page-link active" : "page-link"
              }
              onClick={() => props.onChangePage(v)}
            >
              {v}
            </button>
          </li>
        ))}
        {props.currentPage < props.totalPage && (
          <li className="page-item">
            <button
              className="page-link"
              aria-label="Next"
              onClick={() => props.onChangePage(props.currentPage + 1)}
            >
              <span aria-hidden="true">Next</span>
            </button>
          </li>
        )}
      </ul>
    </nav>
  );
};

export default Pagination;
