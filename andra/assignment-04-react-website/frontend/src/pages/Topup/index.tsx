import TopupForm from "../../components/Form/Topup";
import "./style.scss";

const Topup = () => {
  return (
    <div className="container">
      <div className="d-flex vh-90 flex-column align-items-center justify-content-center">
        <h2 className="topup__title">Top Up</h2>
        <TopupForm />
      </div>
    </div>
  );
};

export default Topup;
