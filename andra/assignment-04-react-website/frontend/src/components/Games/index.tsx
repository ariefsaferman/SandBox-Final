import { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import { useDispatch } from "react-redux";
import { toast, ToastContainer } from "react-toastify";
import { fetchUser, UserDispatch } from "../../store/slices/userSlice";
import "./style.scss";

const GamesSection = () => {
  const [prizes, setPrizes] = useState<number[]>([
    0, 0, 100000, 100000, 100000, 500000, 500000, 500000, 1000000,
  ]);
  const [chances, setChances] = useState<number>(3);
  const [show, setShow] = useState<boolean>(false);
  const [chosen, setChosen] = useState<number | undefined>(undefined);
  const [cookies] = useCookies(["token"]);
  const dispatch: UserDispatch = useDispatch();
  const api = process.env.REACT_APP_API_URL;

  useEffect(() => {
    setPrizes(shuffle(prizes));
  }, []);

  function shuffle(array: number[]) {
    const newArray = [...array];
    const length = newArray.length;

    for (let start = 0; start < length; start++) {
      const randomPosition = Math.floor(
        (newArray.length - start) * Math.random()
      );
      const randomItem = newArray.splice(randomPosition, 1);

      newArray.push(...randomItem);
    }

    return newArray;
  }

  const handleClick = (i: number) => {
    setShow(true);
    setChosen(i);
    if (prizes[i] === 0) {
      alert("You got nothing");
      setChances(chances - 1);
      return;
    }
    topup(prizes[i]);
  };

  const topup = (amount: number) => {
    const payload = {
      amount: amount,
      source_of_fund_id: 3,
    };
    const requestOptions = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + cookies.token,
      },
      body: JSON.stringify(payload),
    };

    fetch(api + "/transactions/top-up", requestOptions)
      .then((res) => {
        if (!res.ok) throw new Error("Failed to topup");
        return res.json();
      })
      .then((res) => {
        if (res.data) {
          dispatch(fetchUser(cookies.token));
          alert("You got " + amount);
        }
      })
      .catch((err) => toast.error(err.message))
      .finally(() => {
        setPrizes(shuffle(prizes));
        setChosen(undefined);
        setShow(false);
        setChances(chances - 1);
      });
  };

  return (
    <div className="games">
      <ToastContainer />
      <h2 className="games__title">Games</h2>
      <p className="games__subtitle">Choose random box below to get reward!</p>
      <p className="games__subtitle">Chance: {chances}</p>
      <section className="games__boxes">
        {prizes.map((prize, index) => (
          <div
            key={index}
            className={
              show && index === chosen
                ? "games__box shadow bg-chosen text-white"
                : chances === 0
                ? "games__box shadow disabled"
                : "games__box shadow"
            }
            onClick={() => handleClick(index)}
          >
            <span className={show ? "" : "d-none"}>{prize}</span>
          </div>
        ))}
      </section>
    </div>
  );
};

export default GamesSection;
