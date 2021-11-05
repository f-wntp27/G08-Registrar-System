import { ProfessorsInterface } from "./IManageCourse";

export interface PrefixesInterface {
    ID: number;
    Value: string;
}

export interface FacultiesInterface {
    ID: number;
    Name: number;
}

export interface DepartmentsInterface {
    ID: number;
    Name: string;

    FacultyID: number;
    Faculty: FacultiesInterface;
}

export interface StudentRecordsInterface {
    ID: number;
    StudentCode: string;
    Prefix: PrefixesInterface;
    FirstName: string;
    LastName: string;
    PersonalId: string;

    Department: DepartmentsInterface;
    Advisor: ProfessorsInterface;
}
