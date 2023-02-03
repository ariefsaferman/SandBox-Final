import { fireEvent, render, screen } from "@testing-library/react";
import Transfer from "..";
import Providers from "../../../components/Providers";

test("render", () => {
  const view = render(<Transfer />, { wrapper: Providers });
  expect(view).toMatchSnapshot();
});

test("transfer form rendered", () => {
  render(<Transfer />, { wrapper: Providers });

  const form = screen.getByTestId("transfer-form");
  expect(form).toBeInTheDocument();
});

test("from input pre-filled with user wallet id", () => {
  render(<Transfer />, { wrapper: Providers });

  const fromInput = screen.getByTestId("from-input");
  expect(fromInput).not.toHaveValue("");
});

test("error when one or more input is empty", async () => {
  render(<Transfer />, { wrapper: Providers });

  const btnSubmit = screen.getByTestId("btn-submit");
  expect(btnSubmit).toBeInTheDocument();

  fireEvent.click(btnSubmit);
  const errors = await screen.findAllByTestId("error");

  expect(errors).not.toHaveLength(0);
});

test("error when amount is less than 1000", async () => {
  render(<Transfer />, { wrapper: Providers });

  const amountInput = screen.getByPlaceholderText("Amount");
  expect(amountInput).toBeInTheDocument();

  fireEvent.change(amountInput, { target: { value: 500 } });

  const btnSubmit = screen.getByTestId("btn-submit");
  expect(btnSubmit).toBeInTheDocument();

  fireEvent.click(btnSubmit);

  const error = await screen.findByText("Minimum transfer amount is IDR 1.000");
  expect(error).toBeInTheDocument();
});

test("error when amount is more than 50000000", async () => {
  render(<Transfer />, { wrapper: Providers });

  const amountInput = screen.getByPlaceholderText("Amount");
  expect(amountInput).toBeInTheDocument();

  fireEvent.change(amountInput, { target: { value: 60000000 } });

  const btnSubmit = screen.getByTestId("btn-submit");
  expect(btnSubmit).toBeInTheDocument();

  fireEvent.click(btnSubmit);

  const error = await screen.findByText(
    "Maximum transfer amount is IDR 50.000.000"
  );
  expect(error).toBeInTheDocument();
});

test("show modal when success", async () => {
  render(<Transfer />, { wrapper: Providers });

  const toInput = screen.getByPlaceholderText("To");
  expect(toInput).toBeInTheDocument();
  const amountInput = screen.getByPlaceholderText("Amount");
  expect(amountInput).toBeInTheDocument();
  const descriptionInput = screen.getByPlaceholderText("Description");
  expect(descriptionInput).toBeInTheDocument();

  fireEvent.change(toInput, { target: { value: 2 } });
  fireEvent.change(amountInput, { target: { value: 1000 } });
  fireEvent.change(descriptionInput, { target: { value: "test" } });

  const btnSubmit = screen.getByTestId("btn-submit");
  expect(btnSubmit).toBeInTheDocument();

  fireEvent.click(btnSubmit);

  const modal = await screen.findByTestId("transaction-modal");
  expect(modal).toBeInTheDocument();
});
