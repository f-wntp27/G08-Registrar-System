import { StudentInterface } from "./IStudent";

export interface EnrollmentInterface {
  ID: number;
  UserID: number;
  User: StudentInterface;
  EnrollYear: number;
  EnrollTrimester: number;
  TotalCredit: number;
  EnrollDateTime: Date | null;
}
