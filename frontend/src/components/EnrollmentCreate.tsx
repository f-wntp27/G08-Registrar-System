import React, { ChangeEvent,
    useEffect,
    useState,
    SyntheticEvent, 
    Fragment
} from 'react';
import { Link as RouterLink } from "react-router-dom";
import Box  from '@material-ui/core/Box';
import Container from '@material-ui/core/Container';
import Grid from '@material-ui/core/Grid';
import Paper from '@material-ui/core/Paper';
import Button from "@material-ui/core/Button";
import FormControl from "@material-ui/core/FormControl";
import FormHelperText from '@material-ui/core/FormHelperText';
import MenuItem from '@material-ui/core/MenuItem';
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import Select from '@material-ui/core/Select';
import Snackbar from "@material-ui/core/Snackbar";
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import DeleteIcon from '@material-ui/icons/Delete'
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import { MuiPickersUtilsProvider, KeyboardDateTimePicker } from "@material-ui/pickers/";
import DateFnsUtils from "@date-io/date-fns";
import { StudentRecordsInterface } from  "../models/IStudentRecord";
import { EnrollmentTypesInterface, 
  EnrollmentsInterface, 
  EnrollmentItemsInterface,
} from "../models/IEnrollRegistration";
import { CoursesInterface, ManageCoursesInterface } from "../models/IManageCourse";
  
const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) => 
  createStyles({
    root: {flexGrow: 2},
    container: {marginTop: theme.spacing(3)},
    paper: {padding: theme.spacing(3)},
    table: {minWidth: 600},
  })
);

