import { PaymentInterface } from "./IPayment";
import { PlaceInterface } from "./IPlace";
import { EnrollmentInterface } from "./IEnrollment";

//การทำให้ชื่อของ attribute ของ frontend กับ backend ตรงกัน
//
export interface BillInterface {
  ID: number;
  PaymentTypeID: number;
  PaymentType: PaymentInterface;
  EnrollmentID: number;
  Enrollment: EnrollmentInterface;
  PlaceID: number;
  Place: PlaceInterface;
  TotalPrice: number;
  BillTime: Date | null;
}
