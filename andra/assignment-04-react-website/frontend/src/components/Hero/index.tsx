import { useEffect } from "react";
import { useCookies } from "react-cookie";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "../../store";
import { fetchUser, UserDispatch } from "../../store/slices/userSlice";
import "./style.scss";

const Hero = () => {
  const { user } = useSelector((state: RootState) => state.user);
  const [cookies] = useCookies(["token"]);
  const dispatch: UserDispatch = useDispatch();

  useEffect(() => {
    dispatch(fetchUser(cookies.token));
  }, [dispatch, cookies.token]);

  return (
    <div className="container hero">
      <h2 className="hero__title">Good Morning, {user?.first_name}!</h2>
      <div className="d-flex justify-content-between hero__subtitle">
        <p>Account: {user?.wallet.id}</p>
        <p>Balance:</p>
      </div>
      <div className="d-flex justify-content-end hero__balance">
        <h3>
          IDR{" "}
          {user?.wallet.balance.toLocaleString("id-ID", {
            minimumFractionDigits: 2,
          })}
        </h3>
      </div>
    </div>
  );
};

export default Hero;
