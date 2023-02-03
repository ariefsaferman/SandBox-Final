export default interface IUser {
  id: number;
  email: string;
  first_name: string;
  last_name: string;
  wallet_id: number;
  wallet: IWallet;
}

interface IWallet {
  id: number;
  balance: number;
}
