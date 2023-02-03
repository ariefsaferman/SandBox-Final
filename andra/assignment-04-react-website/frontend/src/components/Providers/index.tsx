import React from "react";
import { CookiesProvider } from "react-cookie";
import { Provider } from "react-redux";
import { BrowserRouter } from "react-router-dom";
import { store } from "../../store";

type Props = {
  children: React.ReactNode;
};

const Providers = (props: Props) => {
  return (
    <CookiesProvider>
      <Provider store={store}>
        <BrowserRouter>{props.children}</BrowserRouter>
      </Provider>
    </CookiesProvider>
  );
};

export default Providers;