export default function EnrollmentCreate() {
  const classes = useStyles();
  const [courses, setCourses] = useState<CoursesInterface[]>([]);
  const [selectedCourse, setSelectedCourse] = useState<Partial<CoursesInterface>>({ID: 0});
  const [manageCourses, setManageCourses] = useState<ManageCoursesInterface[]>([]);
  const [enrollmentTypes, setEnrollmentTypes] = useState<EnrollmentTypesInterface[]>([]);
  const [enrollment, setEnrollment] = useState<Partial<EnrollmentsInterface>>({EnrollYear: 2021, EnrollTrimester: 1});
  const [enrollmentItem, setEnrollmentItem] = useState<Partial<EnrollmentItemsInterface>>({EnrollmentTypeID: 0});
  const [enrollmentItemArray, setEnrollmentItemArray] = useState<Partial<EnrollmentItemsInterface>[]>([]);
  const [enrollmentItemSubmit, setEnrollmentItemSubmit] = useState<Partial<EnrollmentItemsInterface>[]>([]);
  const [totalCredit, setTotalCredit] = useState<number>(0);

  // If combo box is not selected, it's alert error.
  const [hasError, setHasError] = useState<boolean>(false);

  // If Year or Trimester is not in range, it's alert error.
  const [outOfRange, setOutOfRange] = useState<boolean>(false);
  const [outOfRangeMsg, setOutOfRangeMsg] = useState<string>("");

  /* --- Year and Trimester components --- */
  const handleYearTrimesterChange = (event: ChangeEvent<{id?: string; value: unknown}>) => {
    const id = event.target.id as keyof typeof enrollment;
    const value = event.target.value as string;
    setEnrollment({...enrollment, [id]: value === "" ? "" : parseInt(value),});
  }

  /* --- Course component --- */
  const getCourses = async() => {
    const apiUrl = "http://localhost:8080/courses";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if(res.data) {
          setCourses(res.data)
        } else {
          console.log("else")
        }
      });
  };
  const handleCourseChange = (event: ChangeEvent<{name?: string; value: unknown}>) => {
    const name = event.target.name as keyof typeof selectedCourse;
    setSelectedCourse({...selectedCourse, [name]: event.target.value,});

    // Find ManageCourse from course ID after select course
    getManageCoursesFromCourseID(event.target.value as number);
  };

  /* --- Manage Course component --- */  
  //getManageCoursesFromCourseID with parameter
  const getManageCoursesFromCourseID = async(CourseID: number | undefined) => {
    if (CourseID === 0) {
      setManageCourses([]);
      return
    }
    const apiUrl = "http://localhost:8080/manage_courses/course/" + CourseID;
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    }

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if(res.data) {
          setManageCourses(res.data)
        } else {
          console.log("else")
        }
      });
  } 

  /* --- Enrollment Type component --- */
  const getCourseTypes = async() => {
    const apiUrl = "http://localhost:8080/enrollment_types";
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    }

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if(res.data) {
          setEnrollmentTypes(res.data)
        } else {
          console.log("else")
        }
      });
  } 

  /* --- Enrollment Item component --- */
  const handleEnrollmentItemChange = (event: ChangeEvent<{name?: string; value: unknown}>) => {
    const name = event.target.name as keyof typeof enrollmentItem;
    setEnrollmentItem({...enrollmentItem, [name]: event.target.value,});
  };
  
  const dropItemTable = (index: number) => {
    let afterDropItem = enrollmentItemArray.filter((_, i) => i !== index);
    let afterDropSubmit = enrollmentItemSubmit.filter((_, i) => i !== index);
    setEnrollmentItemArray(afterDropItem);
    setEnrollmentItemSubmit(afterDropSubmit);
    setTotalCredit(totalCredit - (enrollmentItemArray[index].ManageCourse?.Course.Credit || 0));
  }

  const addItemTable = (item: ManageCoursesInterface) => {
    setHasError(false);
    if (enrollmentItem.EnrollmentTypeID === 0) {
      setHasError(true);
      return
    }
    let itemTable = {
      ManageCourse: item,
      EnrollmentType: enrollmentTypes.find(et => et.ID === enrollmentItem.EnrollmentTypeID),
    }
    let itemSubmit = {
      ManageCourseID: item.ID,
      EnrollmentTypeID: enrollmentItem.EnrollmentTypeID,
    }
    setEnrollmentItemArray([...enrollmentItemArray, itemTable]);
    setEnrollmentItemSubmit([...enrollmentItemSubmit, itemSubmit]);
    setTotalCredit(totalCredit + (item.Course.Credit || 0));

    // Reset field after add course to table
    setSelectedCourse({ID: 0});
    setManageCourses([]);
    setEnrollmentItem({EnrollmentTypeID: 0});
  };

  /* --- Enroll Date Time component --- */
  const [enrollDateTime, setEnrollDateTime] = useState<Date|null>(new Date());
  const handleEnrollDateTime = (date: Date | null) => {
    setEnrollDateTime(date);
  }

  /* --- SUBMIT --- */
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  // If enrollment item is empty, it's alert warning.
  const [isItemEmpty, setIsItemEmpty] = useState<boolean>(false);
  const handleClose = (event?: SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
    setIsItemEmpty(false);
    setOutOfRange(false);
  };
  const submit = () => {
    if (Number(enrollment?.EnrollYear) < 2010 || Number(enrollment?.EnrollYear) > 2050) {
      setOutOfRange(true);
      setOutOfRangeMsg("กรุณากรอกปีการศึกษาเป็นปัจจุบัน");
      return
    }
    if (Number(enrollment?.EnrollTrimester) < 1 || Number(enrollment?.EnrollTrimester) > 3) {
      setOutOfRange(true);
      setOutOfRangeMsg("กรุณากรอกภาคการศึกษาในช่วง 1-3");
      return
    }
    if (enrollmentItemSubmit.length === 0) {
      setIsItemEmpty(true);
      return
    }
    let student: StudentRecordsInterface = JSON.parse(localStorage.getItem("student") || "");
    let data = {
      OwnerID: student.ID,
      EnrollYear: enrollment.EnrollYear,
      EnrollTrimester: enrollment.EnrollTrimester,
      TotalCredit: totalCredit,
      EnrollDateTime: enrollDateTime,
      EnrollmentItems: enrollmentItemSubmit,
    };
    const apiUrl = "http://localhost:8080/enrollment";
    const requestOptionsPost = {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };

    fetch(apiUrl, requestOptionsPost)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setSuccess(true);
        } else {
          setError(true);
        }
    });
  }
  console.log(enrollment);

  useEffect(() => {
    getCourses();
    getCourseTypes();
  }, []);
  
  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Snackbar open={success} autoHideDuration={3000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="success">
            บันทึกข้อมูลสำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar open={error} autoHideDuration={3000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="error">
            บันทึกข้อมูลไม่สำเร็จ
          </Alert>
        </Snackbar>
        <Snackbar open={isItemEmpty} autoHideDuration={3000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="warning">
            กรุณาเลือกรายวิชาที่ต้องการลงทะเบียน
          </Alert>
        </Snackbar>
        <Snackbar open={outOfRange} autoHideDuration={3000} onClose={handleClose}>
          <Alert onClose={handleClose} severity="warning">
            {outOfRangeMsg}
          </Alert>
        </Snackbar>
        <Paper className={classes.paper}>
          <Box display="flex" style={{marginBottom: ".5rem"}}>
            <Box flexGrow={1}>
              <Typography
                component="h2"
                variant="h6"
                color="primary"
                gutterBottom
              >
                ลงทะเบียนรายวิชา
              </Typography>
            </Box>
            <Box flexGrow={1}>
              <Button
                component={RouterLink}
                to="/enrollReg/history"
                variant="contained"
                color="primary"
                style={{float: "right"}}
              >
                ผลการลงทะเบียน
              </Button>
            </Box>
          </Box>
          <Grid container spacing={1}>

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined" size="small">
                <p>ปีการศึกษา</p>
                <TextField
                  id="EnrollYear"
                  type="number"
                  inputProps={{ min: 2010, max: 2050 }}
                  value={enrollment.EnrollYear}
                  defaultValue={2021}
                  onChange={handleYearTrimesterChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined" size="small">
                <p>ภาคการศึกษา</p>
                <TextField
                  id="EnrollTrimester"
                  type="number"
                  inputProps={{ min: 1, max: 3 }}
                  value={enrollment.EnrollTrimester}
                  defaultValue={1}
                  onChange={handleYearTrimesterChange}
                />
              </FormControl>
            </Grid>
            
            <Grid item xs={12}>
              <p>รายวิชาที่ลงทะเบียน</p>
              <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="enroll-item-list">
                  <TableHead>
                    <TableRow>
                      <TableCell align="center" width="10%"><b>Course Code</b></TableCell>
                      <TableCell align="center" width="15%"><b>Course Name</b></TableCell>
                      <TableCell align="center" width="5%"><b>Group</b></TableCell>
                      <TableCell align="center" width="5%"><b>Credit</b></TableCell>
                      <TableCell align="center" width="10%"><b>Enroll Type</b></TableCell>
                      <TableCell align="center" width="10%"><b>Total Credit: {totalCredit}</b></TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {enrollmentItemArray?.map((item: Partial<EnrollmentItemsInterface>, index) => (
                      <TableRow key={index}>
                        <TableCell align="center">{item?.ManageCourse?.Course.CourseCode}</TableCell>
                        <TableCell align="left">{item?.ManageCourse?.Course.Name}</TableCell>
                        <TableCell align="center">{item?.ManageCourse?.Group}</TableCell>
                        <TableCell align="center">{item?.ManageCourse?.Course.Credit}</TableCell>
                        <TableCell align="center">{item?.EnrollmentType?.Name}</TableCell>
                        <TableCell align="right"><Button onClick={() => dropItemTable(index)}><DeleteIcon /></Button></TableCell> 
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </Grid>
            
            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined" size="small">
                <p>รายวิชา</p>
                <Select
                  value={selectedCourse.ID}
                  inputProps={{name: "ID"}}
                  onChange={handleCourseChange}
                >
                  <MenuItem value={0} key={0}>เลือกรายวิชา</MenuItem>
                  {courses.map((item: CoursesInterface) => (
                    <MenuItem value={item.ID} key={item.ID}>{item.CourseCode} {item.Name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl fullWidth variant="outlined" size="small" error={hasError}>
                <p>ประเภท</p>
                <Select
                  value={enrollmentItem.EnrollmentTypeID}
                  inputProps={{name: "EnrollmentTypeID"}}
                  onChange={handleEnrollmentItemChange}
                >
                  <MenuItem value={0} key={0}>เลือกประเภท</MenuItem>
                  {enrollmentTypes.map((item: EnrollmentTypesInterface) => (
                    <MenuItem value={item.ID} key={item.ID}>{item.Name}</MenuItem>
                  ))}
                </Select>
                {hasError ? (<FormHelperText>Required</FormHelperText>) : (<FormHelperText></FormHelperText>)}
              </FormControl>
            </Grid>

            <Grid item xs={12}>
              {selectedCourse.ID !== 0 && (
                <Paper className={classes.paper} variant="outlined" >
                  <Table className={classes.table}>
                    <TableBody>
                      <Fragment>
                        {manageCourses.map((item: ManageCoursesInterface, index) => (
                          <TableRow hover key={index}>
                            <TableCell width="5%">
                              <Button onClick={() => addItemTable(item)} color="primary" variant="outlined">
                                เลือก
                              </Button>
                            </TableCell>
                            <TableCell width="40%">
                              <b>Group:</b> {item.Group}<br />
                              <b>Time:</b> {item.TeachingTime}
                            </TableCell>
                            <TableCell width="55%">
                              <b>Room:</b> {item.Room.Number}<br />
                              <b>Lecturer: {item.Professor.TeacherName}</b>
                            </TableCell>
                          </TableRow>
                        ))}
                      </Fragment>
                    </TableBody>
                  </Table>
                </Paper>
              )}
            </Grid>
            
            <Grid item xs={6}>
              <FormControl fullWidth>
                <p>วันที่และเวลา</p>
                <MuiPickersUtilsProvider utils={DateFnsUtils}>
                  <KeyboardDateTimePicker
                    disableToolbar
                    variant="inline"
                    id="enrollDateTime"
                    name="enrollDateTime"
                    value={enrollDateTime}
                    onChange={handleEnrollDateTime}
                    margin="normal"
                    label="กรุณาเลือกวันที่และเวลา"
                    minDate={new Date("2018-01-01T00:00")}
                    format="yyyy/MM/dd hh:mm a"
                    KeyboardButtonProps={{
                      'aria-label': 'change date',
                    }}
                  />
                </MuiPickersUtilsProvider>
              </FormControl>
            </Grid>
            
            <Grid item xs={6}>
              <Button
                onClick={submit}
                color="primary"
                variant="contained" 
                style={{float: "right", marginTop: "75px"}}
              >
                บันทึกการลงทะเบียน
              </Button>
            </Grid>
            
          </Grid>
        </Paper>
      </Container>
    </div>
  );
}