import { render } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";
import App from "./App";

test("renders", () => {
  const view = render(<App />, { wrapper: BrowserRouter });
  expect(view).toMatchSnapshot();
});
