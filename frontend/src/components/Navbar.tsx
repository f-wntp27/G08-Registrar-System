import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import IconButton from "@material-ui/core/IconButton";
import MenuIcon from "@material-ui/icons/Menu";
import MonetizationOnIcon from "@material-ui/icons/MonetizationOn";
import HistoryIcon from "@material-ui/icons/History";
import Drawer from "@material-ui/core/Drawer";
import List from "@material-ui/core/List";
import ListItem from "@material-ui/core/ListItem";
import ListItemIcon from "@material-ui/core/ListItemIcon";
import ListItemText from "@material-ui/core/ListItemText";
import HomeIcon from "@material-ui/icons/Home";

import { StudentInterface } from "../models/IStudent";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    title: {
      flexGrow: 1,
    },
    small: {
      width: theme.spacing(3),
      height: theme.spacing(3),
    },

    drawer: { width: 250 },
  })
);

export default function ButtonAppBar() {
  const classes = useStyles();
  const users: StudentInterface = JSON.parse(
    localStorage.getItem("student") || ""
  );
  const [openDrawer, setOpenDrawer] = useState(false);

  //ฟังชั่น เปิดปิด
  const toggleDrawer = (state: boolean) => (event: any) => {
    if (
      event.type === "keydown" &&
      (event.key === "Tab" || event.key === "Shift")
    ) {
      return;
    }
    setOpenDrawer(state);
  };

  const SignOut = () => {
    localStorage.clear();
    window.location.href = "/";
  };

  console.log(users);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const menu = [
    { name: "ใบชำระเงิน", icon: <MonetizationOnIcon />, path: "/create" },
    { name: "ประวัติใบชำระเงิน", icon: <HistoryIcon />, path: "/history" },
  ];
  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            edge="start"
            className={classes.menuButton}
            color="inherit"
            aria-label="menu"
            onClick={toggleDrawer(true)}
          >
            <MenuIcon />
          </IconButton>
          <Drawer open={openDrawer} onClose={toggleDrawer(false)}>
            <List
              className={classes.drawer}
              onClick={toggleDrawer(false)}
              onKeyDown={toggleDrawer(false)}
            >
              <ListItem button component={RouterLink} to="/home">
                <ListItemIcon>
                  <HomeIcon />
                </ListItemIcon>
                <ListItemText>หน้าเเรก</ListItemText>
              </ListItem>

              {menu.map((item, index) => (
                <ListItem
                  key={index}
                  button
                  component={RouterLink}
                  to={item.path}
                >
                  <ListItemIcon>{item.icon}</ListItemIcon>
                  <ListItemText>{item.name}</ListItemText>
                </ListItem>
              ))}
            </List>
          </Drawer>
          <Typography variant="h6" className={classes.title}>
            เมนู
          </Typography>
          <div style={{ marginRight: ".5rem" }}>
            <Typography align="right" variant="subtitle2">
              {users?.StudentCode} {users?.Prefix.Value}
              {users?.FirstName} {users?.LastName}
            </Typography>
          </div>
          <div>
            <Button
              onClick={SignOut}
              color="primary"
              variant="contained"
              style={{ marginRight: 12 }}
            >
              ออกจากระบบ
            </Button>
          </div>
        </Toolbar>
      </AppBar>
    </div>
  );
}
