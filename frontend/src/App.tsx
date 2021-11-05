import React, { Fragment, useEffect, useState } from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import NavBar from './components/NavBar';
import Home from './components/Home';
import Enrollments from './components/Enrollment';
import EnrollmentCreate from './components/EnrollmentCreate';
import SignIn from './components/SignIn';

function App() {
  const [token, setToken] = useState<string>("");

  useEffect(() => {
    const getToken = localStorage.getItem("token");
    if (getToken) {
      setToken(getToken);
    }
  }, []);

  if (!token) {
    return <SignIn />
  }

  return (
    <Router>
      <div>
        {token && (
          <Fragment>
            <NavBar />
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/enrollReg/history" component={Enrollments} />
              <Route exact path="/enrollReg/create" component={EnrollmentCreate} />
            </Switch>
          </Fragment>
        )}
      </div>
    </Router>
  );
}

export default App;
