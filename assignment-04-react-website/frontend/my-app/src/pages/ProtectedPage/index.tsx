import React from "react";
import { useCookies } from "react-cookie";
import { Navigate, Outlet, useLocation } from "react-router-dom";

export default function   ProtectedPage() {
  let location = useLocation();
  const [cookies] = useCookies(["token"]);

  if (cookies.token) {
    return <Outlet />;
  }

  return <Navigate to="login" replace state={{ from: location }} />;
}

export function ProtectReversePage() {
  let location = useLocation();
  const [cookies] = useCookies(["token"]);

  if (!cookies.token) {
    return <Outlet />;
  }

  return <Navigate to="/" replace state={{ from: location }} />;
}
  