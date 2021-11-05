import { StudentRecordsInterface } from "./IStudentRecord";
import { ManageCoursesInterface } from "./IManageCourse";

export interface EnrollmentTypesInterface {
    ID: number;
    Name: string;
}

export interface EnrollmentsInterface {
    ID: number;
    EnrollYear: number;
    EnrollTrimester: number;
    TotalCredit: number;
    EnrollDateTime: Date;

    OwnerID: number;
    Owner: StudentRecordsInterface;

    EnrollmentItems: EnrollmentItemsInterface[];
}

export interface EnrollmentItemsInterface {
    ID: number;

    EnrollmentID: number;
    Enrollment: EnrollmentsInterface;

    ManageCourseID: number;
    ManageCourse: ManageCoursesInterface;

    EnrollmentTypeID: number;
    EnrollmentType: EnrollmentTypesInterface;
}