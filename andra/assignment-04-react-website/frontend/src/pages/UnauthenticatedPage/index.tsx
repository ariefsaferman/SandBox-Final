import { useCookies } from "react-cookie";
import { Navigate, Outlet, useLocation } from "react-router-dom";

const UnauthenticatedPage = () => {
  const [cookies] = useCookies(["token"]);
  const location = useLocation();
  const from = location.state?.from?.pathname || "/";

  if (cookies.token) {
    return <Navigate to={from} replace state={{ from: location }} />;
  }

  return <Outlet />;
};

export default UnauthenticatedPage;
