import { ReactComponent as SearchIcon } from "../../assets/images/search.svg";
import { ShowFilters } from "../../interfaces/show_filter";
import "./style.scss";

type Props = {
  showFilter: ShowFilters;
  setShowFilter: (showFilter: ShowFilters) => void;
  sortBy: string;
  setSortBy: (sortBy: string) => void;
  sortDirection: string;
  setSortDirection: (sortDirection: string) => void;
  search: string;
  setSearch: (search: string) => void;
};

const   FilterTransaction = (props: Props) => {
  return (
    <div className="filter__transaction d-flex gap-2 flex-wrap align-items-center justify-content-between">
      <div className="filter__transaction__show d-flex align-items-center">
        <label htmlFor="show" className="m-0">
          Show
        </label>
        <select
          className="form-select"
          onChange={(e) => {
            e.preventDefault();
            props.setShowFilter(e.target.value as ShowFilters);
          }}
        >
          <option value={ShowFilters.LAST_TEN}>Last 10 transaction</option>
          <option value={ShowFilters.THIS_MONTH}>This month</option>
          <option value={ShowFilters.LAST_MONTH}>Last month</option>
          <option value={ShowFilters.THIS_YEAR}>This year</option>
          <option value={ShowFilters.LAST_YEAR}>Last year</option>
        </select>
      </div>
      <div className="filter__transaction__sort row gx-0 gy-2">
        <div className="col-lg-6 d-flex align-items-center">
          <label htmlFor="sort" className="m-0 me-3">
            Sort by
          </label>
          <select
            className="form-select me-2"
            defaultValue={props.sortBy}
            onChange={(e) => {
              e.preventDefault();
              props.setSortBy(e.target.value);
            }}
          >
            <option value="date">Date</option>
            <option value="amount">Amount</option>
            <option value="to">To</option>
          </select>
          <select
            className="form-select"
            defaultValue={props.sortDirection}
            onChange={(e) => {
              e.preventDefault();
              props.setSortDirection(e.target.value);
            }}
          >
            <option value="asc">Ascending</option>
            <option value="desc">Descending</option>
          </select>
        </div>
        <div className="col-lg-6">
          <div className="input-group filter__search">
            <span className="input-group-text border-end-0">
              <SearchIcon />
            </span>
            <input
              type="text"
              className="form-control border-start-0"
              placeholder="Search"
              value={props.search}
              onChange={(e) => {
                e.preventDefault();
                props.setSearch(e.target.value);
              }}
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default FilterTransaction;
