export interface IUser {
  id: number;
  first_name: string;
  last_name: string;
  wallet_id: number;
  wallet: IWallet;
}

export interface IWallet {
  id: number;
  balance: number;
}
