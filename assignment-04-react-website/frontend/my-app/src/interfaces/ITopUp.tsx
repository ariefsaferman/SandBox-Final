export interface ITopUp {
  id: number;
  to_wallet_id: number | null;
  to_user: number | null;
  amount: number;
  description: string;
  source_of_fund_id: number;
  wallet_id: number;
}
