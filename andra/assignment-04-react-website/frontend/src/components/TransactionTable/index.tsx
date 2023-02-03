import moment from "moment";
import ITransaction from "../../interfaces/transaction";
import "./style.scss";

type Props = {
  data: ITransaction[];
};

const TransactionTable = (props: Props) => {
  return (
    <div className="table-responsive transaction__table">
      <table className="table table-striped table-bordered">
        <thead>
          <tr>
            <th scope="col">Date & Time</th>
            <th scope="col">Type</th>
            <th scope="col">From / To</th>
            <th scope="col">Description</th>
            <th scope="col">Amount</th>
          </tr>
        </thead>
        <tbody>
          {props.data.map((transaction) => (
            <tr key={transaction.id}>
              <td>
                {moment(transaction.created_at, "YYYY-MM-DD HH:mm:ss").format(
                  "HH:mm - DD MMMM YYYY"
                )}
              </td>
              <td>{transaction.to_wallet_id ? "DEBIT" : "CREDIT"}</td>
              <td>
                {transaction.to_wallet_id ?? transaction.source_of_fund_id}
              </td>
              <td>{transaction.description}</td>
              <td
                className={
                  transaction.to_wallet_id
                    ? "transaction__debit"
                    : "transaction__credit"
                }
              >
                {transaction.amount.toLocaleString("id-ID", {
                  minimumFractionDigits: 2,
                })}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default TransactionTable;
