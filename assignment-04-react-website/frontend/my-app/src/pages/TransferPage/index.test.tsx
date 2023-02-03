import React from "react";
import { fireEvent, render, screen } from "@testing-library/react";
import { Provider } from "react-redux";
import TransferPage from "./index";
import { CookiesProvider } from "react-cookie";
import { store } from "../../store";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import userEvent from "@testing-library/user-event";

type Props = {
  children: React.ReactNode;
};

test("Should return error when user input invalid to wallet", async () => {
  render(
    <Provider store={store}>
      <BrowserRouter>
        <TransferPage />
      </BrowserRouter>
    </Provider>
  );
  const screenTransfer = screen.getByTestId("toForm");
  await userEvent.click(screenTransfer);
  await userEvent.keyboard("0");
  fireEvent.click(screen.getByTestId("submitBtn"));
  const screenText = screen.getByTestId("error");
  expect(screenText).toBeInTheDocument();
});

const Providers = (props: Props) => {
  return (
    <CookiesProvider>
      <Provider store={store}>
        <BrowserRouter>{props.children}</BrowserRouter>
      </Provider>
    </CookiesProvider>
  );
};

test("render snapshot", () => {
  const view = render(<TransferPage />, { wrapper: Providers });
  expect(view).toMatchSnapshot();
});
