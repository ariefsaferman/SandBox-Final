import { configureStore } from "@reduxjs/toolkit";
import logger from "redux-logger";
import userReducer from "./slices/userSlice";
import thunk from "redux-thunk";

export const store = configureStore({
  reducer: {
    user: userReducer,
  },
  middleware: [logger, thunk],
});

export type RootState = ReturnType<typeof store.getState>;
