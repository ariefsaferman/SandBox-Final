const data = [];
let income = 0;
let expense = 0;
let balance = 0;

document.getElementById("form_id").addEventListener("submit", function (event) {
  event.preventDefault();
  let form = event.target.elements;
  let transactionInfo = form.transaction.value;
  let amount = form.amount.value;
  const transactionDetail = {
    transaction: transactionInfo,
    amount: amount,
  };

  intAmount = parseInt(amount);
  balance += intAmount;
  if (amount >= 0) {
    income += intAmount;
  } else {
    expense += Math.abs(intAmount);
  }

  data.push(transactionDetail);
  fillTable();
});

const fillTable = () => {
  const table = document.getElementById("table_id");
  const incomeId = document.getElementById("income");
  const expenseId = document.getElementById("expense");
  const balanceId = document.getElementById("balance");

  table.innerHTML = "";
  incomeId.innerHTML = numberWithCommas(income);
  expenseId.innerHTML = numberWithCommas(expense);
  balanceId.innerHTML = numberWithCommas(balance);
  data.forEach(function (value, index) {
    const decide = value.amount >= 0 ? "green" : "red";
    const sign = value.amount > 0 ? "+" : "";

    // Filling the table
    table.innerHTML += `<tr>
    <td style="
    font-size: 30px;
    color: #0f1a1abf;
    padding-left: 2rem;"> 
    <div class="d-flex">
    <button
    class="btn" onclick="deleteItem(${index})">
      <span class="delete__icon me-2 fs-3">
        <i class="fa-solid fa-x"></i>
      </span>
    </button>
    ${value.transaction}
  </div>
  </div>
    </td>

    <td style=
    "font-size: 30px;
    color: #0f1a1abf;
    padding-left: 2rem;
    border-right: 10px solid ${decide};
    text-align-last: end">${sign + numberWithCommas(value.amount)}
    </td>

    </tr>`;
  });
};

const deleteItem = (id) => {
  intAmount = parseInt(data[id].amount);
  if (intAmount >= 0) {
    income -= intAmount;
    balance -= intAmount;
  } else {
    expense += intAmount;
    balance -= intAmount;
  }

  data.splice(id);
  fillTable();
};

function numberWithCommas(x) {
  return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".");
}
