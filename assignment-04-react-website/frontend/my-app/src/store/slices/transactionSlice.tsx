import {
  AnyAction,
  createAsyncThunk,
  createSlice,
  ThunkDispatch,
} from "@reduxjs/toolkit";
import { IFilterRequest } from "../../interfaces/IFilterRequest";
import {
  ITransaction,
  ITransactionOuter,
} from "../../interfaces/ITransactions";

export interface ITransactionState {
  transactions: ITransactionOuter | undefined;
  transactionLoading: boolean;
  transactionError: string | null;
}

export const initialState: ITransactionState = {
  transactions: undefined,
  transactionLoading: false,
  transactionError: null,
};

export const fetchTransaction = createAsyncThunk<
  ITransactionOuter,

  IFilterRequest,
  { rejectValue: string }
>(
  "FETCH_TRANSACTIONS",
  ({ token, page, sortBy, sortDir, search, last }, { rejectWithValue }) => {
    const API_URL = process.env.REACT_APP_API_URL;

    const requestOptions = {
      method: "GET",
      headers: {
        "content-type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    };

    return fetch(
      API_URL +
        "/transactions?page=" +
        page +
        "&sortBy=" +
        sortBy +
        "&sortDir=" +
        sortDir +
        "&search=" +
        search +
        "&last=" +
        last,
      requestOptions
    )
      .then((response) => {
        if (!response.ok) {
          throw new Error("failed to fetch transactions");
        }
        return response.json();
      })
      .then((data) => {
        return data.data;
      })
      .catch((err) => {
        return rejectWithValue(err.message);
      });
  }
);

export const transactionSlice = createSlice({
  name: "transaction",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchTransaction.pending, (state) => {
      return { ...state, transactionLoading: true, transactionError: null };
    });
    builder.addCase(fetchTransaction.fulfilled, (state, action) => {
      return {
        ...state,
        transactionLoading: false,
        transactions: action.payload,
      };
    });
    builder.addCase(fetchTransaction.rejected, (state, action) => {
      return action.payload
        ? {
            ...state,
            transactionLoading: false,
            transactionError: action.payload,
          }
        : {
            ...state,
            transactionLoading: false,
            transactionError: "Something went wrong",
          };
    });
  },
});

export default transactionSlice.reducer;
export type transactionDispatch = ThunkDispatch<
  ITransactionState,
  any,
  AnyAction
>;
