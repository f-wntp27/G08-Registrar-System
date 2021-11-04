export interface StudentInterface {
  ID: number;
  FirstName: string;
  LastName: string;
  StudentCode: string;
  Prefix: PrefixInterface;
}

export interface PrefixInterface {
  ID: number;
  Value: string;
}
