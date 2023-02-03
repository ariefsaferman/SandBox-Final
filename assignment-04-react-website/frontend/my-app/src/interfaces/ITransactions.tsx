import { IUser } from "./IUser";

export interface ITransaction {
  id: number;
  to_wallet_id: number;
  to_user: IUser;
  amount: number;
  description: string;
  source_of_fund_id: number | null;
  wallet_id: number;
  created_at: string;
}

export interface ITransactionOuter {
  page: number;
  count: number;
  size: number;
  data: ITransaction[];
}
