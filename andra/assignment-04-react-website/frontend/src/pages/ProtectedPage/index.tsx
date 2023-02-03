import { useCookies } from "react-cookie";
import { Navigate, Outlet, useLocation } from "react-router-dom";

const ProtectedPage = () => {
  const [cookies] = useCookies(["token"]);
  const location = useLocation();

  if (!cookies.token) {
    return <Navigate to={"/login"} replace state={{ from: location }} />;
  }

  return <Outlet />;
};

export default ProtectedPage;
