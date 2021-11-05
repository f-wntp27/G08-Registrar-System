export interface ProfessorsInterface {
    ID: number;
    TeacherName: string;
    TeacherEmail: string;
    ProfessorCode: string;
}

export interface CoursesInterface {
    ID: number;
    CourseCode: string;
    Name: string;
    Credit: number;
}

export interface RoomsInterface {
    ID: number;
    Number: number;
    StudentCount: number;
}

export interface ManageCoursesInterface {
    ID: number;
    Group: number;
    TeachingTime: Date;
    UngraduatedYear: number;
    Trimester: number;

    RoomID: number;
    Room: RoomsInterface;

    CourseID: number;
    Course: CoursesInterface;

    ProfessorID: number;
    Professor: ProfessorsInterface;
}
