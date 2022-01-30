import React, { Fragment, useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import moment from 'moment';
import Container from '@material-ui/core/Container';
import IconButton from '@material-ui/core/IconButton';
import KeyboardArrowDownIcon from '@material-ui/icons/KeyboardArrowDown';
import KeyboardArrowUpIcon from '@material-ui/icons/KeyboardArrowUp';
import Collapse from '@material-ui/core/Collapse';
import { StudentRecordsInterface } from "../models/IStudentRecord";
import { EnrollmentsInterface, EnrollmentItemsInterface } from "../models/IEnrollRegistration";

const useStyles = makeStyles((theme: Theme) => 
  createStyles({
    container: {marginTop: theme.spacing(3)},
    paper: {padding: theme.spacing(3)},
    table: {minWidth: 650},
    tableSpace: {marginTop: 20},
    row: {'& > !': {borderBottom: 'unset'},},
  })
);

export interface EnrollmentProps {
  enrollment: EnrollmentsInterface;
}

function ListEnrollments(enroll: EnrollmentProps) {
  // open for collapse component
  const [open, setOpen] = useState(false);
  const classes = useStyles();

  return (
    <Fragment>
      <TableRow className={classes.row}>
        <TableCell>
          <IconButton aria-label="expand row" size="small" onClick={() => setOpen(!open)}>
            {open ? <KeyboardArrowUpIcon /> : <KeyboardArrowDownIcon />}
          </IconButton>
        </TableCell>
        <TableCell align="left">{enroll.enrollment.Owner.StudentCode}</TableCell>
        <TableCell align="center">{enroll.enrollment.EnrollYear}</TableCell>
        <TableCell align="center">{enroll.enrollment.EnrollTrimester}</TableCell>
        <TableCell align="center">{moment(enroll.enrollment.EnrollDateTime).format("DD/MM/YYYY hh:mm A")}</TableCell>
        <TableCell align="center">{enroll.enrollment.TotalCredit}</TableCell>
      </TableRow>
      <TableRow>
        <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
          <Collapse in={open} timeout="auto" unmountOnExit>
            <Box margin={1}>
              <Table size="small" aria-label="purchases">
                <TableHead>
                  <TableRow>
                    <TableCell align="center" width="10%"><b>Course Code</b></TableCell>
                    <TableCell align="center" width="15%"><b>Course Name</b></TableCell>
                    <TableCell align="center" width="5%"><b>Group</b></TableCell>
                    <TableCell align="center" width="5%"><b>Credit</b></TableCell>
                    <TableCell align="center" width="10%"><b>Enroll Type</b></TableCell>
                  </TableRow>
                </TableHead>
                <TableBody>
                  {enroll.enrollment.EnrollmentItems.map((clist: EnrollmentItemsInterface) => (
                    <TableRow key={clist.ID}>
                      <TableCell align="center">{clist.ManageCourse.Course.CourseCode}</TableCell>
                      <TableCell align="left">{clist.ManageCourse.Course.Name}</TableCell>
                      <TableCell align="center">{clist.ManageCourse.Group}</TableCell>
                      <TableCell align="center">{clist.ManageCourse.Course.Credit}</TableCell>
                      <TableCell align="center">{clist.EnrollmentType.Name}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </Box>
          </Collapse>
        </TableCell>
      </TableRow>
    </Fragment>
  );
}

export default function Enrollments() {
  const classes = useStyles();
  const [enrolls, setEnrolls] = useState<EnrollmentsInterface[]>([]);

  const getEnrolls = async() => {
    let student: StudentRecordsInterface = JSON.parse(localStorage.getItem("student") || "");
    const apiUrl = `http://localhost:8080/enrollments/student/${student.ID}`;
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
        if (res.data) {
          setEnrolls(res.data);
        } else {
          console.log("else");
        }
      });
  };

  useEffect(() => {
    getEnrolls();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Paper className={classes.paper}>
          <Box display="flex">
            <Box flexGrow={1}>
              <Typography
                component="h2"
                variant="h6"
                color="primary"
                gutterBottom
              >
                ผลการลงทะเบียน
              </Typography>
            </Box>
            <Box flexGrow={1}>
              <Button
                component={RouterLink}
                to="/enrollReg/create"
                variant="contained"
                color="primary"
                style={{float: "right"}}
              >
                ลงทะเบียน
              </Button>
            </Box>
          </Box>
          <TableContainer component={Paper} className={classes.tableSpace}>
            <Table className={classes.table} aria-label="simple-table">
              <TableHead>
                <TableRow>
                  <TableCell width="5%" />
                  <TableCell align="left" width="5%"><b>Student ID</b></TableCell>
                  <TableCell align="center" width="10%"><b>Year</b></TableCell>
                  <TableCell align="center" width="10%"><b>Trimester</b></TableCell>
                  <TableCell align="center" width="15%"><b>Enroll Time</b></TableCell>
                  <TableCell align="center" width="10%"><b>TotalCredit</b></TableCell>
                  <TableCell  width="5%"/>
                </TableRow>
              </TableHead>
              <TableBody>
                {enrolls.map((enroll: EnrollmentsInterface, index) => (
                  <ListEnrollments 
                    key={enroll.ID}
                    enrollment={enroll} />
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Paper>
      </Container>
    </div>
  )
}
