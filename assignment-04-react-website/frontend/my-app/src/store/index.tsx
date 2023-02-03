import { configureStore } from "@reduxjs/toolkit";
import thunk from "redux-thunk";
import logger from "redux-logger";
import userReducer from "./slices/userSlice";
import transactionReducer from "./slices/transactionSlice";

export const store = configureStore({
  reducer: {
    user: userReducer,
    transaction: transactionReducer,
  },
  middleware: [logger, thunk],
});

export type RootState = ReturnType<typeof store.getState>;
