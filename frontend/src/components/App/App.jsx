import React, { Component } from 'react';
import './App.scss';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import {Link} from "react-router-dom";

class App extends Component {
  render() {
    return (

      <div className="home">
        <AppBar position="static">
          <Toolbar>
            <Link to="/">
              <Typography variat="title" color="inherit">
                B-BOY DANCE BATTLE HOME
              </Typography>
            </Link>

            <Link to="/login">
              <Button color="inherit">Login</Button>
            </Link>
            <Link to="/signup">
              <Button color="inherit">Signup</Button>
            </Link>
            <Link to=""></Link>
          </Toolbar>
        </AppBar>
      </div>
    );
  }
}

export default App;
