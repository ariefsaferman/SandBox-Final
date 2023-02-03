import {
  AnyAction,
  createAsyncThunk,
  createSlice,
  ThunkDispatch,
} from "@reduxjs/toolkit";
import { IUser } from "../../interfaces/IUser";

export interface IUserState {
  user: IUser | null;
  userLoading: boolean;
  userError: string | null;
}

export const initialState: IUserState = {
  user: null,
  userLoading: false,
  userError: null,
};

export const fetchUser = createAsyncThunk<
  IUser,
  string,
  { rejectValue: string }
>("FETCH_USER", (token, { rejectWithValue }) => {
  const API_URL = process.env.REACT_APP_API_URL;
  const requestOptions = {
    method: "GET",
    headers: {
      "content-type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  };

  return fetch(API_URL + "/details", requestOptions)
    .then((response) => {
      if (!response.ok) {
        if (response.status === 401) {
          throw new Error("Unauthorized");
        }
        throw new Error("Internal Server Error");
      }
      return response.json();
    })
    .then((data) => {
      return data.data;
    })
    .catch((err) => {
      return rejectWithValue(err.message);
    });
});

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchUser.pending, (state) => {
      return { ...state, userLoading: true, userError: null };
    });
    builder.addCase(fetchUser.fulfilled, (state, action) => {
      return { ...state, userLoading: false, user: action.payload };
    });
    builder.addCase(fetchUser.rejected, (state, action) => {
      return action.payload
        ? { ...state, userLoading: false, userError: action.payload }
        : { ...state, userLoading: false, userError: "Something went wrong" };
    });
  },
});

export default userSlice.reducer;
export type UserDispatch = ThunkDispatch<IUserState, any, AnyAction>;
