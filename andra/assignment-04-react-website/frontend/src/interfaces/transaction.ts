import IUser from "./user";

export default interface ITransaction {
  id: number;
  to_wallet_id: number | null;
  to_user: IUser | null;
  amount: number;
  description: string;
  source_of_fund_id: number | null;
  wallet_id: number;
  created_at: string; 
}
