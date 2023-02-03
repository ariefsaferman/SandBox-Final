import {
  AnyAction,
  createAsyncThunk,
  createSlice,
  ThunkDispatch,
} from "@reduxjs/toolkit";
import IUser from "../../interfaces/user";

export interface IUserState {
  user: IUser | null;
  userLoading: boolean;
  userError: string | null;
}

const initialState: IUserState = {
  user: null,
  userLoading: true,
  userError: null,
};

export const fetchUser = createAsyncThunk<
  IUser,
  string,
  { rejectValue: string }
>("FETCH_USER", (token, { rejectWithValue }) => {
  const api = process.env.REACT_APP_API_URL;
  const requestOptions = {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: "Bearer " + token,
    },
  };
  return fetch(api + "/details", requestOptions)
    .then((response) => {
      if (!response.ok) throw new Error("failed to fetch user");
      return response.json();
    })
    .then((res) => {
      if (res.data) {
        return res.data;
      }
    })
    .catch((error) => {
      return rejectWithValue(error.message);
    });
});

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchUser.fulfilled, (state, action) => {
      return { ...state, user: action.payload, userLoading: false };
    });
    builder.addCase(fetchUser.pending, (state, action) => {
      return { ...state, userError: null, userLoading: true };
    });
    builder.addCase(fetchUser.rejected, (state, action) => {
      return action.payload
        ? { ...state, userError: action.payload }
        : { ...state, userError: "unknown error" };
    });
  },
});

export default userSlice.reducer;
export type UserDispatch = ThunkDispatch<IUserState, any, AnyAction>;
