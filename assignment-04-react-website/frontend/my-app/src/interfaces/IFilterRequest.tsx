export interface IFilterRequest {
  page: number;
  size: number;
  sortBy: string;
  sortDir: string;
  search: string;
  token: string;
  last: string;
}
