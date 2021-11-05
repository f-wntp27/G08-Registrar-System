import React, { useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Button from "@material-ui/core/Button";
import Drawer from "@material-ui/core/Drawer";
import Divider from "@material-ui/core/Divider";
import IconButton  from "@material-ui/core/IconButton";
import List from "@material-ui/core/List"
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import MenuIcon from "@material-ui/icons/Menu";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";

import HomeIcon from "@material-ui/icons/Home";
import MenuBookIcon from  "@material-ui/icons/MenuBook";
import AssignmentIcon from "@material-ui/icons/Assignment";
import { StudentRecordsInterface } from "../models/IStudentRecord";


const useStyles = makeStyles((theme: Theme) => (
  createStyles({
    root: { 
      display: "flex",
      flexGrow: 1,
    },
    title: { flexGrow: 1 },
    menuButton: { marginRight: theme.spacing(2) },
    drawer: { width: "600" },
  })
));

function NavBar() {
  const classes = useStyles();
  const [openDrawer, setOpenDrawer] = useState(false);
  const student: StudentRecordsInterface = JSON.parse(localStorage.getItem("student") || "");

  const toggleDrawer = (state: boolean) => (event: any) => {
    if (event.type === "keydown" && (event.key === "Tab" || event.key === "Shift")) {
      return;
    }
    setOpenDrawer(state);
  }

  const menu = [
    { name: "ลงทะเบียนรายวิชา", icon: <MenuBookIcon  />, path: "/enrollReg/create" },
    { name: "ผลการลงทะเบียน", icon: <AssignmentIcon  />, path: "/enrollReg/history" },
  ]

  const SignOut = () => {
    localStorage.clear();
    window.location.href = "/";
  }

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton 
            onClick={toggleDrawer(true)} 
            edge="start" 
            className={classes.menuButton} 
            color="inherit" 
            aria-label="menu"
          >
            <MenuIcon />
          </IconButton>          
          <Drawer open={openDrawer} onClose={toggleDrawer(false)} className={classes.drawer} >
            <List   
              onClick={toggleDrawer(false)} 
              onKeyDown={toggleDrawer(false)}
            >
              <ListItem button component={RouterLink} to="/">
                <ListItemIcon><HomeIcon /></ListItemIcon>
                <ListItemText>หน้าแรก</ListItemText>
              </ListItem>
              <Divider />
              {menu.map((item, index) => (
                <ListItem key={index} button component={RouterLink} to={item.path}>
                  <ListItemIcon>{item.icon}</ListItemIcon>
                  <ListItemText>{item.name}</ListItemText>
                </ListItem>
              ))}
            </List>
          </Drawer>
          <Typography variant="h6" className={classes.title}>
            ระบบลงทะเบียนเรียน
          </Typography>
          <Typography variant="subtitle1" style={{ marginRight: 12 }}>
            {student.StudentCode} {student.Prefix.Value}{student.FirstName} {student.LastName}
          </Typography>
          <Button onClick={SignOut} variant="outlined" color="inherit" style={{ marginRight: 12 }}>
            ออกจากระบบ
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default NavBar;